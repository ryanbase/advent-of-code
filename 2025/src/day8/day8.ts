import { getFileName, readFileFromArgs } from "../../utils/read-file";

const input = await readFileFromArgs();

type Coord = { x: number; y: number; z: number };
type Dist = { boxes: { box1: Coord; box2: Coord }; distance: number };

const coords: Coord[] = input.split("\n").map((loc) => {
  const coords = loc.split(",");
  return {
    x: Number(coords[0]!),
    y: Number(coords[1]!),
    z: Number(coords[2]!),
  };
});

function getDistanceBetweenPoints(p1: Coord, p2: Coord) {
  const x = Math.pow(p2.x - p1.x, 2);
  const y = Math.pow(p2.y - p1.y, 2);
  const z = Math.pow(p2.z - p1.z, 2);
  return Math.sqrt(x + y + z);
}

function getSortedDistances() {
  const distances: Dist[] = [];

  for (let i = 0; i < coords.length; i++) {
    for (let j = i + 1; j < coords.length; j++) {
      const distance = getDistanceBetweenPoints(coords[i]!, coords[j]!);
      distances.push({
        boxes: { box1: coords[i]!, box2: coords[j]! },
        distance,
      });
    }
  }

  return distances.sort((a: Dist, b: Dist) => a.distance - b.distance);
}

function part1() {
  const circuits: Coord[][] = [];
  const distances = getSortedDistances();
  const pairs = getFileName() === "test.txt" ? 10 : 1000;
  const shortestDistances = distances.slice(0, pairs);

  for (let i = 0; i < shortestDistances.length; i++) {
    const dist = distances[i]!;
    let circuit1: Coord[] | undefined;
    let circuit2: Coord[] | undefined;
    for (let j = 0; j < circuits.length; j++) {
      const circuit = circuits[j]!;
      const contatinsBox1 = circuit.includes(dist.boxes.box1);
      const contatinsBox2 = circuit.includes(dist.boxes.box2);
      if (!circuit1 && contatinsBox1) {
        circuit1 = circuit;
      }
      if (!circuit2 && contatinsBox2) {
        circuit2 = circuit;
      }
      if (circuit1 && circuit2) {
        break;
      }
    }

    if (circuit1 && !circuit2) {
      circuit1.push(dist.boxes.box2);
    } else if (circuit2 && !circuit1) {
      circuit2.push(dist.boxes.box1);
    } else if (!circuit1 && !circuit2) {
      circuits.push([dist.boxes.box1, dist.boxes.box2]);
    } else if (circuit1 && circuit2 && circuit1 !== circuit2) {
      circuit1.push(...circuit2);
      const index = circuits.indexOf(circuit2);
      circuits.splice(index, 1);
    }
  }

  circuits.sort((a: Coord[], b: Coord[]) => {
    return b.length - a.length;
  });

  return circuits.slice(0, 3).reduce((acc, curr) => acc * curr.length, 1);
}

function part2() {
  const circuits: Coord[][] = [];
  const distances = getSortedDistances();

  for (let i = 0; i < distances.length; i++) {
    const dist = distances[i]!;
    let circuit1: Coord[] | undefined;
    let circuit2: Coord[] | undefined;
    for (let j = 0; j < circuits.length; j++) {
      const circuit = circuits[j]!;
      const contatinsBox1 = circuit.includes(dist.boxes.box1);
      const contatinsBox2 = circuit.includes(dist.boxes.box2);
      if (!circuit1 && contatinsBox1) {
        circuit1 = circuit;
      }
      if (!circuit2 && contatinsBox2) {
        circuit2 = circuit;
      }
      if (circuit1 && circuit2) {
        break;
      }
    }

    if (circuit1 && !circuit2) {
      circuit1.push(dist.boxes.box2);
    } else if (circuit2 && !circuit1) {
      circuit2.push(dist.boxes.box1);
    } else if (!circuit1 && !circuit2) {
      circuits.push([dist.boxes.box1, dist.boxes.box2]);
    } else if (circuit1 && circuit2 && circuit1 !== circuit2) {
      circuit1.push(...circuit2);
      const index = circuits.indexOf(circuit2);
      circuits.splice(index, 1);
    }

    if (circuits.length === 1 && circuits[0]!.length === coords.length) {
      return dist.boxes.box1.x * dist.boxes.box2.x;
    }
  }
}

console.log(part1());
console.log(part2());
