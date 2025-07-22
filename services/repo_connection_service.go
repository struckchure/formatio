package services

import (
	"strconv"

	"github.com/gofiber/fiber/v3"
	"pkg.formatio/dao"
	"pkg.formatio/lib"
	"pkg.formatio/prisma/db"
	"pkg.formatio/types"
)

type RepoConnectionService struct {
	repoConnectionDAO          dao.IRepoConnectionDao
	githubAccountConnectionDAO dao.IGithubAccountConnectionDao
	githubService              IGithubService
}

// ListRepoConnections implements RepoConnectionServiceInterface.
func (r *RepoConnectionService) ListRepoConnections(args types.ListRepoConnectionArgs) (repoConnections []db.RepoConnectionModel, err error) {
	repoConnections, err = r.repoConnectionDAO.ListRepoConnections(args)
	if err != nil {
		return nil, lib.TranslateDAOError(err)
	}

	return repoConnections, nil
}

// CreateRepoConnection implements RepoConnectionServiceInterface.
func (r *RepoConnectionService) CreateRepoConnection(args types.CreateRepoConnectionArgs) (*db.RepoConnectionModel, error) {
	githubAccountConnection, err := r.githubAccountConnectionDAO.ListConnections(
		types.ListGithubAccountConnectionsArgs{UserId: &args.UserId},
	)
	if err != nil {
		return nil, lib.TranslateDAOError(err)
	}

	if len(githubAccountConnection) == 0 {
		return nil, fiber.NewError(fiber.StatusBadRequest, "no github account connection found")
	}

	repoConnection, err := r.repoConnectionDAO.CreateRepoConnection(args)
	if err != nil {
		return nil, lib.TranslateDAOError(err)
	}

	installationId, _ := githubAccountConnection[0].GithubInstallationID()
	_repoId, _ := repoConnection.RepoID()
	repoId, _ := strconv.Atoi(_repoId)
	repoName, _ := repoConnection.RepoName()

	err = r.githubService.DeployRepo(types.DeployRepoArgs{
		InstallationId: installationId,
		MachineId:      repoConnection.MachineID,
		RepoId:         repoId,
		RepoFullName:   repoName,
	})
	if err != nil {
		return nil, fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return repoConnection, nil
}

// GetRepoConnection implements RepoConnectionServiceInterface.
func (r *RepoConnectionService) GetRepoConnection(args types.GetRepoConnectionArgs) (repoConnection *db.RepoConnectionModel, err error) {
	repoConnection, err = r.repoConnectionDAO.GetRepoConnection(args)
	if err != nil {
		return nil, lib.TranslateDAOError(err)
	}

	return repoConnection, nil
}

// UpdateRepoConnection implements RepoConnectionServiceInterface.
func (r *RepoConnectionService) UpdateRepoConnection(args types.UpdateRepoConnectionArgs) (repoConnection *db.RepoConnectionModel, err error) {
	repoConnection, err = r.repoConnectionDAO.UpdateRepoConnection(args)
	if err != nil {
		return nil, lib.TranslateDAOError(err)
	}

	return repoConnection, nil
}

// DeleteRepoConnection implements RepoConnectionServiceInterface.
func (r *RepoConnectionService) DeleteRepoConnection(args types.DeleteRepoConnectionArgs) (err error) {
	err = r.repoConnectionDAO.DeleteRepoConnection(args)
	if err != nil {
		return lib.TranslateDAOError(err)
	}

	return nil
}

func NewRepoConnectionService(
	repoConnectionDAO dao.IRepoConnectionDao,
	githubAccountConnectionDAO dao.IGithubAccountConnectionDao,
	githubService IGithubService,
) *RepoConnectionService {
	return &RepoConnectionService{
		repoConnectionDAO:          repoConnectionDAO,
		githubAccountConnectionDAO: githubAccountConnectionDAO,
		githubService:              githubService,
	}
}
