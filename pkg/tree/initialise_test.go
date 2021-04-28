package tree

import (
	"fmt"
	"io/fs"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestInitialiseDirectory(t *testing.T) {
	dir := mkTempDir(t)
	err := InitialiseDirectory(dir, []string{"staging"})
	if err != nil {
		t.Fatal(err)
	}
	want := []string{
		"bases/kustomization.yaml",
	}
	if diff := cmp.Diff(want, listTree(t, dir)); diff != "" {
		t.Fatalf("failed to generate files:\n%s", diff)
	}
}

func mkTempDir(t *testing.T) string {
	dir, err := ioutil.TempDir("", "peanut")
	if err != nil {
		t.Fatalf("failed to create TempDir: %s", err)
	}
	t.Cleanup(func() {
		if err := os.RemoveAll(dir); err != nil {
			t.Fatalf("failed to cleanup dir %s: %s", dir, err)
		}
	})
	return dir
}

func listTree(t *testing.T, root string) []string {
	var dirs []string
	err := filepath.WalkDir(root, func(p string, d fs.DirEntry, err error) error {
		if err != nil {
			return fmt.Errorf("failed to WalkDir %s: %w", p, err)
		}
		if d.IsDir() {
			return nil
		}
		dirs = append(dirs, strings.TrimPrefix(p, root+"/"))
		return nil
	})
	if err != nil {
		t.Fatal(err)
	}
	return dirs
}
