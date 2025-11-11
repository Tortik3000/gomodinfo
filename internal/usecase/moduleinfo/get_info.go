package moduleinfo

import (
	"context"

	"github.com/Tortik3000/gomodinfo/internal/entity"
)

// GetInfo resolves repo URL, fetches go.mod, parses it and enriches dependencies
func (uc *UseCase) GetInfo(
	ctx context.Context,
	repoURL string,
) (*entity.ModuleInfo, error) {
	ref, err := uc.repo.Resolve(repoURL)
	if err != nil {
		return nil, err
	}

	modBytes, err := uc.repo.GetGoMod(ctx, ref)
	if err != nil {
		return nil, err
	}

	info, err := uc.parser.Parse(modBytes)
	if err != nil {
		return nil, err
	}

	err = uc.checker.Enrich(ctx, info.Deps)
	if err != nil {
		return nil, err
	}
	return info, nil
}
