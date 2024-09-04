package cmd

import (
	"fmt"
	"os"
	"os/exec"
)

func setupWorkdir(repoURL string) error {
	workDir := fmt.Sprintf("%s/PCWT", getWorkingDir())
	if _, err := os.Stat(fmt.Sprintf("%s/.git", workDir)); os.IsNotExist(err) {
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
		return nil
	} else {
		fmt.Println(fmt.Sprintf("%sRepository already exists.%s", yellow, reset))
		return nil
	}
}

func getWorkingDir() string {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println(fmt.Sprintf("%sError getting working directory: %s%s", red, err, reset))
		os.Exit(1)
	}
	return dir
}
