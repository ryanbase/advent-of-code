import { readFileFromArgs } from "../../utils/read-file";

const input = await readFileFromArgs();

const grid = input.split("\n").map((row) => row.split(""));

function checkDirections(grid: string[][], i: number, j: number) {
  const row = grid[i];
  if (!row || row[j] !== "@") {
    return 0;
  }

  const dirs = [
    [-1, -1],
    [-1, 0],
    [-1, 1],
    [0, -1],
    [0, 1],
    [1, -1],
    [1, 0],
    [1, 1],
  ];
  let count = 0;
  dirs.forEach((dir) => {
    const i1 = dir[0]!;
    const j1 = dir[1]!;
    const row = grid[i + i1];
    if (!row) {
      return;
    }
    if (row[j + j1] === "@") {
      count++;
    }
    if (count > 3) {
      return;
    }
  });
  return count < 4 ? 1 : 0;
}

function part1() {
  let result = 0;

  grid.forEach((row, i) => {
    for (let j = 0; j < row.length; j++) {
      result += checkDirections(grid, i, j);
    }
  });

  return result;
}

function part2() {
  let result = 0;

  let found = true;
  while (found) {
    found = false;

    for (let i = 0; i < grid.length; i++) {
      for (let j = 0; j < grid[i]!.length; j++) {
        const removed = checkDirections(grid, i, j);
        if (removed === 0) {
          continue;
        }
        result++;
        found = true;
        grid[i]![j] = ".";
      }
    }
  }

  return result;
}

console.log(part1());
console.log(part2());
