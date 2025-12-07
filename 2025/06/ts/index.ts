import { Grid, GridHelpers, ParseHelpers } from "@aoc/ts";
import { readFileSync } from "fs";
import { dirname, join } from "path";
import { fileURLToPath } from "url";

const __filename = fileURLToPath(import.meta.url);
const __dirname = dirname(__filename);

type Operator = "*" | "+";
type GridValue = string | Operator | " ";
type Problem = { numbers: number[]; operator: Operator };

const readInput = (filename: string): string => {
  return readFileSync(join(__dirname, "..", filename), "utf-8").trim();
};

const getProblems = (grid: Grid<GridValue>): Problem[] => {
  const dimension = GridHelpers.gridDimensions(grid);
  const operatorCoords = GridHelpers.findAllInGrid(
    grid,
    (value) => value === "*" || value === "+",
  );

  return operatorCoords.map((coord, index) => {
    const startX = coord.x;
    const endX = operatorCoords[index + 1]?.x ?? dimension.width;
    const numbers: number[] = [];

    for (let row = 0; row < coord.y; row++) {
      const digits: string[] = [];
      for (let col = startX; col < endX; col++) {
        const cell = grid[row][col];
        if (cell.trim() !== "" && !isNaN(Number(cell))) {
          digits.push(cell);
        }
      }
      if (digits.length > 0) {
        numbers.push(parseInt(digits.join("")));
      }
    }

    const operator = GridHelpers.getAt(grid, coord);

    return { numbers, operator: operator as Operator };
  });
};

const getProblemsPartTwo = (grid: Grid<GridValue>): Problem[] => {
  const dimension = GridHelpers.gridDimensions(grid);
  const operatorCoords = GridHelpers.findAllInGrid(
    grid,
    (value) => value === "*" || value === "+",
  );

  operatorCoords.reverse();

  return operatorCoords.map((coord, index) => {
    const startY =
      index === 0 ? dimension.height - 1 : operatorCoords[index - 1].y - 2;
    const endY = coord.y;
    const numbers: number[] = [];

    for (let row = startY; row >= endY; row--) {
      const digits: string[] = [];
      for (let col = 0; col < coord.x; col++) {
        const cell = grid[row][col];
        if (cell.trim() !== "" && !isNaN(Number(cell))) {
          digits.push(cell);
        }
      }
      if (digits.length > 0) {
        numbers.push(parseInt(digits.join("")));
      }
    }

    const operator = GridHelpers.getAt(grid, coord);

    return { numbers, operator: operator as Operator };
  });
};

const partOne = (input: string): number => {
  const lines = ParseHelpers.lines(input);
  let sum = 0;

  const grid = GridHelpers.grid<GridValue>(lines);

  const problems = getProblems(grid);

  problems.forEach((problem) => {
    sum += problem.numbers.reduce(
      (acc, curr) => {
        if (problem.operator === "*") {
          return acc * curr;
        } else if (problem.operator === "+") {
          return acc + curr;
        }
        return acc;
      },
      problem.operator === "*" ? 1 : 0,
    );
  });

  return sum;
};

const partTwo = (input: string): number => {
  const lines = ParseHelpers.lines(input);
  let sum = 0;

  const grid = GridHelpers.grid<GridValue>(lines);
  const transposed = GridHelpers.transpose(grid);

  const problems = getProblemsPartTwo(transposed);

  problems.forEach((problem) => {
    sum += problem.numbers.reduce(
      (acc, curr) => {
        if (problem.operator === "*") {
          return acc * curr;
        } else if (problem.operator === "+") {
          return acc + curr;
        }
        return acc;
      },
      problem.operator === "*" ? 1 : 0,
    );
  });

  return sum;
};

const main = () => {
  const input = readInput("input.txt");
  console.log("Part 1:", partOne(input));
  console.log("Part 2:", partTwo(input));
};

main();
