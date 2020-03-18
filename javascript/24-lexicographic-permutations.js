function swap(permutation, i, j) {
  const temp = permutation[i];
  permutation[i] = permutation[j];
  permutation[j] = temp;
}

function reverse(permutation, i, num) {
  let start = i;
  let end = permutation.length - 1;

  let swapIndex = null;
  let swapValue = Infinity;

  while (start <= end) {
    swap(permutation, start, end);

    if (permutation[start] > num && permutation[start] <= swapValue) {
      swapIndex = start;
      swapValue = permutation[start];
    }
    if (permutation[end] > num && permutation[end] <= swapValue) {
      swapIndex = end;
      swapValue = permutation[end];
    }

    start++;
    end--;
  }

  return swapIndex;
}

function nextPermutation(permutation) {
  let index = permutation.length - 1;
  while (permutation[index - 1] > permutation[index]) {
    index--;
  }

  const swapIndex = reverse(permutation, index, permutation[index - 1]);
  if (swapIndex) {
    swap(permutation, index - 1, swapIndex);
  }
}

function lexicographicPermutations(n) {
  const permutation = [0, 1, 2, 3, 4, 5, 6, 7, 8, 9];
  for(let i = 0; i < n; i++) {
    nextPermutation(permutation);
  }

  return permutation.join("");
}

console.log(lexicographicPermutations(999999));