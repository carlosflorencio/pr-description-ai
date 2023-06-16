package ai

import (
	"fmt"
	"os/exec"
	"strings"
)

// Function to execute a shell command
func executeCommand(command string) (string, error) {
	cmd := exec.Command("bash", "-c", command)
	output, err := cmd.Output()
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(output)), nil
}

func CurrentBranchName() (string, error) {
	return executeCommand("git symbolic-ref --short HEAD")
}

// returns a string containing the description of changes for each file
func CompareGitChanges(targetBranch string) (string, error) {

	// Get the current branch
	currentBranch, err := CurrentBranchName()
	if err != nil {
		return "", err
	}

	// Get the list of changed files
	filesChanged, err := executeCommand(fmt.Sprintf("git diff --name-status %s...%s", targetBranch, currentBranch))
	if err != nil {
		return "", err
	}

	// Split the filesChanged string into individual lines
	fileLines := strings.Split(filesChanged, "\n")

	// Generate the description of changes
	description := ""
	for _, fileLine := range fileLines {
		if fileLine == "" {
			continue
		}
		fileInfo := strings.Split(fileLine, "\t")
		filePath := fileInfo[1]

		diff, err := executeCommand(fmt.Sprintf("git diff %s...%s -- %s", targetBranch, currentBranch, filePath))
		if err != nil {
			return "", err
		}

		description += fmt.Sprintf("Changes in file: %s\n\n%s\n\n", filePath, diff)
	}

	return description, nil
}
