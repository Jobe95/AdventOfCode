import { ParseHelpers } from "./parse";

export type Position<T> = { x: number; y: number; value: T };
export type Coord = { x: number; y: number };
export type Grid<T> = T[][];

export type Direction =
  | "up"
  | "down"
  | "left"
  | "right"
  | "upLeft"
  | "upRight"
  | "downLeft"
  | "downRight";

export type CardinalDirection = "up" | "down" | "left" | "right";

const DIRECTION_DELTAS: Record<Direction, Coord> = {
  up: { x: 0, y: -1 },
  down: { x: 0, y: 1 },
  left: { x: -1, y: 0 },
  right: { x: 1, y: 0 },
  upLeft: { x: -1, y: -1 },
  upRight: { x: 1, y: -1 },
  downLeft: { x: -1, y: 1 },
  downRight: { x: 1, y: 1 },
};

const CARDINAL_DIRECTIONS: CardinalDirection[] = [
  "up",
  "down",
  "left",
  "right",
];
const ALL_DIRECTIONS: Direction[] = [
  "up",
  "down",
  "left",
  "right",
  "upLeft",
  "upRight",
  "downLeft",
  "downRight",
];

type GridLoggerOptions = {
  title?: string;
  cellWidth?: number;
  showCoordinates?: boolean;
  highlightPositions?: Position<unknown>[];
  highlightColor?: string;
};

const COLORS = {
  reset: "\x1b[0m",
  dim: "\x1b[2m",
  red: "\x1b[31m",
  green: "\x1b[32m",
  yellow: "\x1b[33m",
  blue: "\x1b[34m",
  magenta: "\x1b[35m",
  cyan: "\x1b[36m",
  bgRed: "\x1b[41m",
  bgGreen: "\x1b[42m",
  bgYellow: "\x1b[43m",
  bgBlue: "\x1b[44m",
};

const logGrid = <T>(grid: Grid<T>, options: GridLoggerOptions = {}): void => {
  const {
    title,
    cellWidth = 3,
    showCoordinates = true,
    highlightPositions = [],
    highlightColor = "bgYellow",
  } = options;

  const rows = grid.length;
  const cols = grid[0]?.length ?? 0;

  const highlightSet = new Set(
    highlightPositions.map(({ x, y }) => `${x},${y}`),
  );

  const pad = (val: string, width: number): string => {
    return val.padStart(Math.floor((width + val.length) / 2)).padEnd(width);
  };

  const horizontalLine = (left: string, mid: string, right: string): string => {
    const inner = Array(cols).fill("─".repeat(cellWidth)).join(mid);
    return `${left}${inner}${right}`;
  };

  console.log("");

  if (title) {
    const titleLine = `${COLORS.cyan}┌${"─".repeat(title.length + 2)}┐${COLORS.reset}`;
    const titleContent = `${COLORS.cyan}│${COLORS.reset} ${COLORS.yellow}${title}${COLORS.reset} ${COLORS.cyan}│${COLORS.reset}`;
    const titleBottom = `${COLORS.cyan}└${"─".repeat(title.length + 2)}┘${COLORS.reset}`;
    console.log(titleLine);
    console.log(titleContent);
    console.log(titleBottom);
    console.log("");
  }

  if (showCoordinates) {
    const header = Array.from({ length: cols }, (_, i) =>
      pad(i.toString(), cellWidth),
    ).join(" ");
    console.log(`${COLORS.dim}    ${header}${COLORS.reset}`);
  }

  console.log(
    `${COLORS.dim}   ${horizontalLine("┌", "┬", "┐")}${COLORS.reset}`,
  );

  grid.forEach((row, rowIdx) => {
    const rowLabel = showCoordinates
      ? `${COLORS.dim}${pad(rowIdx.toString(), 2)} ${COLORS.reset}`
      : "   ";

    const cells = row.map((cell, colIdx) => {
      const cellStr = String(cell);
      const paddedCell = pad(cellStr, cellWidth);
      const isHighlighted = highlightSet.has(`${rowIdx},${colIdx}`);

      if (isHighlighted) {
        const color =
          COLORS[highlightColor as keyof typeof COLORS] ?? COLORS.bgYellow;
        return `${color}${paddedCell}${COLORS.reset}`;
      }

      return paddedCell;
    });

    console.log(
      `${rowLabel}${COLORS.dim}│${COLORS.reset}${cells.join(`${COLORS.dim}│${COLORS.reset}`)}${COLORS.dim}│${COLORS.reset}`,
    );

    if (rowIdx < rows - 1) {
      console.log(
        `${COLORS.dim}   ${horizontalLine("├", "┼", "┤")}${COLORS.reset}`,
      );
    }
  });

  console.log(
    `${COLORS.dim}   ${horizontalLine("└", "┴", "┘")}${COLORS.reset}`,
  );
  console.log(`${COLORS.dim}   Grid size: ${rows}×${cols}${COLORS.reset}`);
  console.log("");
};

