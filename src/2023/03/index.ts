const text = await Bun.file(`${import.meta.dir}/input.txt`).text();
const lines = text.trim().split('\n');
const board = lines.map((l) => l.split(''));

const isNumber = (y: number, x: number) =>
  board?.[y]?.[x] && !isNaN(board[y][x]);
const getDirsPos = (y: number, x: number) => [
  [y, x - 1],
  [y, x + 1],
  [y - 1, x],
  [y + 1, x],
  [y - 1, x - 1],
  [y - 1, x + 1],
  [y + 1, x - 1],
  [y + 1, x + 1],
];

const findNumberAtPosition = (str: string, targetIndex: number) => {
  const regex = /\d+/g;
  let match;

  while ((match = regex.exec(str)) !== null) {
    const number = match[0];
    const numberStartIndex = match.index;
    const numberEndIndex = regex.lastIndex - 1;

    if (targetIndex >= numberStartIndex && targetIndex <= numberEndIndex) {
      return number;
    }
  }

  return null;
};

const removeLeftNumber = (y: number, x: number) => {
  if (board[y]?.[x]) board[y][x] = '.';
  if (isNumber(y, x - 1)) {
    if (board[y]?.[x - 1]) board[y][x - 1] = '.';
    removeLeftNumber(y, x - 1);
  }
};

const removeRightNumber = (y: number, x: number) => {
  if (board[y]?.[x]) board[y][x] = '.';

  if (isNumber(y, x + 1)) {
    if (board[y]?.[x + 1]) board[y][x + 1] = '.';
    removeRightNumber(y, x + 1);
  }
};

const getN = (y: number, x: number) => {
  const n = findNumberAtPosition(board[y].join(''), x);
  return n;
};

const ratios = [];

board.forEach((row, y) => {
  row.forEach((col, x) => {
    if (col === '*') {
      const dirsPos = getDirsPos(y, x);

      const numbers: string[] = [];
      dirsPos.forEach((pos) => {
        if (isNumber(...pos)) {
          const n = getN(...pos);

          if (n) {
            removeLeftNumber(...pos);
            removeRightNumber(...pos);
            numbers.push(n);
          }
        }
      });

      if (numbers.length > 1) {
        ratios.push(numbers.map(Number).reduce((acc, curr) => acc * curr, 1));
      }
      console.log(numbers);
    }
  });
});

console.log(ratios.reduce((acc, curr) => acc + curr));
