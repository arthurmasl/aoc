const text = await Bun.file(`${import.meta.dir}/input.txt`).text();
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

const starts = [...map.keys()].filter((x) => x.at(-1) === 'A');

const loops = starts.reduce((acc, curr) => {
  let steps = 0;
  let counter = 0;
  while (curr.at(-1) !== 'Z') {
    if (counter === instructions.length) counter = 0;
    curr = map.get(curr)[instructions[counter]];
    steps++;
    counter++;
  }

  acc = { ...acc, [curr]: steps };
  return acc;
}, {});

const findLCM = (a: number, b: number) => {
  const gcd = (x: number, y: number): number => (y === 0 ? x : gcd(y, x % y));
  return (a * b) / gcd(a, b);
};

const getRes = (arr: number[]) => {
  let lcm = arr[0];
  for (let i = 1; i < arr.length; i++) {
    lcm = findLCM(lcm, arr[i]);
  }
  return lcm;
};

const res = getRes(Object.values(loops));

console.log(res);
