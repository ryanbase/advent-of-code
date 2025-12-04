import { readFileFromArgs } from "../../utils/read-file";

const input = await readFileFromArgs();

const ranges = input.split(",");

function part1() {
  let result = 0;

  for (let range of ranges) {
    const ids = range.split("-");
    const id1 = Number(ids[0]);
    const id2 = Number(ids[1]);
    for (let i = id1; i <= id2; i++) {
      const id = String(i);
      if (id.length % 2 !== 0) {
        continue;
      }
      const half1 = id.substring(0, id.length / 2);
      const half2 = id.substring(id.length / 2);
      if (half1 === half2) {
        result += Number(id);
      }
    }
  }

  return result;
}

function part2() {
  let result = 0;

  for (let range of ranges) {
    const ids = range.split("-");
    const id1 = Number(ids[0]);
    const id2 = Number(ids[1]);
    for (let i = id1; i <= id2; i++) {
      const id = String(i);
      result += hasRepeats(id);
    }
  }

  return result;
}

function hasRepeats(id: string): number {
  for (let len = 1; len <= id.length / 2; len++) {
    if (id.length % len !== 0) {
      continue;
    }
    const sub = id.substring(0, len);
    let repeats = true;
    for (let i = len; i < id.length; i += len) {
      const comp = id.substring(i, i + len);
      if (sub !== comp) {
        repeats = false;
        break;
      }
    }
    if (repeats) {
      return Number(id);
    }
  }
  return 0;
}

console.log(part1());
console.log(part2());
