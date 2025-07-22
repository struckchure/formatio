package services

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	storm "github.com/Overal-X/formatio.storm"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/go-github/v56/github"
	"github.com/samber/lo"
	"golang.org/x/oauth2"
	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/storage/memory"
	"gopkg.in/yaml.v2"
	v1 "k8s.io/api/core/v1"

	"pkg.formatio/dao"
	"pkg.formatio/lib"
	"pkg.formatio/prisma/db"
	"pkg.formatio/types"
)

type IGithubService interface {
	AuthorizeGithubAccount(types.AuthorizeGithubAccountArgs) (string, error)
	ConnectGithubAccount(types.ConnectGithubAccountArgs) (*string, error)
	DeployRepo(types.DeployRepoArgs) error
	DeployRepoHandler(types.DeployRepoArgs) error
	GetInstallationToken(int) (*string, error)
	ListAccountConnections(types.ListGithubAccountConnectionsArgs) ([]db.GithubAccountConnectionModel, error)
	ListBranches(types.ListBranchesArgs) ([]*github.Branch, error)
	ListCommits(types.ListCommitsArgs) ([]*github.RepositoryCommit, error)
	ListRepositories(types.ListRepositoriesArgs) ([]types.Repository, error)
	UpdateAppAccess(types.AuthorizeGithubAccountArgs) (*string, error)

	exchangeCodeForToken(string) (*oauth2.Token, error)
	generateCloneUrl(string, string) string
	generateJwt() (*string, error)
	getFileFromRepo(types.GetFileFromRepoArgs) (*string, error)
	parseAction(string) (*types.Action, error)
	publishDeploymentLog(types.CreateDeploymentLogArgs) error
}

type GithubService struct {
	env                        lib.Env
	rmq                        lib.RabbitMQ
	redis                      lib.IRedis
	containerManager           lib.IContainerManager
	repoConnectionDao          dao.IRepoConnectionDao
	machineService             IMachineService
	deploymentDAO              dao.IDeploymentDao
	deploymentLogService       *DeploymentLogService
	githubAccountConnectionDAO dao.IGithubAccountConnectionDao
}

