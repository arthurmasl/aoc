const normalizeBoardSmart = (n: number) => {
  if (
    h.y - n < 0 ||
    h.y + n > board.length ||
    h.x - n < 0 ||
    h.x + n > board[0].length
  ) {
    const newSize = n * 2 + 1;
    const newPos = Math.floor(newSize / 2);

    board = Array.from({ length: newSize }, () => Array(newSize).fill(EMPTY));
    h = { x: newPos, y: newPos };
    board[newPos][newPos] = H;
  }
};
