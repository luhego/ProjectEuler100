function multiplyByTwo(number) {
  let carry = 0;
  let result = [];

  for (let i = 0; i < number.length; i++) {
    let temp = number[i] * 2 + carry;
    carry = Math.floor(temp / 10);
    result.push(temp % 10);
  }

  if (carry > 0) {
    result.push(carry);
  }

  return result;
}

function powerDigitSum(exponent) {
  // We use an array to represent a big number
  let number = [2];

  for (let i = 0; i < exponent - 1; i++) {
    number = multiplyByTwo(number);
  }

  let result = 0;
  for (let i = 0; i < number.length; i++) {
    result += number[i];
  }
  return result;
}

console.log(powerDigitSum(1000));
