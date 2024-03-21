const text = await Bun.file(`${import.meta.dir}/input.txt`).text();

type Hand = {
  cards: number[];
  bid: number;
};

const cardToPoint = { A: 14, K: 13, Q: 12, J: 1, T: 10 } as Record<
  string,
  number
>;

const mapToPoints = (cards: string) =>
  cards.split('').map((c) => Number(c) || cardToPoint[c]);

const hands: Hand[] = text
  .trim()
  .split('\n')
  .map((l) => l.split(' '))
  .map((i) => ({ cards: mapToPoints(i[0]), bid: +i[1] }));

const getComb = (cards: number[]) => {
  const counts = Object.values(
    cards.reduce(
      (acc, curr) => ({
        ...acc,
        [curr]: cards.filter((i) => i === curr).length,
      }),
      {},
    ),
  );
  const c = cards.reduce(
    (acc, curr) => ({
      ...acc,
      [curr]: cards.filter((i) => i === curr).length,
    }),
    {},
  );

  const jokers = cards.filter((c) => c === 1).length;

  if (counts.includes(5)) return 7;
  if (counts.includes(4)) {
    if (jokers === 4 || jokers === 1) return 7;
    return 6;
  }
  if (counts.includes(3) && counts.includes(2)) {
    if (jokers >= 2) return 7;
    return 5;
  }
  if (counts.includes(3)) {
    if (jokers === 3) return 6;
    if (jokers === 2) return 7;
    if (jokers === 1) return 6;

    return 4;
  }
  if (counts.filter((c) => c === 2).length === 2) {
    if (jokers === 2) return 6;
    if (jokers === 1 && c[1] === 1) return 5;
    if (jokers === 1) return 4;

    return 3;
  }
  if (counts.includes(2)) {
    if (c[1] === 2) return 4;
    if (c[1] === 1) return 4;
    if (c[1] === 0) return 2;
    return 2;
  }
  if (jokers === 1) return 2;

  return 1;
};

console.log(getComb(mapToPoints('JJJJ2')));

const getFirstHighest = (a: number[], b: number[], i = 0): number[] => {
  if (a[i] === b[i]) return getFirstHighest(a, b, i + 1);
  return a[i] > b[i] ? a : b;
};

const sortHands = (a: Hand, b: Hand) => {
  const combA = getComb(a.cards);
  const combB = getComb(b.cards);

  if (combA === combB) {
    const first = getFirstHighest(a.cards, b.cards);
    return first === a.cards ? 1 : -1;
  }

  return combA - combB;
};

const sortedHands = hands
  .sort(sortHands)
  .map((p, i) => ({ ...p, rank: i + 1 }));

console.log(sortedHands.reduce((acc, curr) => acc + curr.bid * curr.rank, 0));