func (g *GithubService) DeployRepoHandler(args types.DeployRepoArgs) error {
	repoId := strconv.Itoa(args.RepoId)
	var userId *string

	if args.MachineId != "" {
		githubConnections, err := g.githubAccountConnectionDAO.ListConnections(types.ListGithubAccountConnectionsArgs{
			InstallationId: &args.InstallationId,
		})
		if err != nil {
			return nil
		}

		if len(githubConnections) == 0 {
			return errors.New("no github account connection found")
		}

		userId = &githubConnections[0].UserID
	}

	fmt.Println("userId: ", userId)
	repoConnections, err := g.repoConnectionDao.ListRepoConnections(types.ListRepoConnectionArgs{
		RepoId: &repoId,
		// MachineId: &args.MachineId,
		// OwnerId:   userId,
	})
	if err != nil {
		log.Println(err)

		return err
	}

	if len(repoConnections) < 1 {
		err := errors.New("no connections found for repo")
		log.Println(err)

		return nil
	}

	machineId := repoConnections[0].MachineID
	machine, err := g.machineService.GetMachine(types.GetMachineArgs{Id: &machineId})
	if err != nil {
		log.Println(err)

		return err
	}

	mtx := lib.NewMutext().CreateMutext(machineId)
	mtx.Lock()
	defer mtx.Unlock()

	accessToken, err := g.GetInstallationToken(args.InstallationId)
	if err != nil {
		log.Println(err)

		return err
	}

	if args.CommitHash == nil || args.CommitMessage == nil || args.Author == nil {
		commits, err := g.ListCommits(types.ListCommitsArgs{
			Token:        *accessToken,
			Ref:          args.Ref,
			RepoFullName: args.RepoFullName,
			PageSize:     lo.ToPtr(1),
		})
		if err != nil {
			return err
		}

		args.CommitHash = lo.ToPtr(commits[0].GetSHA())
		args.CommitMessage = lo.ToPtr(commits[0].GetCommit().GetMessage())
		args.Author = lo.ToPtr(commits[0].GetCommit().GetAuthor().GetName())
	}

	cloneURL := g.generateCloneUrl(args.RepoFullName, *accessToken)

	actionFileContent, err := g.getFileFromRepo(
		types.GetFileFromRepoArgs{
			RepoURL:  cloneURL,
			FilePath: ".formatio/action.yaml",
		},
	)
	if err != nil {
		log.Println(err)

		return nil
	}

	actionConfig, err := g.parseAction(*actionFileContent)
	if err != nil {
		return errors.Join(errors.New("could not parse action configuration"), err)
	}

	deployment, err := g.deploymentDAO.CreateDeployment(types.CreateDeploymentArgs{
		MachineId:        machineId,
		RepoConnectionId: repoConnections[0].ID,
		CommitHash:       *args.CommitHash,
		CommitMessage:    *args.CommitMessage,
		Actor:            *args.Author,
		Status:           db.DeploymentStatusInProgress,
	})
	if err != nil {
		log.Println(err)

		return err
	}

	payload, err := json.Marshal(deployment)
	if err != nil {
		log.Println(err)
	}

	err = g.rmq.Publish(lib.PublishArgs{
		Queue:   types.DEPLOYMENT_NOTIFICATION_EVENT,
		Content: string(payload),
	})
	if err != nil {
		log.Println("publishing to DEPLOYMENT_NOTIFICATION_EVENT: ", err)
	}

	deploymentPods, err := g.containerManager.ListDeploymentPods(lo.Must(machine.ContainerID()))
	if err != nil {
		return err
	}

	agent := storm.NewAgent()

	ic := storm.InventoryConfig{}
	for _, pod := range deploymentPods.Items {
		if pod.Status.Phase == v1.PodRunning {
			ic.Servers = append(ic.Servers, storm.Server{
				Name:        pod.Name,
				Host:        pod.Status.PodIP,
				Port:        22,
				User:        "formatio",
				SshPassword: "password",
			})
		}
	}

	if len(ic.Servers) == 0 {
		err := errors.New("pods are offline, inventory is empty")

		if err := g.publishDeploymentLog(types.CreateDeploymentLogArgs{
			DeploymentId: deployment.ID,
			JobId:        "pre-job.Installing Storm",
			Message:      err.Error(),
		}); err != nil {
			return err
		}

		return nil
	}

	err = g.publishDeploymentLog(types.CreateDeploymentLogArgs{
		DeploymentId: deployment.ID,
		JobId:        "pre-job.Installing Storm",
		Message:      "$ storm agent install -m prod -i ./inventory.yaml",
	})
	if err != nil {
		log.Println(err)

		return err
	}

	err = agent.Install(storm.InstallArgs{Ic: ic, Mode: "prod"})
	if err != nil {
		log.Println(err)

		if err := g.publishDeploymentLog(types.CreateDeploymentLogArgs{
			DeploymentId: deployment.ID,
			JobId:        "pre-job.Installing Storm",
			Message:      err.Error(),
		}); err != nil {
			return err
		}

		if _, err := g.deploymentDAO.UpdateDeployment(types.UpdateDeploymentArgs{
			Id:     deployment.ID,
			Status: lo.ToPtr(db.DeploymentStatusFailed),
		}); err != nil {
			return err
		}

		return err
	}

	err = g.publishDeploymentLog(types.CreateDeploymentLogArgs{
		DeploymentId: deployment.ID,
		JobId:        "pre-job.Installing Storm",
		Message:      "Storm is Ready!",
	})
	if err != nil {
		log.Println(err)

		return err
	}

	logCallback := func(i any) {
		stringPayload := i.(string)
		stringPayloadLines := strings.Split(stringPayload, "\n")
		for _, line := range stringPayloadLines {
			stepOutput := storm.WorkflowStepOutputStruct{}
			err := json.Unmarshal([]byte(line), &stepOutput)
			if err != nil {
				log.Println("$ ", line, "\n", err)

				return
			}

			if !strings.HasPrefix(stepOutput.Path, "__builtin__") {
				err := g.publishDeploymentLog(types.CreateDeploymentLogArgs{
					DeploymentId: deployment.ID,
					JobId:        stepOutput.Path,
					Message:      "$ " + stepOutput.Command,
				})

				if err != nil {
					log.Println(err)

					return
				}

				err = g.publishDeploymentLog(types.CreateDeploymentLogArgs{
					DeploymentId: deployment.ID,
					JobId:        stepOutput.Path,
					Message:      stepOutput.Message,
				})

				if err != nil {
					log.Println(err)

					return
				}
			}
		}
	}

	setupWc := storm.WorkflowConfig{
		Name:      "Setup",
		Directory: "/home/formatio",
		Jobs: []storm.Job{
			{
				Name: "pre-job",
				Steps: []storm.Step{
					{
						Name: "Remove old repo",
						Run:  `sudo rm -rf code || echo "Directory 'code' does not exist, nothing to remove"`,
					},
					{
						Name: "Cloning git repo",
						Run:  fmt.Sprintf("git clone %s code", cloneURL),
					},
				},
			},
		},
	}
	if args.Ref != "" {
		setupWc.Jobs[0].Steps = append(setupWc.Jobs[0].Steps, storm.Step{
			Name: "Checkout commit",
			Run:  fmt.Sprintf("cd code && git checkout %s", args.Ref),
		})
	}

	err = agent.Run(
		agent.AgentWithCallback(logCallback, storm.StepOutputTypeJson),
		agent.AgentWithConfigs(setupWc, ic),
	)

	if err != nil {
		log.Println(err)

		if _, err := g.deploymentDAO.UpdateDeployment(types.UpdateDeploymentArgs{
			Id:     deployment.ID,
			Status: lo.ToPtr(db.DeploymentStatusFailed),
		}); err != nil {
			return err
		}

		return err
	}

	wc := storm.WorkflowConfig{
		Name:      actionConfig.Name,
		Directory: "/home/formatio/code",
		Jobs:      actionConfig.Jobs,
	}
	// TODO: check storm version on remove machine matches local
	// TODO: add agent's callback func for logs
	err = agent.Run(
		agent.AgentWithConfigs(wc, ic),
		agent.AgentWithCallback(logCallback, storm.StepOutputTypeJson),
	)
	if err != nil {
		log.Println(err)

		_, err := g.deploymentDAO.UpdateDeployment(types.UpdateDeploymentArgs{
			Id:     deployment.ID,
			Status: lo.ToPtr(db.DeploymentStatusFailed),
		})

		return err
	}

	_, err = g.deploymentDAO.UpdateDeployment(types.UpdateDeploymentArgs{
		Id:     deployment.ID,
		Status: lo.ToPtr(db.DeploymentStatusSuccessful),
	})

	return err
}

