const path = `${import.meta.dir}/input.txt`;
const file = Bun.file(path);

const text = await file.text();
const data = text.trim().split('\n').map(Number);

let timesIncreased = 0;
let increaseCounter = 0;
const counters = [];

for (const [i, curr] of data.entries()) {
  const prev = data[i - 1];
  if (!prev) continue;

  if (prev < curr) {
    timesIncreased += 1;
    increaseCounter += 1;
  } else {
    counters.push(increaseCounter);
    increaseCounter = 0;
  }
}

console.log(timesIncreased);
console.log(counters.reduce((acc, curr) => acc + curr, 0));
