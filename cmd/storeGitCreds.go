package cmd

import (
	"fmt"
	"os/exec"
)

func storeGitCreds(username, token string) error {
	if username == "" || token == "" {
		return fmt.Errorf("git username and personal access token must be provided")
	}

	gitCmd := exec.Command("git", "config", "--global", "credential.helper", "store")
	if err := gitCmd.Run(); err != nil {
		return fmt.Errorf("error configuring git credential helper: %w", err)
	}

	creds := fmt.Sprintf("protocol=https\nhost=github.com\nusername=%s\npassword=%s", username, token)
	cmd := exec.Command("bash", "-c", fmt.Sprintf("echo -e \"%s\" | git credential approve", creds))
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("error storing credentials: %w", err)
	}
	fmt.Println(fmt.Sprintf("%sGit credentials stored successfully.%s", green, reset))
	return nil
}
