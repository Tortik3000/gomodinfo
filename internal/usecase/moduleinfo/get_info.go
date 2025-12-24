package moduleinfo

import (
	"context"
	"fmt"

	"github.com/Tortik3000/gomodinfo/internal/model"
)

func (uc *UseCase) GetInfo(
	ctx context.Context,
	repoURL string,
) ([]*model.ModuleInfo, error) {
	_, err := uc.repo.Resolve(repoURL)
	if err != nil {
		return nil, err
	}

	mods, err := uc.repo.GetGoMods(ctx, repoURL)
	if err != nil {
		return nil, err
	}

	var results []*model.ModuleInfo
	for relPath, modBytes := range mods {
		info, err := uc.parser.Parse(modBytes)
		if err != nil {
			return nil, fmt.Errorf("can't parse %s, %w", relPath, err)
		}
		if relPath != "go.mod" {
			info.Name = fmt.Sprintf("%s (%s)", info.Name, relPath)
		}

		err = uc.checker.Enrich(ctx, info.Deps)
		if err != nil {
			return nil, err
		}
		results = append(results, info)
	}

	return results, nil
}
