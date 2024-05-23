package main

import (
	"strconv"
)

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

// No IF
func FizzBuzz(n int) string {

	fbMap := map[int]map[int]string{
		0: {
			0: "f",
		},
		1: {
			0: "b",
		},
		2: {
			0: "z",
		},
	}

	keys := []int{3, 5, 15}

	fb := []byte("")
	for i, k := range keys {
		v := n % k

		r := fbMap[i][v]

		fb = append(fb, r...)
	}

	fbResult := map[string]string{
		"":  strconv.Itoa(n),
		"f": "Fizz",
		"b": "Buzz",
	}

	return fbResult[string(fb)]
}
