const workerURL = new URL('worker.ts', import.meta.url).href;

const text = await Bun.file(`${import.meta.dir}/example.txt`).text();
export const lines = text.trim().split('\n\n');

const seeds = lines[0].split(': ')[1].split(' ').map(Number);
const seedsRanges = [];

for (let i = 0; i < seeds.length; i += 2) {
  seedsRanges.push({
    start: seeds[i],
    end: seeds[i] + seeds[i + 1] - 1,
    size: seeds[i + 1],
  });
}

console.log(seedsRanges);

const locations: number[] = [];
let workers = seedsRanges.length;

seedsRanges.forEach((range, id) => {
  const worker = new Worker(workerURL, { smol: false });

  worker.postMessage({ id, range });

  worker.onmessage = (event) => {
    locations.push(event.data);
    // console.log('workers left: ', workers);
    console.log('seed nr', id, event.data);
    console.log(locations);
    console.log('result: ', Math.min(...locations));
  };

  worker.addEventListener('close', () => {
    workers--;

    if (workers === 0) {
      console.log('------');
      console.log('result: ', Math.min(...locations));
    }
  });
});
