function isAbundantNumber(n) {
  let properDivisorsSum = 1;

  for (let i = 2; i * i <= n; i++) {
    if (n % i === 0) {
      properDivisorsSum += i;

      if (i * i !== n) {
        properDivisorsSum += n / i;
      }
    }
  }

  return properDivisorsSum > n;
}

function sumOfNonAbundantNumbers(n) {
  const abundantNumbers = [];
  for (let i = 1; i <= n; i++) {
    if (isAbundantNumber(i)) {
      abundantNumbers.push(i);
    }
  }

  const sumOfAbundantNumbers = new Set();
  for (let i = 0; i < abundantNumbers.length; i++) {
    for (let j = i; j < abundantNumbers.length; j++) {
      if (abundantNumbers[i] + abundantNumbers[j] > n) {
        continue;
      }
      sumOfAbundantNumbers.add(abundantNumbers[i] + abundantNumbers[j]);
    }
  }

  const sortedSumOfAbundantNumbers = Array.from(sumOfAbundantNumbers).sort((a, b) => a - b);

  let result = 0;
  let current = 1;
  for (const s of sortedSumOfAbundantNumbers) {
    result += ((s - 1) * s) / 2;
    if (current > 1) {
      result -= ((current - 1) * current) / 2;
    }

    current = s + 1;
  }

  return result;
}

console.log(sumOfNonAbundantNumbers(28123));
