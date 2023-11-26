const path = `${import.meta.dir}/input.txt`;
const file = Bun.file(path);

const text = await file.text();
const textArr = text.trim().split('\n');

const getRanges = (data: string) =>
  data
    .split(',')
    .map((i) => i.split('-'))
    .map((i) => [+i[0], +i[1]]);

const makeArray = (range: number[]) =>
  Array.from({ length: 9 }, (_, i) =>
    i + 1 >= range[0] && i + 1 <= range[1] ? `${i + 1}` : '.',
  );

const contains = (a: number[], b: number[]) =>
  (a[0] >= b[0] && a[1] <= b[1]) || (b[0] >= a[0] && b[1] <= a[1]);
const overlaps = (a: number[], b: number[]) => a[0] <= b[1] && b[0] <= a[1];

const checkFn = (cb: any) => (data: string) => {
  const [range1, range2] = getRanges(data);
  return cb(range1, range2);
};

const containsResult = textArr.filter(checkFn(contains)).length;
console.log(containsResult);
const overlapsResult = textArr.filter(checkFn(overlaps)).length;
console.log(overlapsResult);
