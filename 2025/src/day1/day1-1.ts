import { readFileFromArgs } from "../../utils/read-file";

const input = await readFileFromArgs();

const steps = input.split("\n").map((step) => {
  const direction = step.charAt(0);
  const number = Number(step.substring(1));
  return { direction, number };
});

let position = 50;
let password = 0;

for (let step of steps) {
  if (step.direction === "L") {
    position -= step.number % 100;
    if (position < 0) {
      position = 100 + position;
    }
  }
  if (step.direction === "R") {
    position += step.number % 100;
    if (position > 99) {
      position = (position % 99) - 1;
    }
  }

  console.log(JSON.stringify({ step, position }));

  if (position === 0) {
    password++;
  }
}

console.log(password);
