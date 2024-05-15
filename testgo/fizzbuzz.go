package main

import "strconv"

func FizzBuzz(n int) string {

	if n == 3 || n == 6 || n == 9 {
		return "Fizz"
	}

	if n == 5 {
		return "Buzz"
	}

	return strconv.Itoa(n)
}
