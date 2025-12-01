package setup

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
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

	dayPadded := fmt.Sprintf("%02d", selections.Day)

	// Check if any directories already exist
	var existingDirs []string
	for _, lang := range selections.Languages {
		langDir := filepath.Join(projectRoot, fmt.Sprintf("%d", selections.Year), dayPadded, lang)
		if _, err := os.Stat(langDir); err == nil {
			existingDirs = append(existingDirs, langDir)
		}
	}

	// If directories exist, ask for confirmation
	if len(existingDirs) > 0 {
		fmt.Printf("\nThe following directories already exist:\n")
		for _, dir := range existingDirs {
			fmt.Printf("  - %s\n", dir)
		}
		fmt.Print("\nDo you want to override existing directories and files? (y/N): ")

		reader := bufio.NewReader(os.Stdin)
		response, err := reader.ReadString('\n')
		if err != nil {
			return fmt.Errorf("failed to read input: %w", err)
		}

		response = strings.TrimSpace(strings.ToLower(response))
		if response != "y" && response != "yes" {
			return fmt.Errorf("setup cancelled by user")
		}
		fmt.Println()
	}

	input, err := fetcher.FetchInput(selections.Year, selections.Day)
	if err != nil {
		return fmt.Errorf("failed to fetch input: %w", err)
	}

	example, _ := fetcher.FetchExample(selections.Year, selections.Day)

	// Create day directory and write shared input/example files
	dayDir := filepath.Join(projectRoot, fmt.Sprintf("%d", selections.Year), dayPadded)
	if err := os.MkdirAll(dayDir, 0755); err != nil {
		return err
	}

	if err := os.WriteFile(filepath.Join(dayDir, "input.txt"), []byte(input), 0644); err != nil {
		return err
	}

	if err := os.WriteFile(filepath.Join(dayDir, "example.txt"), []byte(example), 0644); err != nil {
		return err
	}

	// Create language-specific directories and files
	for _, lang := range selections.Languages {
		langDir := filepath.Join(dayDir, lang)

		if err := os.MkdirAll(langDir, 0755); err != nil {
			return err
		}

		if err := templates.GenerateLanguageFiles(lang, langDir, selections.Year, selections.Day); err != nil {
			return err
		}

		if lang == "ts" {
			if err := runPnpmInstall(langDir); err != nil {
				fmt.Printf("Warning: failed to run pnpm install: %v\n", err)
			} else {
				fmt.Printf("✓ Dependencies installed for TypeScript\n")
			}
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

func runPnpmInstall(dir string) error {
	cmd := exec.Command("pnpm", "install")
	cmd.Dir = dir
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
