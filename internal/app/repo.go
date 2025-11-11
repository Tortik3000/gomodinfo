package app

import (
	"context"
	"fmt"

	"github.com/spf13/cobra"

	proxychecker "github.com/Tortik3000/gomodinfo/internal/adapters/go_versions/golang_proxy"
	gomodparser "github.com/Tortik3000/gomodinfo/internal/adapters/parser/gomod"
	githubadapter "github.com/Tortik3000/gomodinfo/internal/adapters/vcs/github"
	"github.com/Tortik3000/gomodinfo/internal/messages"
	"github.com/Tortik3000/gomodinfo/internal/usecase/moduleinfo"
)

var githubToken string

// repoCmd represents the repo command
var repoCmd = &cobra.Command{
	Use:   "repo [repo-url]",
	Short: "Get info about go module and dependencies",
	Long:  messages.RepoCmdLongInfo,
	Args:  cobra.ExactArgs(1),

	RunE: RunE,
}

func RunE(_ *cobra.Command, args []string) error {
	repoURL := args[0]

	repoClient := githubadapter.New(githubToken)
	parser := gomodparser.New()
	checker := proxychecker.New()
	uc := moduleinfo.NewUseCase(repoClient, parser, checker)

	modInfo, err := uc.GetInfo(context.Background(), repoURL)
	if err != nil {
		return err
	}

	fmt.Println("Module:", modInfo.Name)
	fmt.Println("Go version:", modInfo.Version)
	fmt.Println("Updatable dependencies:")
	if len(modInfo.Deps) == 0 {
		fmt.Println("  (none)")
		return nil
	}

	for _, dep := range modInfo.Deps {
		if dep.UpdateAvailable {
			fmt.Printf(" - %s: %s -> %s\n", dep.Name, dep.CurrentVersion, dep.LatestVersion)
		}
	}
	return nil
}

func init() {
	rootCmd.AddCommand(repoCmd)
	repoCmd.Flags().StringVarP(&githubToken, "token", "t", "", "GitHub token for private repositories")
}
