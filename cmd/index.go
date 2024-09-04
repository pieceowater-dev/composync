package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"net/url"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "composync",
	Short: "Composync automates updating Docker Compose containers",
	Long:  `Composync automates the updating of your Docker Compose containers by continuously pulling and applying the latest changes from your remote repositories' Docker Compose files.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(fmt.Sprintf("%s%s%s", green, boldText("--- COMPOSYNC! ---"), reset))
		fmt.Println(fmt.Sprintf("%s%s%s", blue, "Composync automates the updating of your Docker Compose containers\n"+
			"by continuously pulling and applying the latest changes from your remote repositories' Docker Compose files.\n"+
			"Enjoy seamless Docker Compose updates!", reset))
		fmt.Println(fmt.Sprintf("by %s%s%s", bold, "pieceowater", reset))
	},
}

var goCmd = &cobra.Command{
	Use:   "go",
	Short: "Run all commands sequentially",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(fmt.Sprintf("%s%s%s", green, boldText("--- COMPOSYNC! ---"), reset))
		fmt.Println(fmt.Sprintf("by %s%s%s", bold, "pieceowater", reset))

		// Retrieve flag values
		repoURL, _ := cmd.Flags().GetString("repo")
		branch, _ := cmd.Flags().GetString("branch")
		scanDir, _ := cmd.Flags().GetString("scan-dir")
		recursive, _ := cmd.Flags().GetBool("recursive")
		username, _ := cmd.Flags().GetString("username")
		token, _ := cmd.Flags().GetString("token")

		// Set default values
		if branch == "" {
			branch = "main" // Default branch value
		}
		if scanDir == "" {
			scanDir = "/" // Default scan directory
		}

		// Validate flags
		if repoURL == "" {
			fmt.Println(fmt.Sprintf("%sError: Repository URL is required.%s", red, reset))
			os.Exit(1)
		}
		if err := validateURL(repoURL); err != nil {
			fmt.Println(fmt.Sprintf("%sError: Invalid repository URL. %s%s", red, err, reset))
			os.Exit(1)
		}
		if username == "" {
			fmt.Println(fmt.Sprintf("%sError: Git username is required.%s", red, reset))
			os.Exit(1)
		}
		if token == "" {
			fmt.Println(fmt.Sprintf("%sError: Git personal access token is required.%s", red, reset))
			os.Exit(1)
		}

		// Output flags for confirmation
		cmd.Flags().VisitAll(func(flag *pflag.Flag) {
			fmt.Printf("%s%s=%s%s\n", green, flag.Name, flag.Value, reset)
		})
		fmt.Printf("%sUsing repository: %s%s\n", blue, repoURL, reset)

		// Execute commands
		commands := []struct {
			name string
			fn   func() error
		}{
			{"storeGitCreds", func() error {
				return storeGitCreds(username, token)
			}},
			{"setupWorkdir", func() error {
				return setupWorkdir(repoURL)
			}},
			{"fetchUpdates", func() error {
				return fetchUpdates(repoURL, branch)
			}},
			{"scanAndApply", func() error {
				return scanAndApply(scanDir, recursive)
			}},
		}

		for _, c := range commands {
			fmt.Printf("%sRunning %s...%s\n", blue, c.name, reset)
			if err := c.fn(); err != nil {
				fmt.Printf("%sError executing %s: %s%s\n", red, c.name, err, reset)
				os.Exit(1)
			}
		}
	},
}

func validateURL(u string) error {
	_, err := url.ParseRequestURI(u)
	if err != nil {
		return fmt.Errorf("invalid URL format")
	}
	return nil
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.PersistentFlags().String("repo", "", "Repository URL")
	rootCmd.PersistentFlags().String("branch", "main", "Git branch")
	rootCmd.PersistentFlags().String("scan-dir", "/", "Directory to scan")
	rootCmd.PersistentFlags().Bool("recursive", false, "Scan directories recursively")
	rootCmd.PersistentFlags().String("username", "", "Git username")
	rootCmd.PersistentFlags().String("token", "", "Git personal access token")

	rootCmd.AddCommand(goCmd)
}

const (
	reset  = "\033[0m"
	red    = "\033[31m"
	green  = "\033[32m"
	yellow = "\033[33m"
	blue   = "\033[34m"
	bold   = "\033[1m"
)

func boldText(text string) string {
	return fmt.Sprintf("%s%s%s", bold, text, reset)
}
