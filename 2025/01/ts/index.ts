import { readFileSync } from "fs";
import { dirname, join } from "path";
import { fileURLToPath } from "url";

const __filename = fileURLToPath(import.meta.url);
const __dirname = dirname(__filename);

const MAX_BOUNDS = 100;
const STARTING_POSITION = 50;

enum Direction {
  LEFT = "L",
  RIGHT = "R",
}

type Instruction = {
  direction: Direction;
  value: number;
};

const readInput = (filename: string): string => {
  return readFileSync(join(__dirname, "..", filename), "utf-8").trim();
};

const parseInstruction = (line: string): Instruction => {
  const direction =
    line[0] === Direction.LEFT ? Direction.LEFT : Direction.RIGHT;
  const value = Number(line.slice(1));
  return { direction, value };
};

const wrapPosition = (position: number, max: number): number =>
  ((position % max) + max) % max;

const applyInstruction = (
  currentPosition: number,
  instruction: Instruction,
): number => {
  const delta =
    instruction.direction === Direction.LEFT
      ? -instruction.value
      : instruction.value;
  return currentPosition + delta;
};

const partOne = (input: string): number => {
  let position = STARTING_POSITION;
  let sum = 0;

  const instructions = input.split("\n").map(parseInstruction);

  for (const instruction of instructions) {
    position = applyInstruction(position, instruction);
    position = wrapPosition(position, MAX_BOUNDS);

    if (position === 0) {
      sum++;
    }
  }
  return sum;
};

const partTwo = (input: string): number => {
  let position = STARTING_POSITION;
  let sum = 0;

  const instructions = input.split("\n").map(parseInstruction);

  for (const instruction of instructions) {
    const previousPosition = position;
    position = applyInstruction(position, instruction);

    if (position <= 0) {
      sum += Math.floor(position / -MAX_BOUNDS) + 1;
      if (previousPosition === 0) {
        sum -= 1;
      }
    } else if (position >= MAX_BOUNDS) {
      sum += Math.floor(position / MAX_BOUNDS);
    }

    position = wrapPosition(position, MAX_BOUNDS);
  }

  return sum;
};

const main = () => {
  const input = readInput("input.txt");
  console.log("Part 1:", partOne(input));
  console.log("Part 2:", partTwo(input));
};

main();
