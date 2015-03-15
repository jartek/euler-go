package main

func isLeapYear(year int, month int) bool {
	return month == 1 && (year%4 == 0 && (year%100 != 0 || year%400 == 0))
}

func FirstDayOfTheMonth(month int, year int, day int) int {
	var endDay int

	daysInMonth := [12]int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}

	endDay = day + daysInMonth[month]
	endDay = endDay % 7

	if isLeapYear(year, month) {
		endDay = (endDay + 1) % 7
	}

	return endDay
}

func FirstDayOfTheYear(year int) int {
	if year == 1900 {
		return 1
	}

	numOfDaysInAYear := 365
	if isLeapYear(year-1, 1) {
		numOfDaysInAYear++
	}
	return (FirstDayOfTheYear(year-1) + numOfDaysInAYear) % 7
}

func SundaysInYear(year int) int {
	var count int

	day := FirstDayOfTheYear(year)

	if day == 0 {
		count++
	}

	for i := 0; i < 11; i++ {
		day = FirstDayOfTheMonth(i, year, day)
		if day == 0 {
			count++
		}
	}

	return count
}

func NumberOfSundays(r []int) int {
	var count int
	for i := r[0]; i <= r[1]; i++ {
		count += SundaysInYear(i)
	}
	return count
}
