import { readFileFromArgs } from "../../utils/read-file";

const input = await readFileFromArgs();

type Device = {
  name: string;
  next: string[];
};

const devices: Record<string, Device> = {};

input.split("\n").forEach((device) => {
  const [name, next] = device.split(": ");
  devices[name!] = { name: name!, next: next!.split(" ") };
});

const memo = new Map<string, number>();

function findPaths(device: Device, count: number) {
  if (count > Object.keys(devices).length) {
    return 0;
  }
  if (device.next[0] === "out") {
    return 1;
  }
  let result = 0;
  for (let i = 0; i < device.next.length; i++) {
    result += findPaths(devices[device.next[i]!]!, count + 1);
  }
  return result;
}

// I did use Cursor to figure out how to make this function more performant, which is
// where the idea for the memo came from.
function findPaths2(device: Device, count: number, visited: Set<string>) {
  if (count > Object.keys(devices).length) {
    return 0;
  }

  const hasDac = visited.has("dac") ? "1" : "0";
  const hasFft = visited.has("fft") ? "1" : "0";
  const memoKey = `${device.name}:${hasDac}:${hasFft}:${visited.size}`;

  if (memo.has(memoKey)) {
    return memo.get(memoKey)!;
  }

  if (device.next.includes("out")) {
    const result = visited.has("dac") && visited.has("fft") ? 1 : 0;
    memo.set(memoKey, result);
    return result;
  }

  const newVisited = new Set(visited);
  newVisited.add(device.name);

  let result = 0;
  device.next.forEach(
    (next) => (result += findPaths2(devices[next]!, count + 1, newVisited))
  );

  if (visited.size < 50) {
    memo.set(memoKey, result);
  }

  return result;
}

function part1() {
  return findPaths(devices["you"]!, 0);
}

function part2() {
  return findPaths2(devices["svr"]!, 0, new Set());
}

// console.log(part1());
console.log(part2());
