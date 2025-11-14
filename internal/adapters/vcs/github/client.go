package githubadapter

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/cli/go-gh/pkg/repository"
	"github.com/google/go-github/v53/github"
	"golang.org/x/oauth2"

	"github.com/Tortik3000/gomodinfo/internal/model"
	modelErr "github.com/Tortik3000/gomodinfo/internal/model/errors/vcs"
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

func (c *Client) Resolve(repoURL string) (*model.RepoRef, error) {
	if repoURL == "" {
		return nil, fmt.Errorf("%w: empty input", modelErr.ErrInvalidRepoReference)
	}

	r, err := repository.Parse(repoURL)
	if err != nil {
		return nil, fmt.Errorf("%w %q: %w", modelErr.ErrInvalidRepoReference, repoURL, err)
	}
	host := r.Host()
	if host == "" {
		host = hostCVS
	}

	return &model.RepoRef{
		Host:  host,
		Owner: r.Owner(),
		Name:  r.Name(),
	}, nil
}

func (c *Client) GetGoMod(ctx context.Context, ref *model.RepoRef) ([]byte, error) {
	if ref.Host != hostCVS {
		return nil, fmt.Errorf("%w: %s", modelErr.ErrUnsupportedHost, ref.Host)
	}

	gh := newGH(ctx, c.token)
	fileContent, _, _, err := gh.Repositories.GetContents(
		ctx,
		ref.Owner,
		ref.Name,
		"go.mod",
		nil,
	)
	if err != nil {
		var rErr *github.ErrorResponse
		if errors.As(err, &rErr) {
			if rErr.Response.StatusCode == http.StatusNotFound {
				return nil, modelErr.ErrNotFound
			}
			return nil, fmt.Errorf("GitHub API error (%d): %w", rErr.Response.StatusCode, err)
		}

		return nil, fmt.Errorf("%w: %v", modelErr.ErrNetwork, err)
	}

	content, err := fileContent.GetContent()
	if err != nil {
		return nil, fmt.Errorf("%w: %v", modelErr.ErrDecodingContent, err)
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
