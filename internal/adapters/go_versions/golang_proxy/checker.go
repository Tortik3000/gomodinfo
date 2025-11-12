package proxychecker

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"path"

	"golang.org/x/mod/module"
	"golang.org/x/mod/semver"

	"github.com/Tortik3000/gomodinfo/internal/model"
	modelErr "github.com/Tortik3000/gomodinfo/internal/model/errors/go_versions"
)

// Checker implements VersionChecker using the public Go golang_proxy
type Checker struct{}

func New() *Checker { return &Checker{} }

func (c *Checker) Enrich(_ context.Context, deps []*model.Dependency) error {
	for _, d := range deps {
		latest, err := latestVersion(d.Name)
		if err == nil && latest != "" {
			d.LatestVersion = latest
		} else if d.LatestVersion == "" {
			d.LatestVersion = d.CurrentVersion
		}

		cur, lat := d.CurrentVersion, d.LatestVersion
		if cur != "" && lat != "" && semver.IsValid(cur) && semver.IsValid(lat) {
			d.UpdateAvailable = semver.Compare(lat, cur) > 0
		} else {
			d.UpdateAvailable = (lat != "") && (lat != cur)
		}
	}
	return nil
}

func latestVersion(modPath string) (string, error) {
	esc, err := module.EscapePath(modPath)
	if err != nil {
		return "", fmt.Errorf("%w: %v", modelErr.ErrInvalidModulePath, err)
	}

	u := url.URL{
		Scheme: "https",
		Host:   "proxy.golang.org",
		Path:   path.Join("/", esc, "@latest"),
	}

	req, _ := http.NewRequestWithContext(context.Background(), http.MethodGet, u.String(), nil)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", fmt.Errorf("%w: %v", modelErr.ErrProxyUnavailable, err)
	}
	defer resp.Body.Close()

	switch resp.StatusCode {
	case http.StatusOK:
	case http.StatusNotFound, http.StatusGone:
		return "", modelErr.ErrModuleNotFound
	default:
		return "", fmt.Errorf("%w: unexpected status %s", modelErr.ErrInvalidProxyResponse, resp.Status)
	}

	var data struct{ Version string }
	if err = json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return "", fmt.Errorf("%w: %v", modelErr.ErrInvalidProxyResponse, err)
	}

	if data.Version == "" {
		return "", modelErr.ErrInvalidProxyResponse
	}

	return data.Version, nil
}
