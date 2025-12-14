import { readFileFromArgs } from "../../utils/read-file";

const input = await readFileFromArgs();

type Machine = {
  lights: (0 | 1)[];
  buttons: number[][];
};

const machines: Machine[] = input.split("\n").map((machine) => {
  const all = machine.split(" ");
  const lights = all[0]!
    .replace("[", "")
    .replace("]", "")
    .split("")
    .map((light) => {
      return light === "." ? 0 : 1;
    });
  const buttons = all.slice(1, all.length - 1).map((button) =>
    button
      .replace("(", "")
      .replace(")", "")
      .split(",")
      .map((b) => Number(b))
  );
  return {
    lights,
    buttons,
  };
});

function configureMachine(machine: Machine) {
  let presses = 0;
  const lights = new Array(machine.lights.length).fill(0);

  return presses;
}

function part1() {
  let total = 0;
  machines.forEach((machine) => (total += configureMachine(machine)));
  return total;
}

console.log(part1());
