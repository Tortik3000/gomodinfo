package moduleinfo

import (
	"context"

	"github.com/Tortik3000/gomodinfo/internal/model"
)

type (
	// repoContentProvider to fetch repository contents (go.mod)
	repoContentProvider interface {
		Resolve(repoURL string) (*model.RepoRef, error)
		GetGoMods(ctx context.Context, repoURL string) (map[string][]byte, error)
	}

	// goModParser to parse go.mod
	goModParser interface {
		Parse(modBytes []byte) (*model.ModuleInfo, error)
	}

	// versionChecker enriches dependencies with latest go_versions and update flags
	versionChecker interface {
		Enrich(ctx context.Context, deps []*model.Dependency) error
	}
)
