const path = `${import.meta.dir}/example.txt`;
const file = Bun.file(path);

const text = await file.text();
const rows = text
  .trim()
  .split('\n')
  .map((row) => row.split('').map(Number));

const visible = [];
const scores = [];
const width = rows[0].length - 1;
const height = rows.length - 1;

const isEveryVisible = (row: number[], col: number) =>
  row.every((i) => i < col);

const getScore = (row: number[], col: number) => {
  const i = row.findIndex((i) => i >= col);
  return i >= 0 ? row.slice(0, i + 1).length : row.length;
};

for (const [y, row] of rows.entries()) {
  for (const [x, col] of row.entries()) {
    if (y === 0 || x === 0 || x === width || y === height) {
      visible.push(col);
      continue;
    }

    const vRow = [...Array.from({ length: height + 1 }, (_, i) => rows[i][x])];
    const l = row.slice(0, x);
    const r = row.slice(x + 1, width + 1);
    const t = vRow.slice(0, y);
    const b = vRow.slice(y + 1, height + 1);

    scores.push(
      getScore(l.reverse(), col) *
        getScore(r, col) *
        getScore(t.reverse(), col) *
        getScore(b, col),
    );

    if (
      isEveryVisible(l, col) ||
      isEveryVisible(r, col) ||
      isEveryVisible(t, col) ||
      isEveryVisible(b, col)
    ) {
      visible.push(col);
    }
  }
}

// console.log(visible.length);
// console.log(scores.sort((a, b) => b - a)[0]);

function* traverse(node: number[][] | number): Generator<number> {
  if (Array.isArray(node)) {
    for (const row of node) {
      for (const col of row) {
        yield* traverse(col);
      }
    }
  } else {
    yield node;
  }
}

const items = traverse(rows);
for (const i of items) {
  console.log(i);
}
