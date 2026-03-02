package schema

import (
	"fmt"
	"path/filepath"

	schemapkg "github.com/losisin/helm-values-schema-json/pkg"
)

// Regenerate generates values.schema.json from the given values.yaml file.
func Regenerate(valuesPath, outputPath string) error {
	absValues, err := filepath.Abs(valuesPath)
	if err != nil {
		return fmt.Errorf("resolving values path: %w", err)
	}
	absOutput, err := filepath.Abs(outputPath)
	if err != nil {
		return fmt.Errorf("resolving output path: %w", err)
	}

	cfg := &schemapkg.Config{
		Input:      []string{absValues},
		OutputPath: absOutput,
		Draft:      2020,
		Indent:     4,
	}

	if err := schemapkg.GenerateJsonSchema(cfg); err != nil {
		return fmt.Errorf("generating schema: %w", err)
	}
	return nil
}
