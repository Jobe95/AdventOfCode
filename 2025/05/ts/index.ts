import { MathHelpers, ParseHelpers } from "@aoc/ts";
import { readFileSync } from "fs";
import { dirname, join } from "path";
import { fileURLToPath } from "url";

const __filename = fileURLToPath(import.meta.url);
const __dirname = dirname(__filename);

type Range = { start: number; end: number };

const readInput = (filename: string): string => {
  return readFileSync(join(__dirname, "..", filename), "utf-8").trim();
};

const partOne = (input: string): number => {
  const lines = ParseHelpers.lines(input).filter(String);
  const ranges: Range[] = [];
  const ingredients = [];
  let sum = 0;

  for (const line of lines) {
    if (line.includes("-")) {
      ranges.push({
        start: Number(line.split("-")[0]),
        end: Number(line.split("-")[1]),
      });
    } else {
      ingredients.push(Number(line));
    }
  }

  ingredients.forEach((id) => {
    if (
      ranges.some((range) =>
        MathHelpers.inRangeInclusive(id, range.start, range.end),
      )
    ) {
      sum++;
    }
  });
  return sum;
};

const partTwo = (input: string): number => {
  const lines = ParseHelpers.lines(input).filter(String);
  const ranges: Range[] = [];
  let sum = 0;

  let current: Range = { start: 0, end: 0 };

  for (const line of lines) {
    if (line.includes("-")) {
      ranges.push({
        start: Number(line.split("-")[0]),
        end: Number(line.split("-")[1]),
      });
    }
  }

  const sorted = ranges.sort((a, b) => a.start - b.start);

  const merged: Range[] = [];
  for (let i = 0; i < sorted.length; i++) {
    const next = sorted[i];
    if (next.start <= current.end + 1) {
      current.end = Math.max(current.end, next.end);
    } else {
      merged.push(next);
      current = next;
    }
  }

  for (const range of merged) {
    sum += range.end - range.start + 1;
  }
  return sum;
};

const main = () => {
  const input = readInput("example.txt");
  console.log("Part 1:", partOne(input));
  console.log("Part 2:", partTwo(input));
};

main();
