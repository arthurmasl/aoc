const path = `${import.meta.dir}/example.txt`;
const file = Bun.file(path);

const text = await file.text();
const data = text
  .trim()
  .split('\n')
  .map((str, y) =>
    str.split('').map((s, x) => ({
      height: +s,
      x,
      y,
    })),
  );

const visible = new Set();

const width = data[0].length;
const height = data.length;

const getLarger = (arr: any[], n: number) => arr.filter((i) => i.height > n);

const removeDup = (arr: any[], reverse = false) => {
  const newArr: any[] = [];
  const sortedArr = reverse
    ? arr.sort((a, b) => a.x - b.x)
    : arr.sort((a, b) => a.x - b.x).reverse();

  sortedArr.forEach((x, i) => {
    if (!newArr.includes(x)) {
      newArr.push(arr.splice(i, 1)[0]);
    }
  });

  return newArr;
};

const add = (row: any[]) => {
  const left = row.slice(0, 1)[0];
  const right = row.slice(-1)[0];

  visible.add(left);
  visible.add(right);

  removeDup(getLarger(row, left.height)).forEach((i) => visible.add(i));
  removeDup(getLarger(row, right.height), true).forEach((i) => visible.add(i));
};

for (const [y, row] of data.entries()) {
  for (const [x, col] of row.entries()) {
    if (x === 0) {
      add(row);
    }

    if (y === 0) {
      const row = [...Array.from({ length: height }, (_, i) => data[i][x])];
      add(row);
    }
  }
}

console.log(removeDup(getLarger(data[1], 2)));
console.log(removeDup(getLarger(data[1], 2), true));
// removeDup(getLarger(data[1], 2));
// removeDup(getLarger(data[1], 2), true);

console.log(visible.size);
