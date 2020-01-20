function isPrime(n) {
  if (n <= 3) {
    return n > 1;
  } else if (n % 2 == 0 || n % 3 == 0) {
    return false;
  }

  let i = 5;
  while (i * i <= n) {
    if (n % i == 0 || n % (i + 2) == 0) {
      return false;
    }
    i += 6;
  }
  return true;
}

function nthPrime(n) {
  let primeCounter = 0;
  let i = 1;

  while (true) {
    if (isPrime(i)) {
      primeCounter++;
    }
    if (primeCounter == n) {
      break;
    }
    i++;
  }
  return i;
}

console.log(nthPrime(10001));
