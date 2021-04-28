package features

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/cucumber/godog"
	"github.com/google/go-cmp/cmp"
)

type kustomFeature struct {
	root       string
	lastOutput []byte
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
	arg := strings.Split(cmdArg, " ")
	cmd := exec.CommandContext(context.TODO(), arg[0], arg[1:]...)
	cmd.Env = append(os.Environ(),
		fmt.Sprintf("KAPP_TEMP=%s", k.root),
	)
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to exec command %q: %w", cmdArg, err)
	}
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("failed to exec command %q: %w", cmdArg, err)
	}
	k.lastOutput = output
	return nil
}

func (k *kustomFeature) iShouldGetTheMessage(arg1 string) error {
	if diff := cmp.Diff(arg1, string(k.lastOutput)); diff != "" {
		return fmt.Errorf("failed to match output: %s", diff)
	}
	return nil
}

func (k *kustomFeature) cleanup(sc *godog.Scenario, err error) {
	if err := os.RemoveAll(k.root); err != nil {
		log.Fatalf("failed to remove the directory %s: %s", k.root, err)
	}
}

func InitializeScenario(ctx *godog.ScenarioContext) {
	kf := &kustomFeature{}
	ctx.AfterScenario(kf.cleanup)

	ctx.Step(`^I run successfully "([^"]*)"$`, kf.iRunSuccessfully)
	ctx.Step(`^a temporary directory in the environment$`, kf.aTemporaryDirectoryInTheEnvironment)
	ctx.Step(`^I should get the message "([^"]*)"$`, kf.iShouldGetTheMessage)
}