func (g *GithubService) DeployRepo(args types.DeployRepoArgs) error {
	payload, err := json.Marshal(args)
	if err != nil {
		log.Println(err)

		return err
	}

	return g.rmq.Publish(lib.PublishArgs{
		Queue:   types.DEPLOYMENT_DEPLOY_REPO_QUEUE,
		Content: string(payload),
	})
}

func (g *GithubService) ListBranches(args types.ListBranchesArgs) ([]*github.Branch, error) {
	// Set up OAuth2 with your access token
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: args.Token})
	tc := oauth2.NewClient(ctx, ts)

	// Create a new GitHub client
	client := github.NewClient(tc)

	owner, _ := lo.First(strings.Split(args.RepoFullName, "/"))
	repo, _ := lo.Last(strings.Split(args.RepoFullName, "/"))

	args.Page = lo.Ternary(args.Page == nil, lo.ToPtr(1), args.Page)
	args.PageSize = lo.Ternary(args.PageSize == nil, lo.ToPtr(10), args.PageSize)

	branches, _, err := client.Repositories.ListBranches(ctx, owner, repo, &github.BranchListOptions{
		ListOptions: github.ListOptions{Page: *args.Page, PerPage: *args.PageSize},
	})

	return branches, err
}

func (g *GithubService) ListCommits(args types.ListCommitsArgs) ([]*github.RepositoryCommit, error) {
	// Set up OAuth2 with your access token
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: args.Token})
	tc := oauth2.NewClient(ctx, ts)

	// Create a new GitHub client
	client := github.NewClient(tc)

	owner, _ := lo.First(strings.Split(args.RepoFullName, "/"))
	repo, _ := lo.Last(strings.Split(args.RepoFullName, "/"))

	args.Page = lo.Ternary(args.Page == nil, lo.ToPtr(1), args.Page)
	args.PageSize = lo.Ternary(args.PageSize == nil, lo.ToPtr(10), args.PageSize)

	// Get the list of commits from the repository (replace "main" with your branch)
	commits, _, err := client.Repositories.ListCommits(ctx, owner, repo, &github.CommitsListOptions{
		SHA:         args.Ref,
		ListOptions: github.ListOptions{Page: *args.Page, PerPage: *args.PageSize},
	})
	if err != nil {
		log.Printf("error fetching commits: %v", err)

		return nil, err
	}

	if len(commits) == 0 {
		return nil, errors.New("no commits found")
	}

	return commits, nil
}

