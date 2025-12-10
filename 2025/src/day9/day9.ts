import { getFileName, readFileFromArgs } from "../../utils/read-file";

const input = await readFileFromArgs();

type Coord = {
  x: number;
  y: number;
};

const coords: Coord[] = input.split("\n").map((vals) => {
  const nums = vals.split(",");
  return { x: Number(nums[0]), y: Number(nums[1]) };
});

function part1() {
  let largest = 0;

  for (let i = 0; i < coords.length; i++) {
    for (let j = i + 1; j < coords.length; j++) {
      const p1 = coords[i]!;
      const p2 = coords[j]!;
      const area = (Math.abs(p1.x - p2.x) + 1) * (Math.abs(p1.y - p2.y) + 1);
      largest = Math.max(area, largest);
    }
  }

  return largest;
}

function part2() {
  let largest = 0;

  for (let i = 0; i < coords.length; i++) {
    for (let j = i + 1; j < coords.length; j++) {
      const p1 = coords[i]!;
      const p2 = coords[j]!;

      // validate the rectangle

      const area = (Math.abs(p1.x - p2.x) + 1) * (Math.abs(p1.y - p2.y) + 1);
      largest = Math.max(area, largest);
    }
  }

  return largest;
}

console.log(part1());
console.log(part2());
