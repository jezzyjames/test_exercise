package main

import "strconv"

func FizzBuzz(n int) string {

	myMap := map[int]string{
		3: "Fizz",
		5: "Buzz",
	}

	result := myMap[n]
	if result != "" {
		return result
	}
	return strconv.Itoa(n)
}
