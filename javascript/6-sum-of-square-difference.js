function sumSquareDifference(n) {
  const sumOfSquared = (n * (n + 1) * (2 * n + 1)) / 6;
  let squareOfSum = (n * (n + 1)) / 2;
  squareOfSum = squareOfSum * squareOfSum;

  return Math.abs(sumOfSquared - squareOfSum);
}

console.log(sumSquareDifference(100));
