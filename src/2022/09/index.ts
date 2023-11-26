const path = `${import.meta.dir}/input.txt`;
const file = Bun.file(path);

const text = await file.text();
const lines = text.trim().split('\n');

type Direction = 'L' | 'U' | 'R' | 'D';
type Movement = [Direction, number];
type Vec2 = [x: number, y: number];

const ROPE_LENGTH = 10;
const rope = Array.from({ length: ROPE_LENGTH }, () => [0, 0]);

const head = rope[0] as Vec2;
const tail = rope.at(-1) as Vec2;

const visited = new Set<string>(['0,0']);

const areTouching = (distance: Vec2) =>
  distance.every((coord) => Math.abs(coord) <= 1);

lines.forEach((movement) => {
  const [dir, amount] = movement.split(' ') as Movement;
  console.log(dir, amount);

  for (let i = 0; i < +amount; i++) {
    if (dir === 'L') head[0]--;
    if (dir === 'U') head[1]++;
    if (dir === 'R') head[0]++;
    if (dir === 'D') head[1]--;

    for (let j = 1; j < ROPE_LENGTH; j++) {
      const knot = rope[j];
      const prev = rope[rope.indexOf(knot) - 1];
      const distance: Vec2 = [knot[0] - prev[0], knot[1] - prev[1]];

      if (areTouching(distance)) continue;

      knot[0] -= Math.sign(distance[0]);
      knot[1] -= Math.sign(distance[1]);

      if (knot === tail) visited.add(knot.join(','));
    }
  }
});

console.log(visited.size);
