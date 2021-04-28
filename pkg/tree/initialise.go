package tree

import (
	"fmt"
	"os"
	"path/filepath"

	kustypes "sigs.k8s.io/kustomize/api/types"
	"sigs.k8s.io/yaml"
)

// InitialiseDirectory creates a new directory structure for deploying from
// environments.
func InitialiseDirectory(root string, envs []string) error {
	return initialBases(root)
}

func initialBases(root string) error {
	kus := kustypes.Kustomization{
		TypeMeta: kustypes.TypeMeta{
			APIVersion: kustypes.KustomizationVersion,
			Kind:       kustypes.KustomizationKind,
		},
	}
	return marshalToFile(filepath.Join(root, "bases/kustomization.yaml"), kus)
}

func marshalToFiles(m map[string]interface{}) error {
	for k, v := range m {
		if err := marshalToFile(k, v); err != nil {
			return err
		}
	}
	return nil
}

func marshalToFile(filename string, v interface{}) error {
	b, err := yaml.Marshal(v)
	if err != nil {
		return fmt.Errorf("failed to marshal %#v for %s: %w", v, filename, err)
	}
	if err := os.MkdirAll(filepath.Dir(filename), 0755); err != nil {
		return fmt.Errorf("failed to MkdirAll(%s): %w", filename, err)
	}
	if err := os.WriteFile(filename, b, 0644); err != nil {
		return fmt.Errorf("failed to WriteFile to %s: %w", filename, err)
	}
	return nil
}
