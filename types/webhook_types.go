package types

type Commit struct {
	ID        string    `json:"id"`
	TreeID    string    `json:"tree_id"`
	Distinct  bool      `json:"distinct"`
	Message   string    `json:"message"`
	Timestamp string    `json:"timestamp"`
	URL       string    `json:"url"`
	Author    Author    `json:"author"`
	Committer Committer `json:"committer"`
	Added     []string  `json:"added"`
	Removed   []string  `json:"removed"`
	Modified  []string  `json:"modified"`
}

type Author struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Username string `json:"username"`
}

type Committer struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Username string `json:"username"`
}

type PushEvent struct {
	Ref        string `json:"ref"`
	Before     string `json:"before"`
	After      string `json:"after"`
	Repository struct {
		ID       int    `json:"id"`
		Name     string `json:"name"`
		FullName string `json:"full_name"`
		Private  bool   `json:"private"`
		HTMLURL  string `json:"html_url"`
	} `json:"repository"`
	Pusher       Pusher       `json:"pusher"`
	Sender       Sender       `json:"sender"`
	Installation Installation `json:"installation"`
	Created      bool         `json:"created"`
	Deleted      bool         `json:"deleted"`
	Forced       bool         `json:"forced"`
	BaseRef      interface{}  `json:"base_ref"`
	Compare      string       `json:"compare"`
	Commits      []Commit     `json:"commits"`
	HeadCommit   Commit       `json:"head_commit"`
}

type Pusher struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type Sender struct {
	Login     string `json:"login"`
	ID        int    `json:"id"`
	NodeID    string `json:"node_id"`
	AvatarURL string `json:"avatar_url"`
	URL       string `json:"url"`
	HTMLURL   string `json:"html_url"`
}

type Installation struct {
	ID     int    `json:"id"`
	NodeID string `json:"node_id"`
}

type Permissions struct {
	Contents string `json:"contents"`
	Metadata string `json:"metadata"`
}

type TokenInfo struct {
	Token               string      `json:"token"`
	ExpiresAt           string      `json:"expires_at"`
	Permissions         Permissions `json:"permissions"`
	RepositorySelection string      `json:"repository_selection"`
}

type AuthorizeGithubAccountArgs struct {
	UserId      string `swaggerignore:"true"`
	RedirectUrl string `query:"redirectUrl"`
}

type ConnectGithubAccountArgs struct {
	UserId string `json:"userId"`
	Code   string `json:"code"`
}

type ListBranchesArgs struct {
	Token        string
	RepoFullName string
	PageSize     *int
	Page         *int
}

type ListCommitsArgs struct {
	Token        string
	Ref          string
	RepoFullName string
	PageSize     *int
	Page         *int
}

type DeployRepoArgs struct {
	InstallationId int
	RepoId         int
	MachineId      string
	RepoFullName   string
	Ref            string
	CommitHash     *string
	CommitMessage  *string
	Author         *string
}