func (g *GithubService) ListRepositories(args types.ListRepositoriesArgs) ([]types.Repository, error) {
	userConnections, err := g.githubAccountConnectionDAO.ListConnections(types.ListGithubAccountConnectionsArgs{
		UserId: &args.UserId,
	})
	if err != nil {
		return nil, err
	}

	if len(userConnections) == 0 {
		return []types.Repository{}, errors.New("no github account has been connected")
	}

	token, err := g.generateJwt()
	if err != nil {
		return nil, err
	}

	client := github.NewClient(
		oauth2.NewClient(
			context.Background(),
			oauth2.StaticTokenSource(&oauth2.Token{AccessToken: *token}),
		))

	githubUsername, _ := userConnections[0].GithubUsername()

	installation, _, err := client.Apps.FindUserInstallation(context.Background(), githubUsername)
	if err != nil {
		return nil, err
	}

	installationToken, _ := g.GetInstallationToken(int(*installation.ID))

	args.PageNumber = lo.Ternary(args.PageNumber == nil, lo.ToPtr(1), args.PageNumber)
	args.PageSize = lo.Ternary(args.PageSize == nil, lo.ToPtr(10), args.PageSize)

	_repos, _, err := github.NewClient(nil).WithAuthToken(*installationToken).Apps.ListRepos(
		context.Background(),
		&github.ListOptions{Page: *args.PageNumber, PerPage: *args.PageSize},
	)
	if err != nil {
		return nil, err
	}

	var repos []types.Repository = make([]types.Repository, 0)
	for _, repo := range _repos.Repositories {
		repos = append(
			repos,
			types.Repository{
				ID:       int(*repo.ID),
				Name:     *repo.Name,
				FullName: *repo.FullName,
				Private:  *repo.Private,
				HTMLURL:  *repo.HTMLURL,
			},
		)
	}

	return repos, nil
}

func (g *GithubService) AuthorizeGithubAccount(args types.AuthorizeGithubAccountArgs) (string, error) {
	link := fmt.Sprintf("https://github.com/login/oauth/authorize?client_id=%s&login=&state=%s", g.env.GH_APP_CLIENT_ID, args.UserId)

	g.redis.SetItem(args.UserId, args.RedirectUrl)

	return link, nil
}

