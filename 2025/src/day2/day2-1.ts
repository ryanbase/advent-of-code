import { readFileFromArgs } from "../../utils/read-file";

const input = await readFileFromArgs();

const ranges = input.split(",");

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

console.log(result);
