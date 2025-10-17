# 🎄 Advent of Code Setup Tool

A CLI tool to streamline your Advent of Code workflow across multiple programming languages.

## Installation

```bash
./tools/setup-cli/install.sh
```

This will:

- Build the `aoc` binary
- Install it to `~/.local/bin/aoc`
- Check if the directory is in your PATH

If needed, add to your `~/.zshrc`:

```bash
export PATH="$HOME/.local/bin:$PATH"
```

## Quick Start

### 1. Configure Session

```bash
aoc config set-session <your-session-cookie>
```

Get your session cookie from [adventofcode.com](https://adventofcode.com) (DevTools → Application → Cookies → session)

### 2. Create a Day

```bash
aoc setup
```

Select year, day, and language(s). Done!

## Example

```bash
$ aoc setup
# Select: 2024 → 01 → go, ts

$ tree 2024/01/
2024/01/
├── go
│   ├── input.txt
│   ├── example.txt
│   └── main.go
└── ts
    ├── input.txt
    ├── example.txt
    ├── index.ts
    └── tsconfig.json
```

## Commands

```bash
aoc setup                      # Interactive setup
aoc config set-session <key>   # Save session cookie
aoc config show-session        # Show current session
```

## Language Support

**Go** - `main.go` with `partOne()` and `partTwo()` functions  
**TypeScript** - `index.ts` with file reading utilities

### Adding More Languages

Edit `tools/setup-cli/internal/templates/templates.go`:

1. Add case in `GenerateLanguageFiles()`
2. Create generation function (e.g., `generatePythonFiles()`)
3. Update language list in `internal/ui/interactive.go`

## License

MIT
