package messages

const RepoCmdLongInfo = `Analyzes the go.mod file of the specified repository and shows information about the module and its dependencies.

Arguments:
  repo-url — a repository URL.

What the command does:
  • Clones the repository to a temporary directory.
  • Locates and reads all go.mod files in the repository.
  • Determines the module path and the required Go version.
  • Prints the list of dependencies that have available updates.

Examples:
  gomodinfo repo https://github.com/owner/repo
  gomodinfo repo https://gitlab.com/owner/repo

The output includes: module name, Go version, and the list of dependencies that can be updated.`

const RootCmdLongInfo = `gomodinfo is a command-line tool for analyzing go.mod and a module's dependencies.

Main features:
  • Get module information (module path, required Go version).
  • Print the list of dependencies that have available updates.

Supported commands:
  repo — analyzes a repository's go.mod by URL.

Usage examples:
  gomodinfo repo https://github.com/golang/go
  gomodinfo repo https://gitlab.com/owner/repo

Tips:
  • The repository can be any valid Git URL.`
