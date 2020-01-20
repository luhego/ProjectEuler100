function primeSummation(n) {
  // We build the sieve of Eratosthenes
  // If sieve[i] = true, i is a prime number
  const sieve = Array(n).fill(true);
  for (let p = 2; p * p <= n; p++) {
    if (sieve[p] === true) {
      for (let i = p * p; i <= n; i += p) {
        sieve[i] = false;
      }
    }
  }

  let result = 0;
  for (let i = 2; i < n; i++) {
    if (sieve[i] === true) {
      result += i;
    }
  }
  return result;
}

console.log(primeSummation(2000000));
