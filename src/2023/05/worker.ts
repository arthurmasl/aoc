const text = await Bun.file(`${import.meta.dir}/example.txt`).text();
export const lines = text.trim().split('\n\n');

// prevents TS errors
declare var self: Worker;

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
  const range = ranges.find((r) => seed >= r.src && seed < r.src + r.range);

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

self.onmessage = (event: MessageEvent) => {
  const { range, id } = event.data;
  console.log('Worker started', id);

  let smallest = Infinity;

  let remaining = range.size;
  let start = range.start;

  while (remaining > 0) {
    let seed = range.size - remaining + start;
    const location = getLocation(seed);

    smallest = Math.min(location, smallest);
    remaining--;
  }

  if (remaining === 0) {
    postMessage(smallest);
    console.log('worker terminated', id);

    process.exit();
  }
};
