package main

import (
	"sort"
	"strconv"
)

func FizzBuzz(n int) string {

	fizzBuzzMap := map[int]string{
		3: "Fizz",
		5: "Buzz",
	}

	keys := []int{}
	for key := range fizzBuzzMap {
		keys = append(keys, key)
	}
	sort.Ints(keys)

	for _, k := range keys {
		if n%k == 0 {
			return fizzBuzzMap[n]
		}
	}

	return strconv.Itoa(n)
}
