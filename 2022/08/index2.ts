const path = `${import.meta.dir}/example.txt`;
const file = Bun.file(path);

const text = await file.text();
const data = text
  .trim()
  .split('\n')
  .map((str, y) =>
    str.split('').map((s, x) => ({
      height: +s,
      x,
      y,
    })),
  );

const visible = new Set();

const width = data[0].length - 1;
const height = data.length - 1;

for (const [y, row] of data.entries()) {
  for (const [x, col] of row.entries()) {
    // outer
    if (y === 0 || x === 0 || x === width || y === height) {
      visible.add(col);

      // top outer
      if (y === 0 && x > 0 && x < width) {
        const peekNext = (curr: any, i: number) => {
          const next = data[i]?.[x];
          if (!next) return;

          if (next.height > curr.height) {
            visible.add(next);
            peekNext(next, --i);
          }
        };
        peekNext(col, 1);
      }

      // bottom outer

      if (y === height && x > 0 && x < width) {
        const peekNext = (curr: any, i: number) => {
          const next = data[i]?.[x];
          if (!next) return;

          if (next.height > curr.height) {
            visible.add(next);
            peekNext(next, --i);
          }
        };
        peekNext(col, height - 1);
      }

      // left outer
      if (x === 0 && y > 0 && y < height) {
        const peekNext = (curr: any, i: number) => {
          const next = data[y]?.[i];
          if (!next) return;

          if (next.height > curr.height) {
            visible.add(next);
            peekNext(next, ++i);
          }
        };
        peekNext(col, 1);
      }

      //right outer
      if (x === width && y > 0 && y < height) {
        const ns: any[] = [];
        const peekNext = (curr: any, i: number) => {
          const next = data[y]?.[i];
          if (!next) return;

          // if (next.height > curr.height) {
          // visible.add(next);
          ns.push(next);
          peekNext(next, --i);
          // }
        };
        peekNext(col, width - 1);
        console.log(ns.filter((n) => n.height > col.height));
      }
    }
  }
}
console.log(visible.size);
