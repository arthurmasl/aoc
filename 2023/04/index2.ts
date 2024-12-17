const text = await Bun.file(`${import.meta.dir}/example.txt`).text();
const lines = text.trim().split('\n');

const scores: number[] = [];

lines.forEach((card) => {
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
      score === 0 ? (score += 1) : (score *= 2);
    }
  });

  scores.push(score);
});

const res = scores.reduce((acc, curr) => acc + curr, 0);
console.log(scores);
