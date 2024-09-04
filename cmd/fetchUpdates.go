package cmd

import (
	"fmt"
	"os"
	"os/exec"
)

func fetchUpdates(repoURL string, branch string) error {
	if branch == "" {
		branch = "main" // Default branch value
	}
	workDir := fmt.Sprintf("%s/PCWT", getWorkingDir())
	if err := os.Chdir(workDir); err != nil {
		return fmt.Errorf("error changing directory: %w", err)
	}

	remoteCommit := getGitCommit(repoURL, branch)
	localCommit := getLocalCommit(branch)

	if remoteCommit != localCommit {
		fmt.Println(fmt.Sprintf("%sChanges detected in the remote repository.%s", green, reset))
		fmt.Println(fmt.Sprintf("%sPulling the latest changes...%s", blue, reset))
		cmd := exec.Command("git", "pull", "origin", branch)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		if err := cmd.Run(); err != nil {
			return fmt.Errorf("error pulling changes: %w", err)
		}
		return nil
	} else {
		fmt.Println(fmt.Sprintf("%sNo changes detected in the remote repository.%s", yellow, reset))
		return nil
	}
}

func getGitCommit(repo, branch string) string {
	cmd := exec.Command("git", "ls-remote", repo, branch)
	output, err := cmd.Output()
	if err != nil {
		fmt.Println(fmt.Sprintf("%sError getting remote commit: %s%s", red, err, reset))
		os.Exit(1)
	}
	return string(output[:40])
}

func getLocalCommit(branch string) string {
	cmd := exec.Command("git", "rev-parse", branch)
	output, err := cmd.Output()
	if err != nil {
		fmt.Println(fmt.Sprintf("%sError getting local commit: %s%s", red, err, reset))
		os.Exit(1)
	}
	return string(output[:40])
}
