package main

func removeMultiples(arr []bool, n int) []bool {
	i := 2
	num := i * n
	for num < len(arr) {
		arr[num] = false
		i = i + 1
		num = n * i
	}
	return arr
}

func prepareSieve(arr []bool) []bool {
	for i := 0; i < len(arr); i++ {
		if arr[i] == true {
			arr = removeMultiples(arr, i)
		}
	}
	return arr
}

func SumOfPrimes(limit float64) float64 {
	var sum float64
	var arr []bool = make([]bool, int(limit))

	for i := 2; i < len(arr); i++ {
		arr[i] = true
	}

	arr = prepareSieve(arr)
	for j := 0; j < int(limit); j++ {
		if arr[j] == true {
			sum = sum + float64(j)
		}
	}
	return sum
}
