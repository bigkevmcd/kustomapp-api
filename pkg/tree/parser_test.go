package tree

import (
	"testing"

	"github.com/go-git/go-billy/v5/osfs"
	"github.com/google/go-cmp/cmp"
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
