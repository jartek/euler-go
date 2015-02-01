/*
  a² + b² = c²

  a + b + c = 1000 (given)

  c = 1000 - a - b

  a² + b² = 1000² + a² + b² - 2000a - 2000b + 2ab

  0 = 1000² - 2000(a+b) + 2ab

  2a(1000 - b) = 1000² - 2000b

  if num = 1000

  a(num - b) = (num * num)/2 - (num * b)

  a = ((num * num)/2 - (num * b))/(num - b)

*/

package main

import (
	"fmt"
	"math"
)

func PythagoreanTriplet(num float64) int64 {
	var a, b float64
	for b = num/2 - 1; b > 0; b-- {
		a = ((num*num)/2 - (num * b)) / (num - b)
		if a == math.Floor(a) {
			fmt.Println(a, b)
			return int64(a * b * (num - a - b))
		}
	}
	return -1
}
