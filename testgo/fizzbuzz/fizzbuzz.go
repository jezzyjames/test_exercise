package main

import (
	"strconv"
)

// No IF
func FizzBuzz(n int) string {

	fbMap := map[int]map[int]string{
		0: {
			0: "Fizz",
		},
		1: {
			0: "Buzz",
		},
	}

	keys := []int{3, 5}

	fb := []byte("")
	for i, k := range keys {
		v := n % k

		r := fbMap[i][v]

		fb = append(fb, r...)
	}

	fbResult := map[string]string{
		"":         strconv.Itoa(n),
		"Fizz":     "Fizz",
		"Buzz":     "Buzz",
		"FizzBuzz": "FizzBuzz",
	}

	return fbResult[string(fb)]
}

// // 1 IF
// func FizzBuzz(n int) string {

// 	fizzBuzzMap := map[int]string{
// 		3:  "Fizz",
// 		5:  "Buzz",
// 		15: "FizzBuzz",
// 	}

// 	keys := []int{}
// 	for key := range fizzBuzzMap {
// 		keys = append(keys, key)
// 	}
// 	sort.Ints(keys)

// 	result := strconv.Itoa(n)
// 	for _, k := range keys {
// 		if n%k == 0 {
// 			result = fizzBuzzMap[k]
// 		}
// 	}

// 	return result
// }
