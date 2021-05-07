package tree

import (
	"path/filepath"
	"testing"

	"github.com/go-git/go-billy/v5"
	"github.com/go-git/go-billy/v5/osfs"
	"github.com/google/go-cmp/cmp"
	kustypes "sigs.k8s.io/kustomize/api/types"
)

func TestTargets(t *testing.T) {
	fs := osfs.New("testdata")

	envs, err := Targets(fs)
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

func AddTarget(fs billy.Filesystem, env, target string) error {
	return marshalToFile(filepath.Join(root, name, "bases/kustomization.yaml"), makeKustomization(func(k *kustypes.Kustomization) {
		k.Resources = []string{"../../bases"}
		k.CommonLabels = map[string]string{
			"com.bigkevmcd/environment": name,
		}
	}))
}

func TestAddTarget(t *testing.T) {
	dir := mkTempDir(t)
	err := InitialiseDirectory(dir, []string{"staging"})
	if err != nil {
		t.Fatal(err)
	}

	if err := AddTarget(osfs.New(dir), "staging", "staging-us"); err != nil {
		t.Fatal(err)
	}

	want := []string{
		"bases/kustomization.yaml",
		"staging/bases/kustomization.yaml",
		"staging/staging-us/kustomization.yaml",
	}
	if diff := cmp.Diff(want, listTree(t, dir)); diff != "" {
		t.Fatalf("failed to generate files:\n%s", diff)
	}
}
