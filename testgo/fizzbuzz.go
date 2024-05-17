package main

import (
	"strconv"
	"strings"
)

func FizzBuzz(n int) string {

	result := []string{}

	fizzBuzzMap := map[int]string{
		3:  "Fizz",
		5:  "Buzz",
		15: "FizzBuzz",
	}

	for i := n; i <= n; i++ {
		value := strconv.Itoa(i)
		for key := range fizzBuzzMap {
			if i%key == 0 {
				value = fizzBuzzMap[key]
				continue

			}
		}

		result = append(result, value)
	}

	return strings.Join(result[:], ",")
}
