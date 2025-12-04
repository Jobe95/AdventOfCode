const lines = (input: string): string[] => {
  return input.split("\n");
};

const blocks = (input: string): string[] => {
  return input.split("\n\n");
};

const chars = (input: string): string[] => {
  return input.split("");
};

const digits = (input: string): number[] => {
  return input.split("").map(Number);
};

export const ParseHelpers = {
  lines,
  blocks,
  chars,
  digits,
};
