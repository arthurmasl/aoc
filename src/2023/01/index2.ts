const text = await Bun.file(`${import.meta.dir}/example.txt`).text();
const lines = text.trim().split('\n');

const nStrings = [
  'one',
  'two',
  'three',
  'four',
  'five',
  'six',
  'seven',
  'eight',
  'nine',
];

const strToNum = (str: string) => {
  str = str.replace('one', '1');
  str = str.replace('two', '2');
  str = str.replace('three', '3');
  str = str.replace('four', '4');
  str = str.replace('five', '5');
  str = str.replace('six', '6');
  str = str.replace('seven', '7');
  str = str.replace('eight', '8');
  str = str.replace('nine', '9');

  return str;
};

const replace = (str: string) => {
  const n = nStrings
    .map((ns) => ({ i: str.search(ns), n: ns }))
    .filter((n) => n.i >= 0)
    .sort((a, b) => a.i - b.i);

  const s = str
    .split('')
    .map(Number)
    .map((ss, si) => ({ n: ss, i: si }))
    .filter((s) => s.n >= 0)
    .sort((a, b) => a.i - b.i);

  if (n.length && s.length) {
    const f = n[0]?.i < s[0]?.i ? n[0].n : s[0].n;
    const l =
      n[n.length - 1].i < s[s.length - 1].i
        ? n[n.length - 1].n
        : s[s.length - 1].n;

    return `${strToNum(String(f))}${strToNum(String(l))}`;
  }

  if (!s.length) {
    const f = n[0].n;
    const l = n[n.length - 1].n;

    return `${strToNum(String(f))}${strToNum(String(l))}`;
  }

  if (!n.length) {
    const f = s[0].n;
    const l = s[s.length - 1].n;

    return `${f}${l}`;
  }
};

const nums = lines.map((l) => replace(l));
console.log(nums);
console.log(nums.map(Number).reduce((a, b) => a + b, 0));
