const text = await Bun.file(`${import.meta.dir}/example.txt`).text();
const lines = text.trim().split('\n');

const nStrings = {
  one: 1,
  two: 2,
  three: 3,
  four: 4,
  five: 5,
  six: 6,
  seven: 7,
  eight: 8,
  nine: 9,
};

const strToNum = (str: string) =>
  Object.keys(nStrings).includes(str) ? +str.replace(str, nStrings[str]) : +str;

const replace = (str: string) => {
  const r = [];

  for (let x = 0; x < str.length; x++) {
    for (let y = 0; y <= 5; y++) {
      const s = str.substring(x, y);

      if (s.length === 1 && Number(s)) {
        r.push(s);
        break;
      }
      if (Object.keys(nStrings).includes(s)) {
        r.push(s);
        break;
      }
    }
  }

  return `${strToNum(r[0])}${strToNum(r[r.length - 1])}`;
};

const nums = lines.map((l) => replace(l));
console.log(nums);
console.log(nums.map(Number).reduce((a, b) => a + b, 0)); //278
