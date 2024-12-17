const text = await Bun.file(`${import.meta.dir}/example.txt`).text();
const lines = text.trim().split('\n');
const board = lines.map((l) => l.split(''));
const notAssignedBoard = lines.map((l) => l.split(''));
const assignedBoard = lines.map((l) => l.split('').map((x) => '.'));

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

const removeLeftNumber = (y: number, x: number) => {
  if (notAssignedBoard[y]?.[x]) notAssignedBoard[y][x] = '.';
  if (isNumber(y, x - 1)) {
    if (notAssignedBoard[y]?.[x - 1]) notAssignedBoard[y][x - 1] = '.';
    removeLeftNumber(y, x - 1);
  }
};

const removeRightNumber = (y: number, x: number) => {
  if (notAssignedBoard[y]?.[x]) notAssignedBoard[y][x] = '.';

  if (isNumber(y, x + 1)) {
    if (notAssignedBoard[y]?.[x + 1]) notAssignedBoard[y][x + 1] = '.';
    removeRightNumber(y, x + 1);
  }
};

const assignLeftNumber = (y: number, x: number) => {
  assignedBoard[y][x] = board[y][x];
  if (isNumber(y, x - 1)) {
    assignedBoard[y][x - 1] = board[y][x - 1];
    assignLeftNumber(y, x - 1);
  }
};

const assignRightNumber = (y: number, x: number) => {
  assignedBoard[y][x] = board[y][x];
  if (isNumber(y, x + 1)) {
    assignedBoard[y][x + 1] = board[y][x + 1];
    assignRightNumber(y, x + 1);
  }
};

board.forEach((row, y) => {
  row.forEach((col, x) => {
    if (isNaN(col) && col !== '.') {
      const dirsPos = getDirsPos(y, x);

      dirsPos.forEach((pos) => {
        if (isNumber(...pos)) {
          removeLeftNumber(...pos);
          removeRightNumber(...pos);

          assignLeftNumber(...pos);
          assignRightNumber(...pos);
        }
      });
    }
  });
});

const sum = assignedBoard
  .flat()
  .map((x) => +x)
  .join('')
  .replaceAll('NaN', ' ')
  .split(' ')
  .filter(Number)
  .map(Number);
// .reduce((acc, curr) => acc + curr, 0);

console.log(sum);
// console.log(sum2);

// iterate
// if symbol ->
// check nearby
// delete numbers
//
// 448437 - bad
