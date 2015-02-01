package main

import "math"

func reverse(num float64) float64 {
	var tmp float64
	for num > 0 {
		tmp = tmp*10 + math.Floor(math.Mod(num, 10))
		num = math.Floor(num / 10)
	}
	return tmp
}

func isPalindrome(num float64) bool {
	return num == reverse(num)
}

func isDivisible(mul float64, num float64) bool {
	var i, j float64
	for i = num; i > 0; i-- {
		for j = num; j > 0; j-- {
			if mul > i*j {
				break
			}
			if i*j == mul {
				return true
			}
		}
	}
	return false
}

func LargestPalindrome(n float64) float64 {
	var i, num, max float64
	for i = 0; i < n; i++ {
		num = num + math.Pow(10, i)*9
	}
	max = num * num
	for max > 0 {
		if isPalindrome(max) && isDivisible(max, num) {
			return max
		}
		max = max - 1
	}

	return -1
}
