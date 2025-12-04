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

import (
	"os"
	"path/filepath"
	"strings"
)

func readInput(filename string) string {
	data, err := os.ReadFile(filepath.Join("..", filename))
	if err != nil {
		panic(err)
	}
	return strings.TrimSpace(string(data))
}

func partOne(input string) int {
	return 0
}

func partTwo(input string) int {
	return 0
}

func main() {
	input := readInput("example.txt")
	println("Part 1:", partOne(input))
	println("Part 2:", partTwo(input))
}`

	return os.WriteFile(filepath.Join(targetDir, "main.go"), []byte(content), 0644)
}

func generateTsFiles(targetDir string) error {
	indexContent := `import { readFileSync } from 'fs';
import { fileURLToPath } from 'url';
import { dirname, join } from 'path';

const __filename = fileURLToPath(import.meta.url);
const __dirname = dirname(__filename);

const readInput = (filename: string): string => {
  return readFileSync(join(__dirname, '..', filename), 'utf-8').trim();
};

const partOne = (input: string): number => {
  return 0;
};

const partTwo = (input: string): number => {
  return 0;
};

const main = () => {
  const input = readInput('example.txt');
  console.log('Part 1:', partOne(input));
  console.log('Part 2:', partTwo(input));
};

main();
`

	tsconfigContent := `{
  "compilerOptions": {
    "target": "ES2022",
    "module": "ES2022",
    "lib": ["ES2022"],
    "moduleResolution": "bundler",
    "esModuleInterop": true,
    "skipLibCheck": true,
    "strict": true,
    "forceConsistentCasingInFileNames": true,
    "resolveJsonModule": true,
    "outDir": "./dist",
    "rootDir": "."
  },
  "include": ["*.ts"],
  "exclude": ["node_modules", "dist"]
}
`

	packageJsonContent := `{
  "name": "aoc-solution",
  "version": "1.0.0",
  "type": "module",
  "scripts": {
    "start": "tsx index.ts",
    "build": "tsc",
    "dev": "tsx watch index.ts"
  },
  "dependencies": {
    "@aoc/ts": "link:../../../lib/ts"
  },
  "devDependencies": {
    "@types/node": "^22.10.2",
    "tsx": "^4.19.2",
    "typescript": "^5.7.2"
  }
}
`

	if err := os.WriteFile(filepath.Join(targetDir, "index.ts"), []byte(indexContent), 0644); err != nil {
		return err
	}

	if err := os.WriteFile(filepath.Join(targetDir, "tsconfig.json"), []byte(tsconfigContent), 0644); err != nil {
		return err
	}

	if err := os.WriteFile(filepath.Join(targetDir, "package.json"), []byte(packageJsonContent), 0644); err != nil {
		return err
	}

	gitignoreContent := `node_modules/
dist/
*.log
.DS_Store
`

	return os.WriteFile(filepath.Join(targetDir, ".gitignore"), []byte(gitignoreContent), 0644)
}
