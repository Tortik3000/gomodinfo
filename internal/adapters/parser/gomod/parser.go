package gomodparser

import (
	"golang.org/x/mod/modfile"

	"github.com/Tortik3000/gomodinfo/internal/entity"
)

// Parser implements GoModParser using golang.org/x/mod/modfile
type Parser struct{}

func New() *Parser {
	return &Parser{}
}

func (p *Parser) Parse(modBytes []byte) (*entity.ModuleInfo, error) {
	mf, err := modfile.Parse("go.mod", modBytes, nil)
	if err != nil {
		return nil, err
	}

	info := &entity.ModuleInfo{}

	if mf.Module != nil {
		info.Name = mf.Module.Mod.Path
	}
	if mf.Go != nil {
		info.Version = mf.Go.Version
	}
	for _, r := range mf.Require {
		info.Deps = append(info.Deps,
			&entity.Dependency{
				Name:           r.Mod.Path,
				CurrentVersion: r.Mod.Version,
			})
	}
	return info, nil
}
