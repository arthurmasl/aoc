const path = `${import.meta.dir}/input.txt`;
const file = Bun.file(path);

const text = await file.text();
const textArr = text.trim().split('');

const getPairs = (str: string[], size = 4) => {
  const pairs = [];

  for (let i = 0; i <= str.length - size; i++) {
    pairs.push(str.slice(i, i + size));
  }

  const firstMath = Array(
    ...pairs.map((p) => new Set(p)).find((s) => s.size === size)!,
  ).join('');

  return str.join('').indexOf(firstMath) + size;
};

console.time('t1');
console.log(getPairs(textArr, 14));
console.timeEnd('t1');
// 3-4ms
