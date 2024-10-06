package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

func setupWorkdir(repoURL string) error {
	workDir := getWorkingDir()

	if _, err := os.Stat(filepath.Join(workDir, ".git")); os.IsNotExist(err) {
		fmt.Println(fmt.Sprintf("%sSetting up the working directory...%s", green, reset))
		if err := os.MkdirAll(workDir, os.ModePerm); err != nil {
			return fmt.Errorf("error creating working directory: %w", err)
		}

		fmt.Println(fmt.Sprintf("%sCloning the repository from %s...%s", green, repoURL, reset))
		cmd := exec.Command("git", "clone", repoURL, workDir)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		if err := cmd.Run(); err != nil {
			return fmt.Errorf("error cloning repository: %w", err)
		}

		fmt.Println(fmt.Sprintf("%sRepository cloned successfully.%s", green, reset))
	} else {
		fmt.Println(fmt.Sprintf("%sRepository already exists. Pulling latest changes...%s", yellow, reset))
		cmd := exec.Command("git", "-C", workDir, "pull")
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		if err := cmd.Run(); err != nil {
			return fmt.Errorf("error pulling latest changes: %w", err)
		}
		fmt.Println(fmt.Sprintf("%sRepository updated successfully.%s", green, reset))
	}
	return nil
}

func getRootDir() string {
	dir, err := os.UserHomeDir()
	if err != nil {
		fmt.Println(fmt.Sprintf("%sError getting root directory: %s%s", red, err, reset))
		os.Exit(1)
	}
	return dir
}

func getWorkingDir() string {
	return filepath.Join(getRootDir(), "PCWT")
}
