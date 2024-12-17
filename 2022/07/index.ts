const path = `${import.meta.dir}/input.txt`;
const file = Bun.file(path);

const text = await file.text();
const commands = text.trim().split('\n');

type Dir = {
  files: File[];
  dirs: Map<string, Dir>;
  name: string;
  size: number;
  parent: Dir | null;
};

type File = {
  name: string;
  size: number;
};

const fs: Dir = {
  name: '/',
  dirs: new Map(),
  files: [],
  size: 0,
  parent: null,
};

let currentDir = fs;

const cd = (loc: string) => {
  if (loc === '/' || loc === '..') {
    currentDir = currentDir.parent || currentDir;
    return;
  }

  if (!currentDir.dirs.get(loc)) {
    currentDir.dirs.set(loc, {
      name: loc,
      files: [],
      dirs: new Map(),
      size: 0,
      parent: currentDir,
    });
  }

  currentDir = currentDir.dirs.get(loc)!;
};

const ls = (commandIndex: number) => {
  const listFirst = commands.slice(commandIndex + 1);
  const lastIndex = listFirst.findIndex((cmd) => cmd[0] === '$');
  const list = listFirst
    .slice(0, lastIndex > 0 ? lastIndex : listFirst.length)
    .map((cmd) => cmd.split(' '));
  const files = list.filter((cmd) => cmd[0] !== 'dir');

  files.forEach(([size, name]) => {
    currentDir.files.push({ name, size: +size });
  });
};

let index = 0;
for (const command of commands) {
  if (command.charAt(0) === '$') {
    const cmd = command.slice(2, 4);
    const loc = command.slice(5);

    switch (cmd) {
      case 'cd':
        cd(loc);
        break;
      case 'ls':
        ls(index);
        break;
      default:
        console.log('unknown commmand');
    }
  }
  index++;
}

function* traverse(node: Dir): Generator<Dir> {
  yield node;

  if (node.dirs) {
    for (const child of node.dirs.values()) {
      yield* traverse(child);
    }
  }
}

const calculateSize = (dir: Dir) => {
  for (const d of traverse(dir)) {
    dir.size += d.files.reduce((acc, curr) => acc + curr.size, 0);
  }
};

const sizes = [];
const sizes100 = [];
for (const dir of traverse(fs)) {
  calculateSize(dir);

  sizes.push(dir.size);
  if (dir.size <= 100_000) {
    sizes100.push(dir.size);
  }
}

const available = 70_000_000;
const needed = 30_000_000;

const total = fs.size;
const unused = available - total;
const left = needed - unused;
console.log({ total, unused, left });

const answer1 = sizes100.reduce((acc, curr) => acc + curr, 0);
const answer2 = sizes.filter((s) => s >= left).sort((a, b) => a - b)[0];

console.log(sizes);
// console.log(answer1);
console.log(answer2);
