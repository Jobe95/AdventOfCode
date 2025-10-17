package templates

import (
	"fmt"
	"os"
	"path/filepath"
)

func GenerateLanguageFiles(language, targetDir string, year, day int) error {
	switch language {
	case "go":
		return generateGoFiles(targetDir)
	case "ts":
		return generateTsFiles(targetDir)
	default:
		return fmt.Errorf("unsupported language: %s", language)
	}
}

func generateGoFiles(targetDir string) error {
	content := `package main

func main() {
	partOne()
	partTwo()
}

func partOne() int {
	return 0
}

func partTwo() int {
	return 0
}`

	return os.WriteFile(filepath.Join(targetDir, "main.go"), []byte(content), 0644)
}

func generateTsFiles(targetDir string) error {
	indexContent := `import * as fs from 'fs';
import * as path from 'path';

function readInput(filename: string): string {
    return fs.readFileSync(path.join(__dirname, filename), 'utf-8').trim();
}

function partOne(input: string): number {
    return 0;
}

function partTwo(input: string): number {
    return 0;
}

function main() {
    const input = readInput('input.txt');
    console.log('Part 1:', partOne(input));
    console.log('Part 2:', partTwo(input));
}

main();
`

	tsconfigContent := `{
  "compilerOptions": {
    "target": "ES2020",
    "module": "commonjs",
    "strict": true,
    "esModuleInterop": true
  }
}
`

	if err := os.WriteFile(filepath.Join(targetDir, "index.ts"), []byte(indexContent), 0644); err != nil {
		return err
	}

	return os.WriteFile(filepath.Join(targetDir, "tsconfig.json"), []byte(tsconfigContent), 0644)
}
