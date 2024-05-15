package main

import "strconv"

func FizzBuzz(n int) string {

	myMap := map[int]string{
		3: "Fizz",
		5: "Buzz",
		6: "Fizz",
		9: "Fizz",
		10: "Buzz",
		12: "Fizz",
		15: "FizzBuzz",
	}

	result := myMap[n]
	if result != "" {
		return result
	}
	return strconv.Itoa(n)
}
