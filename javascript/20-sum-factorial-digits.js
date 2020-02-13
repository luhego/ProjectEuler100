function multiply(number, x) {
  let result = [];
  let carry = 0;

  for (let i = 0; i < number.length; i++) {
    let temp = number[i] * x + carry;
    carry = Math.floor(temp / 10);
    result.push(temp % 10);
  }

  while (carry > 0) {
    result.push(carry % 10);
    carry = Math.floor(carry / 10);
  }

  return result;
}

function sumFactorialDigits(n) {
  let number = [1];

  for (let i = 1; i <= n; i++) {
    number = multiply(number, i);
  }

  let result = 0;
  for (let i = 0; i < number.length; i++) {
    result += number[i];
  }

  return result;
}

console.log(sumFactorialDigits(100));
