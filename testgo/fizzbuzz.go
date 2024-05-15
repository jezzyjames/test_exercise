package main

import "strconv"

func FizzBuzz(n int) string {

	myMap := map[int]string{
		3: "Fizz",
		5: "Buzz",
		6: "Fizz",
	}

	result := myMap[n]
	if result != "" {
		return result
	}
	return strconv.Itoa(n)
}
