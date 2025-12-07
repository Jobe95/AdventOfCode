import { Coord, Grid, GridHelpers, ParseHelpers } from "@aoc/ts";
import { readFileSync } from "fs";
import { dirname, join } from "path";
import { fileURLToPath } from "url";

const __filename = fileURLToPath(import.meta.url);
const __dirname = dirname(__filename);

const readInput = (filename: string): string => {
  return readFileSync(join(__dirname, "..", filename), "utf-8").trim();
};

type GridValue = "S" | "." | "^" | "|";

const moveBeamDown = (
  grid: Grid<GridValue>,
  coord: Coord,
): { nextCoord: Coord; hitSplitter: boolean } | undefined => {
  let value = GridHelpers.getAt(grid, coord);

  while (value !== "^") {
    GridHelpers.setAt(grid, coord, "|");

    coord = GridHelpers.move(coord, "down");

    if (!GridHelpers.inBounds(grid, coord.x, coord.y)) {
      return undefined;
    }

    value = GridHelpers.getAt(grid, coord);
  }

  return { nextCoord: coord, hitSplitter: value === "^" };
};

const partOne = (input: string): number => {
  const lines = ParseHelpers.lines(input);
  let grid = GridHelpers.grid<GridValue>(lines);
  let sum = 0;

  const startCoord = GridHelpers.findInGrid(
    grid,
    (val) => val === "S",
  ) as Coord;

  const queue: Coord[] = [GridHelpers.move(startCoord, "down")];
  const visited = new Set<string>(GridHelpers.coordToKey(startCoord));

  while (queue.length > 0) {
    const coord = queue.shift()!;

    const result = moveBeamDown(grid, coord);

    if (result && result.hitSplitter) {
      const key = GridHelpers.coordToKey(result.nextCoord);

      if (visited.has(key)) {
        continue;
      }
      visited.add(key);
      const left = GridHelpers.move(result.nextCoord, "left");
      const right = GridHelpers.move(result.nextCoord, "right");
      queue.push(left, right);
      sum++;
    }
  }

  return sum;
};

const partTwo = (input: string): number => {
  const lines = ParseHelpers.lines(input);
  const grid = GridHelpers.grid<GridValue>(lines);

  const startCoord = GridHelpers.findInGrid(
    grid,
    (val) => val === "S",
  ) as Coord;

  const memo = new Map<string, number>();
  let sum = 0;

  const countTimelines = (coord: Coord): number => {
    const key = GridHelpers.coordToKey(coord);
    if (memo.has(key)) {
      return memo.get(key)!;
    }

    const result = moveBeamDown(grid, coord);

    if (!result) {
      memo.set(key, 1);
      return 1;
    }

    if (result.hitSplitter) {
      const left = GridHelpers.move(result.nextCoord, "left");
      const right = GridHelpers.move(result.nextCoord, "right");

      const total = countTimelines(left) + countTimelines(right);

      memo.set(key, total);
      return total;
    }
    return 0;
  };

  sum = countTimelines(GridHelpers.move(startCoord, "down"));

  return sum;
};

const main = () => {
  const input = readInput("input.txt");
  console.log("Part 1:", partOne(input));
  console.log("Part 2:", partTwo(input));
};

main();
