const single_digits = ["", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"];

const two_digits = [
  "ten", "eleven", "twelve", "thirteen", "fourteen", "fifteen", "sixteen", "seventeen", "eighteen",
  "nineteen"
]

const tens = [
  "", "ten", "twenty", "thirty", "forty", "fifty", "sixty", "seventy", "eighty", "ninety"
];

const thousands = ["", "thousand"];

function helper(number) {
  if (number === 0) {
    return "";
  } else if (number < 10) {
    return single_digits[number];
  } else if(number < 20) {
    return two_digits[number - 10];
  } else if (number < 100) {
    return tens[Math.floor(number / 10)] + helper(number % 10);
  } else {
    let temp = single_digits[Math.floor(number / 100)] + "hundred";
    if (number % 100 !== 0) {
      temp += "and" + helper(number % 100);
    }
    return temp;
  }
}

function numberToWord(number) {
  let result = "";

  let i = 0;
  while (number > 0) {
    if (number % 1000 !== 0) {
      result += helper(number % 1000) + thousands[i] + result;
    }
    number = Math.floor(number / 1000);
    i++;
  }

  return result;
}

function numberLetterCounts(limit) {
  let result = 0;
  for (let i = 1; i <= limit; i++) {
    let strNumber = numberToWord(i);
    result += strNumber.length;
  }
  return result;
}

console.log(numberLetterCounts(1000));
