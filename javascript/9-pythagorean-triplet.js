function specialPythagoreanTriplet(n) {
  let sumOfabc = n;

  for (let c = 3; c <= n; c++) {
    for (let b = 2; b < c; b++) {
      for (let a = 1; a < b; a++) {
        if (a + b + c === sumOfabc && a * a + b * b === c * c) {
          return a * b * c;
        }
      }
    }
  }

  return -1;
}

console.log(specialPythagoreanTriplet(1000));
