package cmd

import (
	"fmt"
	"net/url"
	"os/exec"
	"strings"
)

func storeGitCreds(repoURL, username, token string) error {
	// Parse the repository URL to extract the protocol and host
	parsedURL, err := url.Parse(repoURL)
	if err != nil {
		return fmt.Errorf("invalid repository URL: %w", err)
	}

	// Extract protocol and host from parsed URL
	protocol := strings.Split(parsedURL.Scheme, "+")[0] // Handle schemes like "http+git"
	host := parsedURL.Host

	// Validate that essential parameters are provided
	if protocol == "" || host == "" || username == "" || token == "" {
		return fmt.Errorf("protocol, host, git username, and personal access token must be provided")
	}

	// Configure Git to use credential helper
	gitCmd := exec.Command("git", "config", "--global", "credential.helper", "store")
	if err := gitCmd.Run(); err != nil {
		return fmt.Errorf("error configuring git credential helper: %w", err)
	}

	// Format the credentials for dynamic protocol and host
	creds := fmt.Sprintf("protocol=%s\nhost=%s\nusername=%s\npassword=%s", protocol, host, username, token)
	cmd := exec.Command("bash", "-c", fmt.Sprintf("echo -e \"%s\" | git credential approve", creds))
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("error storing credentials for %s://%s: %w", protocol, host, err)
	}

	fmt.Println(fmt.Sprintf("%sGit credentials stored successfully for %s://%s.%s", green, protocol, host, reset))
	return nil
}
