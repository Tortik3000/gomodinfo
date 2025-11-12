package gomodparser

import (
	"fmt"

	"golang.org/x/mod/modfile"

	"github.com/Tortik3000/gomodinfo/internal/model"
	modelErr "github.com/Tortik3000/gomodinfo/internal/model/errors/parser"
)

// Parser implements GoModParser using golang.org/x/mod/modfile
type Parser struct{}

func New() *Parser {
	return &Parser{}
}

func (p *Parser) Parse(modBytes []byte) (*model.ModuleInfo, error) {
	if len(modBytes) == 0 {
		return nil, modelErr.ErrEmptyGoMod
	}

	mf, err := modfile.Parse("go.mod", modBytes, nil)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", modelErr.ErrInvalidGoModSyntax, err)
	}

	info := &model.ModuleInfo{}
	if mf.Module == nil || mf.Module.Mod.Path == "" {
		return nil, modelErr.ErrMissingModuleDirective
	}
	info.Name = mf.Module.Mod.Path

	if mf.Go == nil || mf.Go.Version == "" {
		return nil, modelErr.ErrMissingGoVersion
	}
	info.Version = mf.Go.Version

	for _, r := range mf.Require {
		info.Deps = append(info.Deps, &model.Dependency{
			Name:           r.Mod.Path,
			CurrentVersion: r.Mod.Version,
		})
	}

	return info, nil
}
