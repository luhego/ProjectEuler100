/*
Formula for nth Fibonacci number:
  phi = (1 + sqrt(5)) / 2
  Fn = round(phi ^ n / sqrt(5))  .... (1)

Formula for counting digits of a number: 1 + floor(log10(n)) ... (2)

Combining (1) and (2)
Formula for the number of digits of the nth Fibonacci number: 1 + floor(n * log10(phi) - log10(sqrt(5))) 
*/

function countDigistOfNthFibonacci(n) {
  const phi = (1 + Math.sqrt(5)) / 2;
  return 1 + Math.floor(n * Math.log10(phi) - Math.log10(Math.sqrt(5)));
}

function digitFibonacci(n) {
  let index = 1;

  while(true) {
    let numDigits = countDigistOfNthFibonacci(index);
    if (numDigits == n) {
      break;
    }
    index++;
  }

  return index;
}

console.log(digitFibonacci(1000))
