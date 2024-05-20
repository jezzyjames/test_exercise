package main

import (
	"sort"
	"strconv"
)

func FizzBuzz(n int) string {

	fizzBuzzMap := map[int]string{
		3:  "Fizz",
		5:  "Buzz",
		15: "FizzBuzz",
	}

	keys := []int{}
	for key := range fizzBuzzMap {
		keys = append(keys, key)
	}
	sort.Ints(keys)

	result := strconv.Itoa(n)
	for _, k := range keys {
		if n%k == 0 {
			result = fizzBuzzMap[k]
		}
	}

	return result
}
