package gitutils

import (
	"fmt"
	"os/exec"
	"strings"
)

// struct for the git logs
type GitLogs struct {
	rawLogs []byte
}

// gets the logs from the git repository
func (g *GitLogs) getLogs() error {
	cmd := exec.Command("git", "log")
	bLogs, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("git log failed: %w", err)
	}
	g.rawLogs = bLogs
	return nil
}

// returns the short hash of the last commit
func (g *GitLogs) LastCommitHash() (string, error) {
	if err := g.getLogs(); err != nil {
		return "", err
	}
	logLines := strings.Split(string(g.rawLogs), "\n")
	if len(logLines) == 0 {
		return "", fmt.Errorf("no commits found")
	}
	parts := strings.Split(logLines[0], " ")
	if len(parts) < 2 || len(parts[1]) < 6 {
		return "", fmt.Errorf("unexpected git log format: %q", logLines[0])
	}
	return parts[1][:6], nil
}
