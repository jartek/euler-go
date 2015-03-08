package main

var cache = map[int]int{}

func collatz(num int) int {
	if num%2 != 0 {
		return (num * 3) + 1
	} else {
		return num / 2
	}
}

func NumInCache(num int) int {
	_, ok := cache[num]
	if !ok {
		return -1
	}
	return cache[num]
}

func CacheNum(num int, length int) {
	cache[num] = length
}

func CollatzSequenceLength(num int) int {
	length := 0
	for num != 1 {
		cachedVal := NumInCache(num)
		if cachedVal == -1 {
			num = collatz(num)
			length = length + 1
		} else {
			length = length + cachedVal
			num = 1
		}
	}
	return length
}

func LongestCollatzSequence(num int) int {
	longest, largest, length := 0, 0, 0
	for i := 2; i <= num; i++ {
		length = CollatzSequenceLength(i)
		CacheNum(i, length)
		if largest < length {
			longest = i
			largest = length
		}
	}
	return longest
}
