const path = `${import.meta.dir}/example.txt`;
const file = Bun.file(path);

const text = await file.text();
const textArr = text.split('\n');

const numbersArr: Array<Array<number>> = [];

for (const i of textArr) {
  if (!numbersArr.length) numbersArr.push([]);
  if (i === '') {
    numbersArr.push([]);
    continue;
  }
  numbersArr.at(-1)?.push(Number(i));
}

const sumsArr = numbersArr.map((n) => n.reduce((acc, curr) => acc + curr, 0));
const sortedArr = sumsArr.sort((a, b) => b - a);
const max = sortedArr[0];
const top3 = sortedArr.slice(0, 3);
const totalTop3 = top3.reduce((acc, curr) => acc + curr);

console.log(textArr);
console.log(numbersArr);
console.log(sumsArr);
console.log(max);
console.log(top3);
console.log(totalTop3);
