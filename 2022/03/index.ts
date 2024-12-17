const path = `${import.meta.dir}/input.txt`;
const file = Bun.file(path);

const text = await file.text();
const textArr = text.trim().split('\n');

const isUpper = (char: string) => char.toUpperCase() === char;
const countPoint = (char: string) => {
  const code = char.charCodeAt(0);
  return isUpper(char) ? code - 65 + 27 : code - 96;
};

const checkTwo = (items: string) => {
  const left = items.split('');
  const right = left.splice(items.length / 2);
  const repeating = left.find((i) => right.includes(i)) ?? '';

  return countPoint(repeating);
};

const resultTwo = textArr.map(checkTwo).reduce((acc, curr) => acc + curr);
console.log(resultTwo);

const checkThree = (bags: string[]) => {
  const three = bags.splice(0, 3);
  const repeating = Array.from(three[0]).find(
    (i) => three[1].includes(i) && three[2].includes(i),
  );

  return countPoint(repeating!);
};

let resultThree = 0;
while (textArr.length) {
  resultThree += checkThree(textArr);
}

console.log(resultThree);
