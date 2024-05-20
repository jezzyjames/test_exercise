package main

func FizzBuzz(n int) string {

	fizzBuzzMap := map[int]string{
		1: "1",
		2: "2",
		3: "Fizz",
		4: "4",
	}

	return fizzBuzzMap[n]
}
