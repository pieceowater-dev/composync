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
	workDir := getWorkingDir()
	if err := os.Chdir(workDir); err != nil {
		return fmt.Errorf("error changing directory: %w", err)
	}

	remoteCommit := getGitCommit(repoURL, branch)
	localCommit, err := getLocalCommit(branch)
	if err != nil {
		fmt.Println(fmt.Sprintf("%sLocal branch not found. Attempting to check out %s...%s", yellow, branch, reset))
		// Attempt to check out the branch
		checkoutCmd := exec.Command("git", "checkout", "-b", branch, "--track", "origin/"+branch)
		checkoutCmd.Stdout = os.Stdout
		checkoutCmd.Stderr = os.Stderr
		if checkoutErr := checkoutCmd.Run(); checkoutErr != nil {
			return fmt.Errorf("failed to checkout branch %s: %w", branch, checkoutErr)
		}
		// Retry getting the commit after checkout
		localCommit, err = getLocalCommit(branch)
		if err != nil {
			return fmt.Errorf("error retrieving local commit after checkout: %w", err)
		}
	}

	if remoteCommit != localCommit {
		fmt.Println(fmt.Sprintf("%sChanges detected in the remote repository.%s", green, reset))
		fmt.Println(fmt.Sprintf("%sPulling the latest changes...%s", blue, reset))
		cmd := exec.Command("git", "pull", "origin", branch)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		if err := cmd.Run(); err != nil {
			return fmt.Errorf("error pulling changes: %w", err)
		}
	} else {
		fmt.Println(fmt.Sprintf("%sNo changes detected in the remote repository.%s", yellow, reset))
	}
	return nil
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

func getLocalCommit(branch string) (string, error) {
	cmd := exec.Command("git", "rev-parse", "--verify", branch)
	output, err := cmd.Output()
	if err != nil {
		return "", fmt.Errorf("error getting local commit: %w", err)
	}
	return string(output[:40]), nil
}
