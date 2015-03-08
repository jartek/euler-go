package main

func NumberInLetters(num int) string {
	if num <= 10 {
		if num == 1 {
			return "one"
		}
		if num == 2 {
			return "two"
		}
		if num == 3 {
			return "three"
		}
		if num == 4 {
			return "four"
		}
		if num == 5 {
			return "five"
		}
		if num == 6 {
			return "six"
		}
		if num == 7 {
			return "seven"
		}
		if num == 8 {
			return "eight"
		}
		if num == 9 {
			return "nine"
		}
		if num == 10 {
			return "ten"
		}
	} else if num < 20 {
		if num == 11 {
			return "eleven"
		}
		if num == 12 {
			return "twelve"
		}
		if num == 13 {
			return "thirteen"
		}
		if num == 14 {
			return "fourteen"
		}
		if num == 15 {
			return "fifteen"
		}
		if num == 16 {
			return "sixteen"
		}
		if num == 17 {
			return "seventeen"
		}
		if num == 18 {
			return "eighteen"
		}
		if num == 19 {
			return "nineteen"
		}
	} else if num < 100 {
		n := num % 10
		if num < 30 {
			return "twenty" + NumberInLetters(n)
		}
		if num < 40 {
			return "thirty" + NumberInLetters(n)
		}
		if num < 50 {
			return "forty" + NumberInLetters(n)
		}
		if num < 60 {
			return "fifty" + NumberInLetters(n)
		}
		if num < 70 {
			return "sixty" + NumberInLetters(n)
		}
		if num < 80 {
			return "seventy" + NumberInLetters(n)
		}
		if num < 90 {
			return "eighty" + NumberInLetters(n)
		}
		if num < 100 {
			return "ninety" + NumberInLetters(n)
		}
	} else if num < 1000 {
		hundreds := num / 100
		remaining := num % 100
		if remaining == 0 {
			return NumberInLetters(hundreds) + "hundred"
		}
		return NumberInLetters(hundreds) + "hundredand" + NumberInLetters(remaining)
	}
	if num == 1000 {
		return "onethousand"
	}
	return ""
}

func NumberLetterCounts(num int) int {
	var sum int
	for i := 1; i <= num; i++ {
		sum = sum + len(NumberInLetters(i))
	}
	return sum
}
