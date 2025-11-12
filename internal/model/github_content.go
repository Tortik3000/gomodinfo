package model

type GitHubContent struct {
	Name     string `json:"name"`
	Path     string `json:"path"`
	Encoding string `json:"encoding"`
	Content  string `json:"content"`
}

// RepoRef describes a VCS repository reference resolved from an URL
type RepoRef struct {
	Host  string
	Owner string
	Name  string
}
