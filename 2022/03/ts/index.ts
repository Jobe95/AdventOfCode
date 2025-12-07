import { ParseHelpers } from "@aoc/ts";
import { readFileSync } from "fs";
import { dirname, join } from "path";
import { fileURLToPath } from "url";

const __filename = fileURLToPath(import.meta.url);
const __dirname = dirname(__filename);

const readInput = (filename: string): string => {
  return readFileSync(join(__dirname, "..", filename), "utf-8").trim();
};

const buildPriorityMap = () => {
  const map: Record<string, number> = {};
  for (let i = 0; i < 26; i++) {
    map[String.fromCharCode(97 + i)] = i + 1;
    map[String.fromCharCode(65 + i)] = i + 27;
  }
  return map;
};

const partOne = (input: string): number => {
  const lines = ParseHelpers.lines(input);
  const priorityMap = buildPriorityMap();

  let sum = 0;
  const matches = [];
  for (const line of lines) {
    const first = new Set(line.slice(0, line.length / 2));
    const second = new Set(line.slice(line.length / 2));

    for (const char of first) {
      if (second.has(char)) {
        matches.push(priorityMap[char]);
      }
    }
  }

  sum = matches.reduce((acc, curr) => acc + curr, 0);

  return sum;
};

const partTwo = (input: string): number => {
  const lines = ParseHelpers.lines(input);
  const priorityMap = buildPriorityMap();

  let sum = 0;
  const matches = [];
  for (let i = 0; i < lines.length; i = i += 3) {
    const group = [lines[i], lines[i + 1], lines[i + 2]];
    const [a, b, c] = group.map((str) => new Set(str));
    for (const char of a) {
      if (b.has(char) && c.has(char)) {
        matches.push(priorityMap[char]);
      }
    }
  }

  sum = matches.reduce((acc, curr) => acc + curr, 0);

  return sum;
};

const main = () => {
  const input = readInput("input.txt");
  console.log("Part 1:", partOne(input));
  console.log("Part 2:", partTwo(input));
};

main();
