package main

func Multiples(n int) int {
	var sum int
	for i := 0; i < n; i++ {
		if i%3 == 0 || i%5 == 0 {
			sum = sum + i
		}
	}
	return sum
}
