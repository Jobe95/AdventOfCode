package setup

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/Jobe95/AdventOfCode/internal/fetcher"
	"github.com/Jobe95/AdventOfCode/internal/templates"
	"github.com/Jobe95/AdventOfCode/internal/ui"
)

func Execute(selections *ui.Selections) error {
	projectRoot, err := findProjectRoot()
	if err != nil {
		return err
	}

	input, err := fetcher.FetchInput(selections.Year, selections.Day)
	if err != nil {
		return fmt.Errorf("failed to fetch input: %w", err)
	}

	example, _ := fetcher.FetchExample(selections.Year, selections.Day)

	dayPadded := fmt.Sprintf("%02d", selections.Day)

	for _, lang := range selections.Languages {
		langDir := filepath.Join(projectRoot, fmt.Sprintf("%d", selections.Year), dayPadded, lang)

		if err := os.MkdirAll(langDir, 0755); err != nil {
			return err
		}

		if err := os.WriteFile(filepath.Join(langDir, "input.txt"), []byte(input), 0644); err != nil {
			return err
		}

		if err := os.WriteFile(filepath.Join(langDir, "example.txt"), []byte(example), 0644); err != nil {
			return err
		}

		if err := templates.GenerateLanguageFiles(lang, langDir, selections.Year, selections.Day); err != nil {
			return err
		}
	}

	fmt.Printf("\n✓ Setup complete: %d/%02d (%s)\n", selections.Year, selections.Day, strings.Join(selections.Languages, ", "))
	return nil
}

func findProjectRoot() (string, error) {
	dir, err := os.Getwd()
	if err != nil {
		return "", err
	}

	for {
		if _, err := os.Stat(filepath.Join(dir, "2024")); err == nil {
			return dir, nil
		}

		parent := filepath.Dir(dir)
		if parent == dir {
			return "", fmt.Errorf("project root not found")
		}
		dir = parent
	}
}
