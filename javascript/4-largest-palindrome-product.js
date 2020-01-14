function isPalindrome(n) {
  const s = n.toString();

  let i = 0;
  let j = s.length - 1;
  while (i <= j) {
    if (s[i] != s[j]) {
      return false;
    }
    i++;
    j--;
  }
  return true;
}

function largestPalindromeProduct(n) {
  const start = Math.pow(10, n - 1);
  const end = Math.pow(10, n) - 1;

  let currProduct;
  let maxProduct = 0;
  for (let i = start; i <= end; i++) {
    for (let j = start; j <= end; j++) {
      currProduct = i * j;
      if (currProduct > maxProduct && isPalindrome(currProduct)) {
        maxProduct = currProduct;
      }
    }
  }
  return maxProduct;
}

console.log(largestPalindromeProduct(3));
