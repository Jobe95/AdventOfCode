import { MathHelpers, ParseHelpers, Point } from "@aoc/ts";
import { readFileSync } from "fs";
import { dirname, join } from "path";
import { fileURLToPath } from "url";

const __filename = fileURLToPath(import.meta.url);
const __dirname = dirname(__filename);

const readInput = (filename: string): string => {
  return readFileSync(join(__dirname, "..", filename), "utf-8").trim();
};

type DisjointSet = {
  roots: number[];
  sizes: number[];
  groups: number;
};

const createDisjointSet = (initialSize: number): DisjointSet => ({
  roots: Array.from({ length: initialSize }, (_, i) => i),
  sizes: Array(initialSize).fill(1),
  groups: initialSize,
});

// Find with path compression (returns new roots array)
const find = (ds: DisjointSet, index: number): [number, number[]] => {
  const { roots } = ds;
  if (roots[index] === index) return [index, roots];
  const [root, newRoots] = find(ds, roots[index]);
  const updatedRoots = [...newRoots];
  updatedRoots[index] = root;
  return [root, updatedRoots];
};

// Union returns a new DisjointSet
const union = (ds: DisjointSet, ia: number, ib: number): DisjointSet => {
  let [ra, rootsA] = find(ds, ia);
  let [rb, rootsB] = find({ ...ds, roots: rootsA }, ib);

  if (ra === rb) return { ...ds, roots: rootsB };

  let newRoots = [...rootsB];
  let newSizes = [...ds.sizes];

  if (newSizes[ra] < newSizes[rb]) {
    [ra, rb] = [rb, ra];
  }

  newRoots[rb] = ra;
  newSizes[ra] += newSizes[rb];

  return {
    roots: newRoots,
    sizes: newSizes,
    groups: ds.groups - 1,
  };
};

const partOne = (input: string): number => {
  const lines = ParseHelpers.lines(input);
  const boxes: Point[] = [];

  let sum = 0;

  for (const line of lines) {
    const [x, y, z] = line.split(",").map((val) => Number(val));
    const point = { x, y, z };
    boxes.push(point);
  }

  const pairs: { i: number; j: number; dist: number }[] = [];
  for (let i = 0; i < boxes.length; i++) {
    for (let j = i + 1; j < boxes.length; j++) {
      pairs.push({
        i,
        j,
        dist: MathHelpers.euclideanDistance(boxes[i], boxes[j]),
      });
    }
  }

  pairs.sort((a, b) => a.dist - b.dist);

  let ds = createDisjointSet(boxes.length);

  for (let k = 0; k < 1000; k++) {
    const { i, j } = pairs[k];
    ds = union(ds, i, j);
  }

  const sizes = ds.sizes.sort((a, b) => b - a);
  sum = MathHelpers.product(sizes.slice(0, 3));

  return sum;
};

const partTwo = (input: string): number => {
  const lines = ParseHelpers.lines(input);
  const boxes: Point[] = [];

  let sum = 0;

  for (const line of lines) {
    const [x, y, z] = line.split(",").map((val) => Number(val));
    const point = { x, y, z };
    boxes.push(point);
  }

  const pairs: { i: number; j: number; dist: number }[] = [];
  for (let i = 0; i < boxes.length; i++) {
    for (let j = i + 1; j < boxes.length; j++) {
      pairs.push({
        i,
        j,
        dist: MathHelpers.euclideanDistance(boxes[i], boxes[j]),
      });
    }
  }

  pairs.sort((a, b) => a.dist - b.dist);

  let ds = createDisjointSet(boxes.length);
  for (const { i, j } of pairs) {
    ds = union(ds, i, j);
    if (ds.groups === 1) {
      sum = boxes[i].x * boxes[j].x;
      break;
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
