const path = `${import.meta.dir}/example.txt`;
const file = Bun.file(path);

const text = await file.text();
const textArr = text.split('\n');

const wins = ['A Y', 'B Z', 'C X'];
const draws = ['A X', 'B Y', 'C Z'];
const loses = ['A Z', 'B X', 'C Y'];

const points: { [key: string]: number } = { X: 1, Y: 2, Z: 3 };
const predicate = (op: string) => (i: string) => i[0] === op;

const checkFn = ([op, _, me]: string, cheat = false): number => {
  let score = points[me] ?? 0;

  if (wins.includes(`${op} ${me}`)) score += 6;
  if (draws.includes(`${op} ${me}`)) score += 3;

  if (cheat) {
    if (me === 'X') return checkFn(loses.find(predicate(op))!);
    if (me === 'Y') return checkFn(draws.find(predicate(op))!);
    if (me === 'Z') return checkFn(wins.find(predicate(op))!);
  }

  return score;
};

const result = textArr.reduce((acc, curr) => acc + checkFn(curr, true), 0);

console.log(result);
