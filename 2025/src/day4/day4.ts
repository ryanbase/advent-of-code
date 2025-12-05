import { readFileFromArgs } from "../../utils/read-file";

const input = await readFileFromArgs();

const grid = input.split("\n");

const gridArray = grid.map((row) => row.split(""));

console.log(gridArray);

function checkDirections(grid: string[], i: number, j: number) {
  if (grid[i]?.charAt(j) !== "@") {
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
    if (grid[i + i1]?.charAt(j + j1) === "@") {
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

    for (let i = 0; i < gridArray.length; i++) {
      for (let j = 0; j < gridArray[i]!.length; j++) {
        // const removed = checkDirections(gridArray, i)
      }
    }

    // grid.forEach((row, i) => {
    //   for (let j = 0; j < row.length; j++) {
    //     const removed = checkDirections(grid, i, j);
    //     if (removed === 1) {
    //       result++;
    //       found = true;
    //       const array = row.split("");

    //       array[j] = ".";
    //       grid[i] = array.join();
    //       row = array.join();
    //     }
    //   }
    // });
  }

  return result;
}

console.log(part1());
console.log(part2());
