import { readFileFromArgs } from "../../utils/read-file";

const input = await readFileFromArgs();

const banks = input.split("\n");

function part1() {
  let result = 0;

  for (let bank of banks) {
    let first = bank.charAt(0);
    let second = bank.charAt(1);
    for (let i = 1; i < bank.length; i++) {
      const curr = bank.charAt(i);
      if (i !== bank.length - 1 && curr > first) {
        first = curr;
        second = bank.charAt(i + 1);
      } else if (curr > second) {
        second = curr;
      }
    }
    result += Number(first + second);
  }

  return result;
}

function part2() {
  let result = 0;

  const findLargestStart = (bank: string, index: number) => {
    let max = bank.charAt(0);
    let maxIndex = 0;
    for (let i = 1; i <= bank.length - (12 - index); i++) {
      if (bank.charAt(i) > max) {
        max = bank.charAt(i);
        maxIndex = i;
      }
    }
    return maxIndex;
  };

  const getJoltage = (bank: string): number => {
    let joltage = "";
    let currBank = bank;
    for (let i = 0; i < 12; i++) {
      if (currBank.length === 12 - i) {
        joltage += currBank;
        break;
      }
      const maxIndex = findLargestStart(currBank, i);
      joltage += currBank.charAt(maxIndex);
      currBank = currBank.substring(maxIndex + 1);
    }
    return Number(joltage);
  };

  for (let bank of banks) {
    result += getJoltage(bank);
  }

  return result;
}

console.log(part1());
console.log(part2());
