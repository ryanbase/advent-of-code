import { readFileFromArgs } from "../../utils/read-file";

const input = await readFileFromArgs();

type Problem = {
  vals: number[];
  op: string;
};

function sumProbs(probs: Problem[]) {
  return probs
    .map((prob) => {
      if (prob.op === "+") {
        let sum = 0;
        prob.vals.forEach((val) => (sum += val));
        return sum;
      } else if (prob.op === "*") {
        let prod = 1;
        prob.vals.forEach((val) => (prod *= val));
        return prod;
      }
      return 0;
    })
    .reduce((acc, curr) => {
      return acc + curr;
    }, 0);
}

function part1() {
  const rows = input
    .split("\n")
    .map((row) => row.split(" ").filter((val) => val !== ""));

  const probs: Problem[] = new Array(rows[0]!.length)
    .fill(undefined)
    .map(() => ({ vals: [], op: "" }));

  rows.forEach((row, i) => {
    row.forEach((val, j) => {
      if (i === rows.length - 1) {
        probs[j]!.op = val;
      } else {
        probs[j]!.vals.push(Number(val));
      }
    });
  });

  return sumProbs(probs);
}

function part2() {
  const rows = input.split("\n").map((row) => row.split(""));

  let cols = rows[0]?.length;

  if (!cols || cols === 0) {
    return 0;
  }

  const probs: Problem[] = [];
  let currProb: Problem = { vals: [], op: "" };

  for (let i = cols - 1; i >= 0; i--) {
    let currNum = "";
    for (let j = 0; j < rows.length; j++) {
      const val = rows[j]![i];
      if (val === "") {
        continue;
      }
      if (val === "*" || val === "+") {
        currProb.op = val;
        continue;
      }
      currNum += val;
    }
    currNum = currNum.trim();
    if (currNum === "") {
      probs.push(currProb);
      currProb = { vals: [], op: "" };
    } else {
      currProb.vals.push(Number(currNum));
    }
  }
  probs.push(currProb);

  return sumProbs(probs);
}

console.log(part1());
console.log(part2());
