export type Point = { x: number; y: number };

const gcd = (a: number, b: number): number => {
  a = Math.abs(a);
  b = Math.abs(b);
  while (b !== 0) {
    const t = b;
    b = a % b;
    a = t;
  }
  return a;
};

const lcm = (a: number, b: number): number => {
  return Math.abs(a * b) / gcd(a, b);
};

const gcdMultiple = (numbers: number[]): number => {
  return numbers.reduce((acc, n) => gcd(acc, n), numbers[0] ?? 1);
};

const lcmMultiple = (numbers: number[]): number => {
  return numbers.reduce((acc, n) => lcm(acc, n), 1);
};

const sum = (numbers: number[]): number => {
  return numbers.reduce((acc, n) => acc + n, 0);
};

const product = (numbers: number[]): number => {
  return numbers.reduce((acc, n) => acc * n, 1);
};

const min = (numbers: number[]): number => {
  return Math.min(...numbers);
};

const max = (numbers: number[]): number => {
  return Math.max(...numbers);
};

const clamp = (value: number, minVal: number, maxVal: number): number => {
  return Math.max(minVal, Math.min(maxVal, value));
};

const mod = (n: number, m: number): number => {
  return ((n % m) + m) % m;
};

const manhattanDistance = (a: Point, b: Point): number => {
  return Math.abs(a.x - b.x) + Math.abs(a.y - b.y);
};

const euclideanDistance = (a: Point, b: Point): number => {
  return Math.sqrt((a.x - b.x) ** 2 + (a.y - b.y) ** 2);
};

const addPoints = (a: Point, b: Point): Point => {
  return { x: a.x + b.x, y: a.y + b.y };
};

const subtractPoints = (a: Point, b: Point): Point => {
  return { x: a.x - b.x, y: a.y - b.y };
};

const scalePoint = (p: Point, scalar: number): Point => {
  return { x: p.x * scalar, y: p.y * scalar };
};

const inBounds = (
  p: Point,
  minX: number,
  maxX: number,
  minY: number,
  maxY: number,
): boolean => {
  return p.x >= minX && p.x <= maxX && p.y >= minY && p.y <= maxY;
};

const range = (start: number, end: number, step = 1): number[] => {
  const result: number[] = [];
  if (step > 0) {
    for (let i = start; i < end; i += step) {
      result.push(i);
    }
  } else {
    for (let i = start; i > end; i += step) {
      result.push(i);
    }
  }
  return result;
};

const rangeInclusive = (start: number, end: number, step = 1): number[] => {
  const result: number[] = [];
  if (step > 0) {
    for (let i = start; i <= end; i += step) {
      result.push(i);
    }
  } else {
    for (let i = start; i >= end; i += step) {
      result.push(i);
    }
  }
  return result;
};

const inRange = (value: number, start: number, end: number) => {
  return value >= start && value < end;
};

const inRangeInclusive = (value: number, start: number, end: number) => {
  return value >= start && value <= end;
};

export const MathHelpers = {
  gcd,
  lcm,
  gcdMultiple,
  lcmMultiple,
  sum,
  product,
  min,
  max,
  clamp,
  mod,
  manhattanDistance,
  euclideanDistance,
  addPoints,
  subtractPoints,
  scalePoint,
  inBounds,
  range,
  rangeInclusive,
  inRange,
  inRangeInclusive,
};
