/**
 * This problem also be solved mathematically using the formula:
 * 1 -> 2
 * 2 -> 6
 * 3 -> 20
 * 4 -> 70
 * n -> factorial(2 * n) / (factorial(n) ** 2) 
 */

function countLatticePaths(startX, startY, goalX, goalY, cache) {
  if (startX > goalX || startY > goalY) {
    return 0;
  }

  if (startX === goalX && startY === goalY) {
    return 1;
  }

  const key = `${startX}|${startY}`;
  if (cache.hasOwnProperty(key)) {
    return cache[key];
  }

  let numPaths = 0;

  // Move to the right
  numPaths += countLatticePaths(startX + 1, startY, goalX, goalY, cache);

  // Move down
  numPaths += countLatticePaths(startX, startY + 1, goalX, goalY, cache);

  cache[key] = numPaths;
  return cache[key]
}

function latticePaths(gridSize) {
  let cache = {}
  return countLatticePaths(0, 0, gridSize, gridSize, cache);
}

console.log(latticePaths(20));
