function isLeapYear(year) {
  return year % 4 === 0 && (year % 100 !== 0 || year % 400 === 0);
}

function getFirstDayOfYear(year) {
  if (year === 1900) {
    return 0;
  }
  
  let currentYear = 1900;
  let firstDayOfYear = 0;
  while (currentYear < year) {
    const lastDayOfYear = isLeapYear(currentYear)
      ? firstDayOfYear + 1
      : firstDayOfYear;
    firstDayOfYear = (lastDayOfYear + 1) % 7;
    currentYear++;
  }

  return firstDayOfYear;
}

function countSundaysPerYear(year) {
  let numSundays = 0;

  const daysPerMonth = {
    1: 31,
    2: 28,
    3: 31,
    4: 30,
    5: 31,
    6: 30,
    7: 31,
    8: 31,
    9: 30,
    10: 31,
    11: 30,
    12: 31
  };

  const leapYear = isLeapYear(year);
  const numDays = leapYear ? 366 : 365;

  if (leapYear) {
    daysPerMonth[2]++;
  }

  const firstOfMonthDays = new Set();
  let day = 1;
  firstOfMonthDays.add(day);
  for (let month = 1; month <= 11; month++) {
    day += daysPerMonth[month];
    firstOfMonthDays.add(day);
  }

  const firstDayOfYear = getFirstDayOfYear(year);
  let sunday = 1 + (6 - firstDayOfYear);

  while (sunday <= numDays) {
    if (firstOfMonthDays.has(sunday)) {
      numSundays++;
    }
    sunday += 7;
  }


  return numSundays;
}

function countSundaysBetweenYears(firstYear, lastYear) {
  let startYear = firstYear;
  let result = 0;
  while (startYear <= lastYear) {
    result += countSundaysPerYear(startYear);
    startYear++;
  }

  return result;
}

console.log(countSundaysBetweenYears(1943, 1946));
console.log(countSundaysBetweenYears(1995, 2000));
console.log(countSundaysBetweenYears(1901, 2000));
