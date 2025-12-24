package golist

import (
	"context"
	"encoding/json"
	"os/exec"

	"github.com/Tortik3000/gomodinfo/internal/adapters/go_versions/golist/dto"
	"github.com/Tortik3000/gomodinfo/internal/model"
)

type Checker struct{}

func New() *Checker {
	return &Checker{}
}

func (c *Checker) Enrich(ctx context.Context, deps []*model.Dependency) error {
	for _, d := range deps {
		latest, err := getLatestVersion(ctx, d.Name)
		if err == nil && latest != "" {
			d.LatestVersion = latest
		} else if d.LatestVersion == "" {
			d.LatestVersion = d.CurrentVersion
		}

		if d.CurrentVersion != "" && d.LatestVersion != "" {
			d.UpdateAvailable = d.CurrentVersion != d.LatestVersion
		}
	}
	return nil
}

func getLatestVersion(ctx context.Context, modPath string) (string, error) {
	cmd := exec.CommandContext(ctx, "go", "list", "-m", "-json", modPath+"@latest")
	output, err := cmd.Output()
	if err != nil {
		return "", err
	}

	info := dto.ModuleInfo{}
	if err := json.Unmarshal(output, &info); err != nil {
		return "", err
	}

	return info.Version, nil
}
