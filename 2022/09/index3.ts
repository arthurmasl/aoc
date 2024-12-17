const file = Bun.file(`${import.meta.dir}/input.txt`);
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

let h: Vec2 = { x: 0, y: 0 };
let t: Vec2 = { x: 0, y: 0 };

const visited: Vec2[] = [t];

const getPos = (curr: Vec2, [dir, n]: Command): Vec2 => {
  if (dir === 'R') return { ...curr, x: curr.x + n };
  if (dir === 'L') return { ...curr, x: curr.x - n };
  if (dir === 'U') return { ...curr, y: curr.y - n };
  if (dir === 'D') return { ...curr, y: curr.y + n };
  return curr;
};

const move = (cmd: Command) => {
  const newPos = getPos(h, cmd);
  h = newPos;

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
};

commands.forEach(([dir, n]) => {
  for (let i = 1; i <= n; i++) {
    move([dir, 1]);
  }
});

console.log(visited.length);
