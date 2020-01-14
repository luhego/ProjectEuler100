function gcd(a, b) {
  if (a == 0) {
    return b;
  }
  return gcd(b % a, a);
}

function lcm(a, b) {
  return (a * b) / gcd(a, b);
}

function smallestMult(n) {
  let result = 1;
  for (let i = 1; i <= n; i++) {
    result = lcm(result, i);
  }
  return result;
}

console.log(smallestMult(20));