func (g *GithubService) ConnectGithubAccount(args types.ConnectGithubAccountArgs) (*string, error) {
	token, err := g.exchangeCodeForToken(args.Code)
	if err != nil {
		return nil, err
	}

	// You now have the user access token in 'token'.
	// You can use it to make authenticated requests to the GitHub API.
	// For example:
	oauthClient := oauth2.NewClient(context.TODO(), oauth2.StaticTokenSource(token))
	httpClient := &http.Client{Transport: oauthClient.Transport}

	// Use the httpClient to make authenticated requests to the GitHub API.
	// For instance, to get the user's information:
	resp, err := httpClient.Get("https://api.github.com/user")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	responseBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// Create a User struct
	var user types.GithubUser
	if err := json.Unmarshal(responseBytes, &user); err != nil {
		return nil, err
	}

	ghToken, err := g.generateJwt()
	if err != nil {
		return nil, err
	}
	ghClient := github.NewClient(
		oauth2.NewClient(
			context.Background(),
			oauth2.StaticTokenSource(&oauth2.Token{AccessToken: *ghToken}),
		),
	)
	installation, _, err := ghClient.Apps.FindUserInstallation(context.Background(), user.Username)
	if err != nil {
		return nil, err
	}

	_, err = g.githubAccountConnectionDAO.CreateConnection(
		types.CreateGithubAccountConnectionArgs{
			UserId:               args.UserId,
			GithubId:             strconv.Itoa(user.Id),
			GithubInstallationId: lo.ToPtr(int(*installation.ID)),
			GithubEmail:          user.Email,
			GithubUsername:       user.Username,
		},
	)
	if err != nil {
		switch e := err.(type) {
		case lib.DatabaseError:
			if e.ErrorCode != lib.ErrorCodeDuplicateEntry {
				return nil, lib.TranslateDAOError(e)
			}
		default:
			return nil, lib.TranslateDAOError(e)
		}
	}

	redirectUrl, err := g.redis.GetItem(args.UserId)
	if err != nil {
		return nil, err
	}

	g.redis.DeleteItem(args.UserId)

	return redirectUrl, nil
}

func (g *GithubService) UpdateAppAccess(args types.AuthorizeGithubAccountArgs) (*string, error) {
	userConnections, err := g.githubAccountConnectionDAO.ListConnections(types.ListGithubAccountConnectionsArgs{
		UserId: &args.UserId,
	})
	if err != nil {
		return nil, err
	}

	if len(userConnections) == 0 {
		return nil, lib.HttpError{Message: "github acccount not authorized"}
	}

	token, err := g.generateJwt()
	if err != nil {
		return nil, err
	}

	client := github.NewClient(
		oauth2.NewClient(
			context.Background(),
			oauth2.StaticTokenSource(&oauth2.Token{AccessToken: *token}),
		))

	installation, _, err := client.Apps.FindUserInstallation(context.Background(), lo.Must(userConnections[0].GithubUsername()))
	if err != nil {
		errResponse := err.(*github.ErrorResponse)
		if errResponse.Response.StatusCode == http.StatusNotFound {
			installationLink := fmt.Sprintf("https://github.com/apps/%s/installations/new", g.env.GH_APP_SLUG)

			return &installationLink, nil
		}

		return nil, err
	}

	_, err = g.githubAccountConnectionDAO.UpdateConnection(types.UpdateGithubAccountConnectionArgs{
		Id:                   userConnections[0].ID,
		GithubInstallationId: lo.ToPtr(int(*installation.ID)),
	})
	if err != nil {
		return nil, lib.TranslateDAOError(err)
	}

	installtionUpdateLink := fmt.Sprintf("https://github.com/settings/installations/%d", int(*installation.ID))

	return &installtionUpdateLink, nil
}

func (g *GithubService) ListAccountConnections(args types.ListGithubAccountConnectionsArgs) (
	connections []db.GithubAccountConnectionModel,
	err error,
) {
	connections, err = g.githubAccountConnectionDAO.ListConnections(types.ListGithubAccountConnectionsArgs{
		UserId: args.UserId,
	})
	if err != nil {
		return nil, lib.TranslateDAOError(err)
	}

	return connections, nil
}

func (g *GithubService) GetInstallationToken(installationId int) (*string, error) {
	appToken, err := g.generateJwt()
	if err != nil {
		return nil, err
	}

	token, _, err := github.NewClient(nil).
		WithAuthToken(*appToken).
		Apps.CreateInstallationToken(
		context.Background(),
		int64(installationId),
		&github.InstallationTokenOptions{},
	)
	if err != nil {
		return nil, err
	}

	return token.Token, nil
}

