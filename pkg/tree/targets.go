package tree

import (
	"fmt"
	"io/fs"

	"github.com/go-git/go-billy/v5"
)

// Target represents a specific cluster within an environment.
//
// e.g. "eu-west" within the "production" environment.
type Target struct {
	Environment string
	Cluster     string
}

// Targets parses a Kubectl layout filesystem, finding all the Targets based at
// the root.
func Targets(fs billy.Filesystem) ([]Target, error) {
	files, err := fs.ReadDir("/")
	if err != nil {
		return nil, fmt.Errorf("failed to read directory: %w", err)
	}

	var targets []Target
	for _, file := range files {
		if file.IsDir() && !isBaseDir(file) {
			t, err := parseDir(fs, file.Name())
			if err != nil {
				return nil, fmt.Errorf("failed to parse directory %q: %w", file.Name(), err)
			}
			targets = append(targets, t...)
		}
	}
	return targets, nil
}

func parseDir(fs billy.Filesystem, root string) ([]Target, error) {
	files, err := fs.ReadDir(root)
	if err != nil {
		return nil, fmt.Errorf("failed to read directory: %w", err)
	}
	var targets []Target
	for _, file := range files {
		if file.IsDir() && !isBaseDir(file) {
			targets = append(targets, Target{Environment: root, Cluster: file.Name()})
		}
	}
	return targets, nil
}

func isBaseDir(fi fs.FileInfo) bool {
	return fi.Name() == "bases"
}
