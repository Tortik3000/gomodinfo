package moduleinfo

import (
	"context"

	"github.com/Tortik3000/gomodinfo/internal/entity"
)

type (
	// repoContentProvider to fetch repository contents (go.mod)
	repoContentProvider interface {
		Resolve(repoURL string) (*entity.RepoRef, error)
		GetGoMod(ctx context.Context, ref *entity.RepoRef) ([]byte, error)
	}

	// goModParser to parse go.mod
	goModParser interface {
		Parse(modBytes []byte) (*entity.ModuleInfo, error)
	}

	// versionChecker enriches dependencies with latest go_versions and update flags
	versionChecker interface {
		Enrich(ctx context.Context, deps []*entity.Dependency) error
	}
)
