function* triangleNumberGenerator(n) {
  let index = 1;
  let currentTriangleNumber = 1;
  while (true) {
    yield currentTriangleNumber;
    index++;
    currentTriangleNumber += index;
  }
}

function countDivisors(n) {
  let numDivisors = 0;
  for (let i = 1; i * i <= n; i++) {
    if (n % i === 0) {
      if (i * i === n) {
        numDivisors++;
      } else {
        numDivisors += 2;
      }
    }
  }
  return numDivisors;
}

function divisibleTriangleNumber(n) {
  const gen = triangleNumberGenerator();

  while(true) {
    const currTriangleNumber = gen.next().value;
    if (countDivisors(currTriangleNumber) > n) {
      return currTriangleNumber;
    }
  }
}

console.log(divisibleTriangleNumber(500));
