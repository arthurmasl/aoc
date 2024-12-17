const text = await Bun.file(`${import.meta.dir}/input.txt`).text();
const lines = text.trim().split('\n');

const copies = Array.from({ length: lines.length }).fill(1);
// const copies = Array.from({ length: 6 }).fill(1);
const scores: number[] = [];

lines.forEach((card, i) => {
  const winning = card
    .slice(card.indexOf(':') + 1, card.indexOf('|'))
    .split(' ')
    .filter(Number)
    .map(Number);
  const numbers = card
    .slice(card.indexOf('|') + 1)
    .split(' ')
    .filter(Number)
    .map(Number);

  // console.log(winning);
  // console.log(numbers);

  let score = 0;
  winning.forEach((w) => {
    if (numbers.includes(w)) {
      score += 1;
    }
  });

  for (let y = 0; y < copies[i]; y++) {
    for (let index = 0; index < score; index++) {
      copies[i + index + 1] += 1;
    }
  }

  scores.push(score);
});

// 1 2 4 8 14 1
// console.log(scores);
console.log(copies);
const res = copies.reduce((acc, curr) => acc + curr, 0);
console.log(res);
