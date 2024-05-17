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
	for k := range fizzBuzzMap {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	value := strconv.Itoa(n)

	for _, key := range keys {
		if n%key == 0 {
			value = fizzBuzzMap[key]
		}
	}

	return value
}