const grid = <T = string>(
  input: string[],
  transform?: (char: string, x: number, y: number) => T,
): Grid<T> => {
  return input.map((line, y) =>
    ParseHelpers.chars(line).map((char, x) =>
      transform ? transform(char, x, y) : (char as unknown as T),
    ),
  );
};

const intGrid = (input: string[]): Grid<number> => {
  return grid(input, (char) => Number(char));
};

const gridDimensions = <T>(g: Grid<T>): { width: number; height: number } => {
  return {
    height: g.length,
    width: g[0]?.length ?? 0,
  };
};

const transpose = <T>(g: Grid<T>): Grid<T> => {
  if (g.length === 0) return [];
  const { width, height } = gridDimensions(g);
  const result: Grid<T> = [];
  for (let x = 0; x < width; x++) {
    const row: T[] = [];
    for (let y = 0; y < height; y++) {
      row.push(g[y]![x]!);
    }
    result.push(row);
  }
  return result;
};

const inBounds = <T>(g: Grid<T>, x: number, y: number): boolean => {
  const { width, height } = gridDimensions(g);
  return x >= 0 && y >= 0 && x < width && y < height;
};

const move = (coord: Coord, direction: Direction): Coord => {
  const delta = DIRECTION_DELTAS[direction];
  return { x: coord.x + delta.x, y: coord.y + delta.y };
};

const moveN = (coord: Coord, direction: Direction, steps: number): Coord => {
  const delta = DIRECTION_DELTAS[direction];
  return { x: coord.x + delta.x * steps, y: coord.y + delta.y * steps };
};

const getAt = <T>(g: Grid<T>, coord: Coord): T | undefined => {
  if (!inBounds(g, coord.x, coord.y)) return undefined;
  return g[coord.y]![coord.x];
};

const setAt = <T>(g: Grid<T>, coord: Coord, value: T): void => {
  if (inBounds(g, coord.x, coord.y)) {
    g[coord.y]![coord.x] = value;
  }
};

const getAdjacent = <T>(
  g: Grid<T>,
  x: number,
  y: number,
  includeDiagonals = true,
): Position<T>[] => {
  const directions = includeDiagonals ? ALL_DIRECTIONS : CARDINAL_DIRECTIONS;
  const adjacents: Position<T>[] = [];

  for (const dir of directions) {
    const delta = DIRECTION_DELTAS[dir];
    const newX = x + delta.x;
    const newY = y + delta.y;

    if (inBounds(g, newX, newY)) {
      adjacents.push({ x: newX, y: newY, value: g[newY]![newX]! });
    }
  }

  return adjacents;
};

const findInGrid = <T>(
  g: Grid<T>,
  predicate: (value: T) => boolean,
): Coord | undefined => {
  for (let y = 0; y < g.length; y++) {
    for (let x = 0; x < g[y]!.length; x++) {
      if (predicate(g[y]![x]!)) {
        return { x, y };
      }
    }
  }
  return undefined;
};

const findAllInGrid = <T>(
  g: Grid<T>,
  predicate: (value: T) => boolean,
): Coord[] => {
  const results: Coord[] = [];
  for (let y = 0; y < g.length; y++) {
    for (let x = 0; x < g[y]!.length; x++) {
      if (predicate(g[y]![x]!)) {
        results.push({ x, y });
      }
    }
  }
  return results;
};

const coordToKey = (coord: Coord): string => `${coord.x},${coord.y}`;

const keyToCoord = (key: string): Coord => {
  const [x, y] = key.split(",").map(Number);
  return { x: x!, y: y! };
};

export const GridHelpers = {
  logGrid,
  grid,
  intGrid,
  gridDimensions,
  transpose,
  inBounds,
  move,
  moveN,
  getAt,
  setAt,
  getAdjacent,
  findInGrid,
  findAllInGrid,
  coordToKey,
  keyToCoord,
  DIRECTION_DELTAS,
  CARDINAL_DIRECTIONS,
  ALL_DIRECTIONS,
};
