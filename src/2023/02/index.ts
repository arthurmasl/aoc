const text = await Bun.file(`${import.meta.dir}/input.txt`).text();
const lines = text.trim().split('\n');

const red = 12;
const green = 13;
const blue = 14;

const getNumber = (str: string) => +str.match(/(\d+)/)[0];

const getNumbers = (arr, color) =>
  arr.reduce(
    (acc, curr) => (curr.includes(color) ? acc + getNumber(curr) : acc),
    0,
  );

const collectCubes = (game) => {
  const id = getNumber(game.slice(0, game.search(':')), 'Game');

  const arr = game
    .slice(game.search(':') + 2)
    .split(';')
    .map((g) => g.split('; '));

  const arr2 = game
    .slice(game.search(':') + 2)
    .split(';')
    .flatMap((g) => g.split(', '));

  const mR = arr2
    .filter((a) => a.includes('red'))
    .map(getNumber)
    .sort((a, b) => b - a)[0];
  const mG = arr2
    .filter((a) => a.includes('green'))
    .map(getNumber)
    .sort((a, b) => b - a)[0];
  const mB = arr2
    .filter((a) => a.includes('blue'))
    .map(getNumber)
    .sort((a, b) => b - a)[0];

  const power = mR * mG * mB;

  const p = arr.map((s) => {
    const sArr = s[0].split(',');

    const r = getNumbers(sArr, 'red');
    const g = getNumbers(sArr, 'green');
    const b = getNumbers(sArr, 'blue');
    return r <= red && g <= green && b <= blue;
  });

  return power;
  return p.every((r) => !!r) && id;
};

collectCubes(lines[0]);

const res = lines
  .map(collectCubes)
  .filter(Number)
  ?.reduce((acc, curr) => acc + curr);

console.log(res);
