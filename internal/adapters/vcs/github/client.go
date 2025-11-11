package githubadapter

import (
	"context"
	"fmt"

	"github.com/cli/go-gh/pkg/repository"
	"github.com/google/go-github/v53/github"
	"golang.org/x/oauth2"

	"github.com/Tortik3000/gomodinfo/internal/entity"
)

// Client implements RepoContentProvider for GitHub
type Client struct {
	token string
}

const hostCVS = "github.com"

func New(token string) *Client {
	return &Client{
		token: token,
	}
}

func (c *Client) Resolve(repoURL string) (*entity.RepoRef, error) {
	if repoURL == "" {
		return nil, fmt.Errorf("invalid repository reference: empty input")
	}

	r, err := repository.Parse(repoURL)
	if err != nil {
		return nil, fmt.Errorf("invalid repository reference %q: %w", repoURL, err)
	}
	host := r.Host()
	if host == "" {
		host = hostCVS
	}

	return &entity.RepoRef{
		Host:  host,
		Owner: r.Owner(),
		Name:  r.Name(),
	}, nil
}

func (c *Client) GetGoMod(ctx context.Context, ref *entity.RepoRef) ([]byte, error) {
	if ref.Host != hostCVS {
		return nil, fmt.Errorf("unsupported host: %s", ref.Host)
	}

	gh := newGH(ctx, c.token)

	fileContent, _, _, err := gh.Repositories.GetContents(ctx,
		ref.Owner,
		ref.Name,
		"go.mod",
		nil,
	)
	if err != nil {
		return nil, err
	}

	content, err := fileContent.GetContent()
	if err != nil {
		return nil, err
	}

	return []byte(content), nil
}

// newGH creates a github.Client using token if provided.
func newGH(ctx context.Context, token string) *github.Client {
	if token != "" {
		ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: token})
		tc := oauth2.NewClient(ctx, ts)
		return github.NewClient(tc)
	}
	return github.NewClient(nil)
}
