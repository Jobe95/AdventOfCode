import { readFileSync } from "fs";
import { dirname, join } from "path";
import { fileURLToPath } from "url";

const __filename = fileURLToPath(import.meta.url);
const __dirname = dirname(__filename);

const readInput = (filename: string): string => {
  return readFileSync(join(__dirname, "..", filename), "utf-8").trim();
};

const findLargestNumber = (line: string, max: number): number => {
  const len = line.length;
  const joltages: number[] = [];

  for (let i = 0; i < len; i++) {
    const digit = Number(line[i]);
    while (
      joltages.length > 0 &&
      digit > joltages[joltages.length - 1] &&
      joltages.length + (len - i - 1) >= max
    ) {
      joltages.pop();
    }
    if (joltages.length < max) {
      joltages.push(digit);
    }
  }

  return Number(joltages.join(""));
};

const partOne = (input: string): number => {
  const lines = input.split("\n");

  const sum = lines
    .map((line) => findLargestNumber(line, 2))
    .reduce((acc, curr) => acc + curr, 0);
  return sum;
};

const partTwo = (input: string): number => {
  const lines = input.split("\n");

  const sum = lines
    .map((line) => findLargestNumber(line, 12))
    .reduce((acc, curr) => acc + curr, 0);
  return sum;
};

const main = () => {
  const input = readInput("input.txt");
  console.log("Part 1:", partOne(input));
  console.log("Part 2:", partTwo(input));
};

main();
