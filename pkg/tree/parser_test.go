package tree

import (
	"fmt"
	"io/fs"
	"testing"

	"github.com/go-git/go-billy/v5"
	"github.com/go-git/go-billy/v5/osfs"
	"github.com/google/go-cmp/cmp"
)

type Target struct {
	Environment string
	Cluster     string
}

func Environments(fs billy.Filesystem) ([]Target, error) {
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

func TestEnvironments(t *testing.T) {
	fs := osfs.New("testdata")

	envs, err := Environments(fs)
	if err != nil {
		t.Fatal(err)
	}

	want := []Target{
		{"prod", "eu-west-1"},
		{"prod", "us-east-1"},
		{"staging", "eu-west-1"},
		{"test", "us-east-1"},
	}

	if diff := cmp.Diff(want, envs); diff != "" {
		t.Fatalf("failed to parse environments:\n%s", diff)
	}
}
