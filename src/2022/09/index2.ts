const path = `${import.meta.dir}/example.txt`;
const file = Bun.file(path);

const text = await file.text();

type Command = [string, number];
type Vec2 = {
  x: number;
  y: number;
};

const commands: Command[] = text
  .trim()
  .split('\n')
  .map((str) => str.split(' '))
  .map(([str, n]) => [str, +n]);

const SIZE = 30;
const H = 'H';
const T = 'T';
const EMPTY = '.';
const VISITED = '#';
const TAIL = true;
const TAIL_LENGTH = 3;

let h: Vec2 = { x: 0, y: 0 };
let t: Vec2 = { x: 0, y: 0 };
let tail = Array.from({ length: TAIL_LENGTH }, () => ({ x: 0, y: 0 }));

const visited: Vec2[] = [];
let board: string[][] = [[]];

const getPos = (curr: Vec2, [dir, n]: Command): Vec2 => {
  if (dir === 'R') return { ...curr, x: curr.x + n };
  if (dir === 'L') return { ...curr, x: curr.x - n };
  if (dir === 'U') return { ...curr, y: curr.y - n };
  if (dir === 'D') return { ...curr, y: curr.y + n };
  return curr;
};

const clear = () =>
  (board = Array.from({ length: SIZE }, () => Array(SIZE).fill(EMPTY)));

const normalizeBoard = () => {
  const newPos = Math.floor(SIZE / 2);

  h = { x: newPos, y: newPos };
  if (TAIL) {
    tail = tail.map(() => ({ x: newPos, y: newPos }));
  } else {
    t = { x: newPos, y: newPos };
  }

  visited.push(t);
  clear();
  draw();
};

const move = (cmd: Command) => {
  const newPos = getPos(h, cmd);
  h = newPos;

  if (!TAIL) {
    const distance = Math.floor(Math.hypot(t.x - newPos.x, t.y - newPos.y));
    if (distance > 1) {
      let tPos = getPos(t, cmd);
      const xDiff = Math.abs(h.x - t.x);
      const yDiff = Math.abs(h.y - t.y);

      xDiff > yDiff ? (tPos.y = h.y) : (tPos.x = h.x);
      t = tPos;

      if (!visited.find((v) => v.x === tPos.x && v.y === tPos.y)) {
        visited.push(tPos);
      }
    }
  } else {
    tail.forEach((ti, i) => {
      const prev = tail[i - 1] || h;
      const distance = Math.floor(Math.hypot(ti.x - prev.x, ti.y - prev.y));

      if (distance > 1) {
        let tPos = getPos(ti, cmd);
        const xDiff = Math.abs(prev.x - ti.x);
        const yDiff = Math.abs(prev.y - ti.y);

        xDiff > yDiff ? (tPos.y = prev.y) : (tPos.x = prev.x);
        tail[i] = tPos;

        if (!visited.find((v) => v.x === tPos.x && v.y === tPos.y)) {
          visited.push(tPos);
        }
      }
    });
  }
};

const draw = () => {
  clear();

  visited.forEach((vec) => {
    // board[vec.y][vec.x] = VISITED;
  });

  board[h.y][h.x] = H;
  if (TAIL) {
    tail.forEach((t, i) => {
      board[t.y][t.x] = `${i + 1}`;
    });
  } else {
    board[t.y][t.x] = T;
  }
};

const execute = ([dir, n]: Command) => {
  for (let i = 1; i <= n; i++) {
    move([dir, 1]);
  }

  draw();
};

normalizeBoard();
commands.forEach(execute);

console.log(board.map((s) => s.join(' ')).join('\n'));
console.log(visited.length);
