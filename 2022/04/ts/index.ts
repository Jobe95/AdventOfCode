import { MathHelpers, ParseHelpers } from "@aoc/ts";
import { readFileSync } from "fs";
import { dirname, join } from "path";
import { fileURLToPath } from "url";

const __filename = fileURLToPath(import.meta.url);
const __dirname = dirname(__filename);

const readInput = (filename: string): string => {
  return readFileSync(join(__dirname, "..", filename), "utf-8").trim();
};

const parseRanges = (line: string) => {
  const [first, second] = line
    .split(",")
    .map((pair) => pair.split("-").map(Number));
  return [
    MathHelpers.rangeInclusive(first[0], first[1]),
    MathHelpers.rangeInclusive(second[0], second[1]),
  ];
};

const containsAll = (arr1: unknown[], arr2: unknown[]) => {
  const set2 = new Set(arr2);
  return arr1.every((item) => set2.has(item));
};

const containsAny = (arr1: unknown[], arr2: unknown[]) => {
  const set2 = new Set(arr2);
  return arr1.some((item) => set2.has(item));
};

const partOne = (input: string): number => {
  const lines = ParseHelpers.lines(input);
  let sum = 0;

  for (const line of lines) {
    const [first, second] = parseRanges(line);

    if (containsAll(first, second) || containsAll(second, first)) {
      sum++;
    }
  }

  return sum;
};

const partTwo = (input: string): number => {
  const lines = ParseHelpers.lines(input);
  let sum = 0;

  for (const line of lines) {
    const [first, second] = parseRanges(line);

    if (containsAny(first, second)) {
      sum++;
    }
  }
  return sum;
};

const main = () => {
  const input = readInput("input.txt");
  console.log("Part 1:", partOne(input));
  console.log("Part 2:", partTwo(input));
};

main();
