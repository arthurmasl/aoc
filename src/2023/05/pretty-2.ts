const text = await Bun.file(`${import.meta.dir}/example.txt`).text();
const lines = text.trim().split('\n\n');

type SeedRange = {
  start: number;
  end: number;
  size: number;
};
const seeds = lines[0].split(': ')[1].split(' ').map(Number);
const seedsRanges = [];

for (let i = 0; i < seeds.length; i += 2) {
  seedsRanges.push({
    start: seeds[i],
    end: seeds[i] + seeds[i + 1] - 1,
    size: seeds[i + 1],
  });
}

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
  let newLocation = seed;
  maps.forEach((map) => {
    newLocation = getSeed(map, newLocation);
  });
  return newLocation;
};

const locations: number[] = [];

let smallest = Infinity;

seedsRanges.forEach((range) => {
  let remaining = range.size;
  let start = range.start;

  while (remaining > 0) {
    let seed = range.size - remaining + start;
    const location = getLocation(seed);

    if (location > start + range.size) {
      break;
    }

    smallest = Math.min(smallest, location);
    remaining--;
    locations.push(location);
  }
});

console.log('------');
console.log('result: ', locations.sort((a, b) => a - b)[0]);
