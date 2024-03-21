const text = await Bun.file(`${import.meta.dir}/example.txt`).text();
const data = text.trim().split('\n\n');

const instructions = data[0]
  .replaceAll('L', '0')
  .replaceAll('R', '1')
  .split('')
  .map(Number);

const map = new Map();

data[1].split('\n').forEach((l) => {
  const [key, value] = l.split(' = ');
  map.set(key, value.slice(1, -1).split(', '));
});
const endsInA = [...map.keys()].filter((x) => x.at(-1) === 'A');

let element = 'AAA';

let counter = 0;
let steps = 0;

while (element !== 'ZZZ') {
  if (counter === instructions.length) counter = 0;
  element = map.get(element)[instructions[counter]];

  steps++;
  counter++;
}

console.log(steps);
