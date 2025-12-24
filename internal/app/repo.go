package app

import (
	"context"
	"fmt"

	"github.com/spf13/cobra"

	proxychecker "github.com/Tortik3000/gomodinfo/internal/adapters/go_versions/golist"
	gomodparser "github.com/Tortik3000/gomodinfo/internal/adapters/parser/gomod"
	githubadapter "github.com/Tortik3000/gomodinfo/internal/adapters/vcs/git"
	"github.com/Tortik3000/gomodinfo/internal/messages"
	"github.com/Tortik3000/gomodinfo/internal/usecase/moduleinfo"
)

var repoCmd = &cobra.Command{
	Use:   "repo [repo-url]",
	Short: "Get info about go module and dependencies",
	Long:  messages.RepoCmdLongInfo,
	Args:  cobra.ExactArgs(1),

	RunE: RunE,
}

func RunE(_ *cobra.Command, args []string) error {
	repoURL := args[0]

	repoClient := githubadapter.New()
	parser := gomodparser.New()
	checker := proxychecker.New()
	uc := moduleinfo.NewUseCase(repoClient, parser, checker)

	modsInfo, err := uc.GetInfo(context.Background(), repoURL)
	if err != nil {
		return err
	}

	for _, modInfo := range modsInfo {
		fmt.Println("--------------------------------")
		fmt.Println("Module:", modInfo.Name)
		fmt.Println("Go version:", modInfo.Version)
		fmt.Println("Updatable dependencies:")
		if len(modInfo.Deps) == 0 {
			fmt.Println("  (none)")
			continue
		}

		foundUpdate := false
		for _, dep := range modInfo.Deps {
			if dep.UpdateAvailable {
				fmt.Printf(" - %s: %s -> %s\n", dep.Name, dep.CurrentVersion, dep.LatestVersion)
				foundUpdate = true
			}
		}
		if !foundUpdate {
			fmt.Println("  (none)")
		}
	}
	return nil
}

func init() {
	rootCmd.AddCommand(repoCmd)
}
