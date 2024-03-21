const text = await Bun.file(`${import.meta.dir}/example.txt`).text();
const players = text
  .trim()
  .split('\n')
  .map((l) => l.split(' '))
  .map((i) => ({ cards: i[0], bid: +i[1], points: 0, rank: 0 }));

const mapToPoints = (cards: string[]) =>
  cards.reverse().reduce((acc, curr, i, _) => ({ ...acc, [curr]: i }), {});

const cards = ['A', 'K', 'Q', 'J', 'T', '9', '8', '7', '6', '5', '4', '3', '2'];
const cardsPoints = mapToPoints(cards);

players.forEach((player) => {
  const pointsHand = player.cards.split('').map((i) => cardsPoints[i]);

  const counts = pointsHand.reduce((acc, curr, _, arr) => {
    return { ...acc, [curr]: arr.filter((i) => i === curr).length };
  }, {});

  console.log(player.cards);

  const countsValues = Object.values(counts);

  const five = countsValues.includes(5) && 7;
  const four = countsValues.includes(4) && 6;
  const fullHouse = countsValues.includes(3) && countsValues.includes(2) && 5;
  const three = countsValues.includes(3) && 4;
  const twoPairs = countsValues.filter((v) => v === 2).length === 2 && 3;
  const pair = countsValues.includes(2) && 2;

  const combs = [
    { five },
    { four },
    { fullHouse },
    { three },
    { twoPairs },
    { pair },
  ];

  let mult = 1;
  let firstComb = pointsHand[0];
  let secondComb = pointsHand[1];
  const pointsSum = pointsHand.reduce((acc, curr) => acc + curr);

  for (const comb of combs) {
    const [key, value] = Object.entries(comb)[0];

    if (!!value) {
      // const comboCards = Object.entries(counts).sort((a, b) => b[1] - a[1]);

      // if (['twoPairs', 'fullHouse'].includes(key)) {
      //   secondComb = +comboCards[1][0];
      // }

      mult = value;
      // firstComb = +comboCards[0][0];
      break;
    }
  }

  // player.points = pointsHand.reduce((acc, curr) => acc + curr) + mult;

  console.log({ pointsSum, mult, firstComb, secondComb });

  player.points = mult * (firstComb > secondComb ? firstComb : secondComb);
});

const sortedPlayers = players
  .sort((a, b) => a.points - b.points)
  .map((p, i) => ({ ...p, rank: i + 1 }));

// console.log(sortedPlayers);
console.log(sortedPlayers.reduce((acc, curr) => acc + curr.bid * curr.rank, 0));
