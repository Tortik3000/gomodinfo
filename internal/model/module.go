package model

type ModuleInfo struct {
	Name    string
	Version string
	Deps    []*Dependency
}

type Dependency struct {
	Name            string
	CurrentVersion  string
	LatestVersion   string
	UpdateAvailable bool
}
