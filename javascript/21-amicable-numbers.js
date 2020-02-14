function getProperDivisorsSum(n) {
  const divisors = [1];

  for (let i = 2; i * i <= n; i++) {
    if (n % i === 0) {
      divisors.push(i);

      if (i * i !== n) {
        divisors.push(n / i);
      }
    }
  }

  let result = 0;
  divisors.forEach(divisor => {
    result += divisor;
  });

  return result;
}

function sumAmicableNum(n) {
  let divisorsSum = [0, 0];

  for (let i = 2; i <= n; i++) {
    divisorsSum.push(getProperDivisorsSum(i));
  }

  let result = 0;
  for (let i = 2; i <= n; i++) {
    if (
      divisorsSum[i] <= n &&
      i !== divisorsSum[i] &&
      i === divisorsSum[divisorsSum[i]]
    ) {
      result += i;
    }
  }

  return result;
}

console.log(sumAmicableNum(10000));
