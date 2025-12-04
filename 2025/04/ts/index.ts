import { GridHelpers, ParseHelpers } from "@aoc/ts";
import { readFileSync } from "fs";
import { dirname, join } from "path";
import { fileURLToPath } from "url";

const __filename = fileURLToPath(import.meta.url);
const __dirname = dirname(__filename);

type Value = "." | "@" | "x";

const readInput = (filename: string): string => {
  return readFileSync(join(__dirname, "..", filename), "utf-8").trim();
};

const partOne = (input: string): number => {
  const lines = ParseHelpers.lines(input);
  const grid = GridHelpers.grid<Value>(lines);

  const relevantPositions = GridHelpers.findAllInGrid(
    grid,
    (val) => val === "@",
  );

  let sum = 0;

  relevantPositions.map(({ x, y }) => {
    const adjacentRolls = GridHelpers.getAdjacent<Value>(
      grid,
      x,
      y,
      true,
    ).filter((pos) => pos.value === "@");
    if (adjacentRolls.length < 4) {
      sum++;
    }
  });

  return sum;
};

const partTwo = (input: string): number => {
  const lines = ParseHelpers.lines(input);
  const grid = GridHelpers.grid<Value>(lines);

  let relevantPositions = GridHelpers.findAllInGrid(grid, (val) => val === "@");

  let sum = 0;

  let toRemove: typeof relevantPositions = [];

  do {
    toRemove = relevantPositions.filter(({ x, y }) => {
      const adjacentRolls = GridHelpers.getAdjacent<Value>(
        grid,
        x,
        y,
        true,
      ).filter((pos) => pos.value === "@");
      return adjacentRolls.length < 4;
    });

    toRemove.forEach(({ x, y }) =>
      GridHelpers.setAt<Value>(grid, { x, y }, "x"),
    );
    relevantPositions = relevantPositions.filter(
      (pos) => !toRemove.some((r) => r.x === pos.x && r.y === pos.y),
    );
    sum += toRemove.length;
  } while (toRemove.length > 0);

  GridHelpers.logGrid(grid, { showCoordinates: true });

  return sum;
};

const main = () => {
  const input = readInput("example.txt");
  console.log("Part 1:", partOne(input));
  console.log("Part 2:", partTwo(input));
};

main();
