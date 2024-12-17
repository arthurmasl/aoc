const pathOperations = `${import.meta.dir}/input.txt`;
const fileOperations = Bun.file(pathOperations);
const textOperations = await fileOperations.text();

const operationsRaw = textOperations.trim().split('\n');
const operations = operationsRaw.map((op) =>
  op.split(' ').filter(Number).map(Number),
);

const stacks = [
  ['Q', 'F', 'L', 'S', 'R'],
  ['T', 'P', 'G', 'Q', 'Z', 'N'],
  ['B', 'Q', 'M', 'S'],
  ['Q', 'B', 'C', 'H', 'J', 'Z', 'G', 'T'],
  ['S', 'F', 'N', 'B', 'M', 'H', 'P'],
  ['G', 'V', 'L', 'S', 'N', 'Q', 'C', 'P'],
  ['F', 'C', 'W'],
  ['M', 'P', 'V', 'W', 'Z', 'G', 'H', 'Q'],
  ['R', 'N', 'C', 'L', 'D', 'Z', 'G'],
];

const makeOperation =
  (reverse = false) =>
  (operation: number[]) => {
    const [count, from, to] = operation;
    let cargo = stacks[from - 1].splice(0, count);
    if (reverse) cargo = cargo.reverse();
    stacks[to - 1].unshift(...cargo);
  };

operations.forEach(makeOperation(false));
const result = stacks.map((s) => s[0]).join('');
console.log(result);
