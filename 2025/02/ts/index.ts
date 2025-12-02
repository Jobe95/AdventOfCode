import { readFileSync } from "fs";
import { dirname, join } from "path";
import { fileURLToPath } from "url";

const __filename = fileURLToPath(import.meta.url);
const __dirname = dirname(__filename);

const readInput = (filename: string): string => {
  return readFileSync(join(__dirname, "..", filename), "utf-8").trim();
};

const isInvalidIdPartOne = (id: number) => {
  const str = id.toString();
  const half = str.length / 2;
  return str.slice(0, half) === str.slice(half);
};

const isInvalidIdPartTwo = (id: number) => {
  const str = id.toString();
  const len = str.length;

  for (let i = 1; i < len; i++) {
    if (len % i !== 0) continue;
    const pattern = str.slice(0, i);
    const repeated = pattern.repeat(len / i);

    if (repeated === str) {
      return true;
    }
  }
  return false;
};

const partOne = (input: string): number => {
  const ranges = input.split(",");
  let sum = 0;
  for (const range of ranges) {
    const firstId = Number(range.split("-")[0]);
    const secondId = Number(range.split("-")[1]);
    for (let id = firstId; id <= secondId; id++) {
      if (isInvalidIdPartOne(id)) {
        sum += id;
      }
    }
  }

  return sum;
};

const partTwo = (input: string): number => {
  const ranges = input.split(",");
  let sum = 0;
  for (const range of ranges) {
    const firstId = Number(range.split("-")[0]);
    const secondId = Number(range.split("-")[1]);
    for (let id = firstId; id <= secondId; id++) {
      if (isInvalidIdPartTwo(id)) {
        sum += id;
      }
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
