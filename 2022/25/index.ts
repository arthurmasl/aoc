const text = await Bun.file(`${import.meta.dir}/input.txt`).text();
const lines = text.trim().split('\n');

// 4890
//
//
//  78125 15625 3125 625 125 5 1
//

const convertToNumber = (str: string) => {
  const ops: number[] = [];

  str.split('').forEach((n, i) => {
    const pow = Math.pow(5, Math.abs(i - str.length) - 1);

    if (n === '2') ops.push(pow * 2);
    if (n === '1') ops.push(pow);
    if (n === '0') ops.push(pow * 0);
    if (n === '-') ops.push(-pow);
    if (n === '=') ops.push(-(pow * 2));
  });

  return ops.reduce((acc, curr) => acc + curr, 0);
};

const numbers = lines.map(convertToNumber);
console.log(numbers.reduce((acc, curr) => acc + curr, 0)); // 4890

const conv = (c: number) => {
  if (c === 0) return '0';
  if (c === 1) return '1';
  if (c === 2) return '2';
  if (c === 3) return '=';
  if (c === 4) return '-';
  return '';
};
const numberToSnafu = (number: number) => {
  let snafu = [];

  while (number > 0) {
    let remainder = number % 5;
    snafu.push(conv(remainder));
    if (remainder >= 3) {
      number += 5;
    }
    number = Math.floor(number / 5);
  }
  return snafu.reverse().join('');
};

console.log(numberToSnafu(34978907874317));
