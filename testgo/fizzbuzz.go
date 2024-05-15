package main

import "strconv"

func FizzBuzz(n int) string {

	if n % 3 == 0 {
		return "Fizz"
	}

	if n == 5 || n == 10 {
		return "Buzz"
	}

	return strconv.Itoa(n)
}
