package messages

const RepoCmdLongInfo = "Analyzes the go.mod file of the specified GitHub repository and shows information about the module and its dependencies.\n\n" +
	"Arguments:\n" +
	"  repo-url — a GitHub repository URL or a short form owner/repo.\n\n" +
	"Flags:\n" +
	"  -t, --token string  GitHub token for accessing private repositories.\n\n" +
	"What the command does:\n" +
	"  • Locates and reads the go.mod file at the repository root.\n" +
	"  • Determines the module path and the required Go version.\n" +
	"  • Prints the list of dependencies that have available updates.\n\n" +
	"Examples:\n" +
	"  gomodinfo repo https://github.com/owner/repo\n" +
	"  gomodinfo repo owner/repo\n" +
	"  gomodinfo repo owner/private-repo --token $GITHUB_TOKEN\n\n" +
	"The output includes: module name, Go version, and the list of dependencies that can be updated."

const RootCmdLongInfo = "gomodinfo is a command-line tool for analyzing go.mod and a module's dependencies.\n\n" +
	"Main features:\n" +
	"  • Get module information (module path, required Go version).\n" +
	"  • Print the list of dependencies that have available updates.\n\n" +
	"Supported commands:\n" +
	"  repo — analyzes a repository's go.mod by URL.\n\n" +
	"Usage examples:\n" +
	"  gomodinfo repo https://github.com/golang/go\n" +
	"  gomodinfo repo owner/repo\n" +
	"  gomodinfo repo https://github.com/owner/private-repo --token $GITHUB_TOKEN\n\n" +
	"Tips:\n" +
	"  • For private repositories, use the --token flag.\n" +
	"  • The repository can be a full GitHub URL or in the owner/repo form."
