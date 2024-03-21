const text = await Bun.file(`${import.meta.dir}/example.txt`).text();
const lines = text.trim().split('\n\n');

const seeds = lines[0].split(': ')[1].split(' ').map(Number);

const maps = lines.slice(1).map((l) =>
  l
    .split('\n')
    .slice(1)
    .map((i) => {
      const d = i.split(' ').map(Number);

      return {
        dest: d[0],
        src: d[1],
        range: d[2],
      };
    }),
);

const getSeed = (ranges: any[], seed: number) => {
  const range = ranges.find((r) => r.src <= seed && seed <= r.src + r.range);
  if (range) {
    const newSeed = range.dest + (seed - range.src);
    return newSeed as number;
  }
  return seed;
};

const getLocation = (seed: number) => {
  let newSeed = seed;
  maps.forEach((map) => {
    newSeed = getSeed(map, newSeed);
  });
  return newSeed;
};

const res: number[] = [];
seeds.forEach((seed) => {
  res.push(getLocation(seed));
});

console.log(res.sort((a, b) => a - b)[0]);
