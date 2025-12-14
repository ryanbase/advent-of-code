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

function findPaths2(device: Device, count: number, visited: string[]) {
  if (count > Object.keys(devices).length) {
    return 0;
  }
  if (device.next[0] === "out") {
    return visited.length >= 2 ? 1 : 0;
  }
  if (device.name === "dac" || device.name === "fft") {
    visited.push(device.name);
  }
  console.log(device.name, count, visited);
  let result = 0;
  for (let i = 0; i < device.next.length; i++) {
    result += findPaths2(devices[device.next[i]!]!, count + 1, [...visited]);
  }
  return result;
}

function part1() {
  return findPaths(devices["you"]!, 0);
}

function part2() {
  return findPaths2(devices["svr"]!, 0, []);
}

// console.log(part1());
console.log(part2());
