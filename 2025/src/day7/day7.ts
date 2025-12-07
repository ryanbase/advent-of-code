import { readFileFromArgs } from "../../utils/read-file";

const input = await readFileFromArgs();

type Point = {
  row: number;
  col: number;
};

const rows = input.split("\n").map((row) => row.split(""));

function part1() {
  const beams = new Array(rows.length)
    .fill(undefined)
    .map(() => new Set<number>());

  const start = rows[0]?.indexOf("S");

  if (!start) {
    return 0;
  }

  beams[1]?.add(start);

  let splits = 0;

  for (let i = 2; i < rows.length; i++) {
    for (let j = 0; j < rows[i]!.length; j++) {
      const val = rows[i]![j];
      if (val === "^" && beams[i - 1]?.has(j)) {
        splits++;
        if (j > 0) {
          beams[i]?.add(j - 1);
        }
        if (j < rows[i]!.length - 1) {
          beams[i]?.add(j + 1);
        }
      } else if (val === "." && beams[i - 1]?.has(j)) {
        beams[i]?.add(j);
      }
    }
  }

  return splits;
}

console.log(part1());
