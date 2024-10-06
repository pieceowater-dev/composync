package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func scanAndApply(scanDir string, recursive bool) error {
	fullScanDir := filepath.Join(getWorkingDir(), scanDir)

	fmt.Printf("%sScanning directory: %s%s\n", blue, fullScanDir, reset)

	err := filepath.Walk(fullScanDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Printf("%sError accessing path %q: %s%s\n", red, path, err, reset)
			return err
		}

		if info.IsDir() {
			if path != fullScanDir && !recursive {
				return filepath.SkipDir
			}
			return nil
		}

		if strings.HasSuffix(info.Name(), "compose.yml") || strings.HasSuffix(info.Name(), "compose.yaml") {
			fmt.Printf("%sFound Docker Compose file: %s%s\n", green, path, reset)
			if err := applyDockerComposeUpdates(path); err != nil {
				fmt.Printf("%sError applying Docker Compose updates: %s%s\n", red, err, reset)
			}
		}

		return nil
	})

	if err != nil {
		return fmt.Errorf("error scanning directory: %w", err)
	}

	fmt.Printf("%sFinished scanning and applying updates.%s\n", green, reset)
	return nil
}

func commandExists(cmd string) bool {
	_, err := exec.LookPath(cmd)
	if err != nil {
		fmt.Printf("Error looking up command %q: %s\n", cmd, err)
		return false
	}
	return true
}

func isDockerComposeAvailable() bool {
	// Check for legacy docker-compose
	if commandExists("docker-compose") {
		return true
	}
	// Check for docker compose v2 integrated command
	return commandExists("docker") && exec.Command("docker", "compose", "version").Run() == nil
}

func applyDockerComposeUpdates(composeFilePath string) error {
	dir := filepath.Dir(composeFilePath)
	fmt.Printf("%sApplying Docker Compose updates in: %s%s\n", blue, dir, reset)

	if err := os.Chdir(dir); err != nil {
		return fmt.Errorf("error changing directory: %w", err)
	}

	var upCmd *exec.Cmd
	var dockerComposeCmd string

	// Determine which Docker Compose command to use
	if commandExists("docker-compose") {
		dockerComposeCmd = "docker-compose"
	} else if isDockerComposeAvailable() {
		dockerComposeCmd = "docker compose"
	} else {
		return fmt.Errorf("docker compose command not found")
	}

	fmt.Printf("%sUsing Docker Compose command: %s%s\n", blue, dockerComposeCmd, reset)

	// Use the appropriate Docker Compose command with the specific compose file
	if dockerComposeCmd == "docker-compose" {
		upCmd = exec.Command("docker-compose", "-f", composeFilePath, "up", "-d")
	} else {
		upCmd = exec.Command("docker", "compose", "-f", composeFilePath, "up", "-d")
	}

	// Execute the up command
	fmt.Printf("%sRunning: %s\n", blue, upCmd.String())
	upCmd.Stdout = os.Stdout
	upCmd.Stderr = os.Stderr
	if err := upCmd.Run(); err != nil {
		return fmt.Errorf("error applying updates: %w", err)
	}

	// Check if Docker is available before pruning
	if commandExists("docker") {
		pruneCmd := exec.Command("docker", "image", "prune", "-f", "-a")
		fmt.Printf("%sRunning: %s\n", blue, pruneCmd.String())
		pruneCmd.Stdout = os.Stdout
		pruneCmd.Stderr = os.Stderr
		if err := pruneCmd.Run(); err != nil {
			return fmt.Errorf("error pruning images: %w", err)
		}
	} else {
		fmt.Printf("%sWarning: Docker command not found. Skipping image prune.%s\n", yellow, reset)
	}

	fmt.Printf("%sUpdates applied and old images pruned successfully in: %s%s\n", green, dir, reset)
	return nil
}
