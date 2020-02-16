var fs = require("fs");

function getNameScore(name) {
  let score = 0;
  for (let i = 0; i < name.length; i++) {
    score += name.charCodeAt(i) - 64;
  }

  return score;
}

function namesScores(arr) {
  arr.sort();

  let score = 0;
  for (let i = 0; i < arr.length; i++) {
    score += getNameScore(arr[i]) * (i + 1);
  }

  return score;
}

const test = fs
  .readFileSync("p022_names.txt")
  .toString()
  .split(",")
  .map(name => name.replace(/"/g, ""));

console.log(namesScores(test));