func (g *GithubService) generateJwt() (*string, error) {
	// Generate JWT as before
	token := jwt.New(jwt.SigningMethodRS256)
	token.Claims = jwt.MapClaims{
		"iss": g.env.GH_APP_ID,
		"iat": jwt.NewNumericDate(time.Now()),
		"exp": jwt.NewNumericDate(time.Now().Add(10 * time.Minute)),
	}

	key, err := jwt.ParseRSAPrivateKeyFromPEM([]byte(g.env.GH_PRIVATE_KEY))
	if err != nil {
		return nil, err
	}

	tokenString, err := token.SignedString(key)
	if err != nil {
		return nil, err
	}

	return &tokenString, nil
}

func (g *GithubService) generateCloneUrl(repoFullName string, accessToken string) string {
	return fmt.Sprintf(
		"https://%s:%s@github.com/%s.git",
		"x-access-token",
		accessToken,
		repoFullName,
	)
}

func (g *GithubService) getFileFromRepo(args types.GetFileFromRepoArgs) (*string, error) {
	// TODO: handler error if file does not exist

	// Clone the repository into memory
	r, err := git.Clone(memory.NewStorage(), nil, &git.CloneOptions{
		URL: args.RepoURL,
	})
	if err != nil {
		return nil, err
	}

	// Resolve the revision (e.g., "HEAD")
	ref, err := r.Head()
	if err != nil {
		return nil, err
	}

	// Get the commit object
	commit, err := r.CommitObject(ref.Hash())
	if err != nil {
		return nil, err
	}

	// Get the tree associated with the commit
	tree, err := commit.Tree()
	if err != nil {
		return nil, err
	}

	// Traverse the tree to find the file
	fileBlob, err := tree.File(args.FilePath)
	if err != nil {
		return nil, err
	}

	// Open the file reader
	reader, err := fileBlob.Reader()
	if err != nil {
		return nil, err
	}
	defer reader.Close()

	// Read and print the file content
	content := make([]byte, fileBlob.Size)
	_, err = reader.Read(content)
	if err != nil {
		return nil, err
	}

	stringContent := string(content)

	return &stringContent, nil
}

func (g *GithubService) exchangeCodeForToken(code string) (*oauth2.Token, error) {
	config := oauth2.Config{
		ClientID:     g.env.GH_APP_CLIENT_ID,
		ClientSecret: g.env.GH_APP_CLIENT_SECRET,
		Endpoint: oauth2.Endpoint{
			AuthURL:  "https://github.com/login/oauth/authorize",
			TokenURL: "https://github.com/login/oauth/access_token",
		},
	}

	// Exchange the code for a token
	token, err := config.Exchange(context.TODO(), code)
	if err != nil {
		return nil, err
	}

	return token, nil
}

func (g *GithubService) publishDeploymentLog(args types.CreateDeploymentLogArgs) error {
	deploymentLog, err := g.deploymentLogService.CreateDeploymentLog(args)
	if err != nil {
		return err
	}

	payload, err := json.Marshal(deploymentLog)
	if err != nil {
		return err
	}

	err = g.rmq.Publish(lib.PublishArgs{
		Queue:   types.DEPLOYMENT_LOG_EVENT_QUEUE,
		Content: string(payload),
	})
	if err != nil {
		return err
	}

	return nil
}

func (g *GithubService) parseAction(content string) (*types.Action, error) {
	contentBytes := []byte(content)

	// Unmarshal YAML data into the struct
	var config types.Action
	if err := yaml.Unmarshal(contentBytes, &config); err != nil {
		return nil, err
	}

	return &config, nil
}

func NewGithubService(
	rmq lib.RabbitMQ,
	redis lib.IRedis,
	env lib.Env,
	container lib.IContainerManager,

	machineService IMachineService,
	repoConnectionDao dao.IRepoConnectionDao,
	deploymentDAO dao.IDeploymentDao,
	deploymentLogService *DeploymentLogService,
	githubAccountConnectionDAO dao.IGithubAccountConnectionDao,
) IGithubService {
	return &GithubService{
		env: env,

		rmq:              rmq,
		redis:            redis,
		containerManager: container,

		repoConnectionDao: repoConnectionDao,
		machineService:    machineService,

		deploymentDAO:              deploymentDAO,
		deploymentLogService:       deploymentLogService,
		githubAccountConnectionDAO: githubAccountConnectionDAO,
	}
}
