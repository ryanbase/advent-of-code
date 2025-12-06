import { readFileFromArgs } from "../../utils/read-file";

const inputStr = await readFileFromArgs();
const input = inputStr.split("\n");

type Range = { min: number; max: number };

const ranges: Range[] = [];
const ids: number[] = [];

let idSection = false;
input.forEach((val) => {
  if (val === "") {
    idSection = true;
    return;
  }
  if (idSection) {
    ids.push(Number(val));
  } else {
    const range = val.split("-");
    ranges.push({ min: Number(range[0]), max: Number(range[1]) });
  }
});

function isIdInRanges(id: number) {
  for (let range of ranges) {
    if (id >= range.min && id <= range.max) {
      return true;
    }
  }
}

function part1() {
  let result = 0;

  ids.forEach((id) => {
    if (isIdInRanges(id)) {
      result++;
    }
  });

  return result;
}

function isInRange(val: number, range: Range) {
  return val >= range.min && val <= range.max;
}

function combineRanges(ranges: Range[]) {
  const combinedRanges: Range[] = [];

  for (let range of ranges) {
    let added = false;
    for (let combinedRange of combinedRanges) {
      if (
        isInRange(range.min, combinedRange) &&
        isInRange(range.max, combinedRange)
      ) {
        added = true;
        break;
      }
      if (
        isInRange(range.min, combinedRange) &&
        range.max > combinedRange.max
      ) {
        combinedRange.max = range.max;
        added = true;
      }
      if (
        isInRange(range.max, combinedRange) &&
        range.min < combinedRange.min
      ) {
        combinedRange.min = range.min;
        added = true;
      }
      if (added) {
        break;
      }
    }
    if (!added) {
      combinedRanges.push(range);
    }
  }

  return combinedRanges;
}

function part2() {
  let result = 0;

  ranges.sort((a: Range, b: Range) => {
    return a.min < b.min ? -1 : 1;
  });

  for (let combinedRange of combineRanges(ranges)) {
    result += combinedRange.max - combinedRange.min + 1;
  }

  return result;
}

console.log(part1());
console.log(part2());
