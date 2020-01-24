// TODO: cache intermediate results for better performance
function longestCollatzSequence(limit) {
  let maxLen = 0;
  let maxNumber;

  limit--;
  while (limit > 0) {
    let currNumber = limit;
    let currLen = 1;
    while (currNumber !== 1) {
      if (currNumber % 2 === 0) {
        currNumber /= 2;
      } else {
        currNumber = 3 * currNumber + 1;
      }
      currLen++;
    }

    if (currLen > maxLen) {
      maxLen = currLen;
      maxNumber = limit;
    }

    limit--;
  }

  return maxNumber;
}

console.log(longestCollatzSequence(1000000));
