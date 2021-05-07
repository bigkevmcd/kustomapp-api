package features

import (
	"context"
	"fmt"
	"io/fs"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/cucumber/godog"
	"github.com/cucumber/messages-go/v10"
	"github.com/google/go-cmp/cmp"
)

type kustomFeature struct {
	root       string
	lastOutput string
}

func (k *kustomFeature) aTemporaryDirectoryInTheEnvironment() error {
	dir, err := ioutil.TempDir("", "kustom")
	if err != nil {
		return fmt.Errorf("failed to create a temporary directory: %w", err)
	}
	k.root = dir
	return nil
}

func (k *kustomFeature) iRunSuccessfully(cmdArg string) error {
	args := strings.Split(cmdArg, " ")
	for i := range args {
		if args[i] == "KAPP_TEMP" {
			args[i] = k.root
			continue
		}
	}

	log.Printf("KEVIN!!!! passing through %#v\n", args)

	cmd := exec.CommandContext(context.TODO(), args[0], args[1:]...)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("failed to exec command %q: %w", cmdArg, err)
	}
	k.lastOutput = strings.TrimSpace(string(output))
	return nil
}

func (k *kustomFeature) iShouldGetTheMessage(arg1 string) error {
	if diff := cmp.Diff(arg1, k.lastOutput); diff != "" {
		return fmt.Errorf("failed to match output: %s", diff)
	}
	return nil
}

func (k *kustomFeature) cleanup(sc *godog.Scenario, err error) {
	if err := os.RemoveAll(k.root); err != nil {
		log.Fatalf("failed to remove the directory %s: %s", k.root, err)
	}
}

func (k *kustomFeature) aTreeOfFilesShouldBeGenerated(want *messages.PickleStepArgument_PickleDocString) error {
	parsed := strings.Split(want.Content, "\n")
	var dirs []string
	err := filepath.WalkDir(k.root, func(p string, d fs.DirEntry, err error) error {
		if err != nil {
			return fmt.Errorf("failed to WalkDir %s: %w", p, err)
		}
		if d.IsDir() {
			return nil
		}
		dirs = append(dirs, strings.TrimPrefix(p, k.root+"/"))
		return nil
	})

	if err != nil {
		return fmt.Errorf("failed to list directories in %s: %w", k.root, err)
	}

	if diff := cmp.Diff(parsed, dirs); diff != "" {
		return fmt.Errorf("failed to match output: %s", diff)
	}
	return nil
}

func InitializeScenario(ctx *godog.ScenarioContext) {
	kf := &kustomFeature{}
	ctx.AfterScenario(kf.cleanup)

	ctx.Step(`^I run successfully "([^"]*)"$`, kf.iRunSuccessfully)
	ctx.Step(`^a temporary directory in the environment$`, kf.aTemporaryDirectoryInTheEnvironment)
	ctx.Step(`^I should get the message "([^"]*)"$`, kf.iShouldGetTheMessage)
	ctx.Step(`^a tree of files should be generated$`, kf.aTreeOfFilesShouldBeGenerated)
}
