const text = await Bun.file(`${import.meta.dir}/input.txt`).text();
const lines = text.trim().split('\n');

const [t, prevRecord] = lines.map(
  (l) => +l.split(': ')[1].trim().replaceAll(' ', ''),
);

const res = [];

let records = [];

for (let hold = 0; hold <= t; hold++) {
  let timeToTravel = t - hold;
  const newDsit = timeToTravel * hold;
  if (newDsit > prevRecord) {
    records.push(true);
  }

  // console.log(
  //   'hold: ',
  //   hold,
  //   'timeToTravel: ',
  //   timeToTravel,
  //   'newDist: ',
  //   newDsit,
  // );
}
res.push(records);

console.log(
  res.map((i) => i.length).reduce((acc, curr) => acc * curr),
  0,
);
