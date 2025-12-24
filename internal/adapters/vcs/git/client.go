package gitadapter

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/Tortik3000/gomodinfo/internal/model"
	modelErr "github.com/Tortik3000/gomodinfo/internal/model/errors/vcs"
)

type Client struct{}

func New() *Client {
	return &Client{}
}

func (c *Client) Resolve(repoURL string) (*model.RepoRef, error) {
	if repoURL == "" {
		return nil, fmt.Errorf("%w: empty input", modelErr.ErrInvalidRepoReference)
	}
	return &model.RepoRef{
		Name: repoURL,
	}, nil
}

func (c *Client) GetGoMods(ctx context.Context, repoURL string) (map[string][]byte, error) {
	tmpDir, err := os.MkdirTemp("", "gomodinfo-*")
	if err != nil {
		return nil, fmt.Errorf("failed to create temp dir: %w", err)
	}
	defer os.RemoveAll(tmpDir)

	cmd := exec.CommandContext(ctx, "git", "clone", "--depth", "1", repoURL, tmpDir)
	if err := cmd.Run(); err != nil {
		return nil, fmt.Errorf("%w: git clone failed: %v", modelErr.ErrNetwork, err)
	}

	mods := make(map[string][]byte)
	err = filepath.Walk(tmpDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && info.Name() == "go.mod" {
			content, err := os.ReadFile(path)
			if err != nil {
				return err
			}
			relPath, _ := filepath.Rel(tmpDir, path)
			mods[relPath] = content
		}
		return nil
	})

	if err != nil {
		return nil, fmt.Errorf("failed to walk repository: %w", err)
	}

	if len(mods) == 0 {
		return nil, modelErr.ErrNotFound
	}

	return mods, nil
}
