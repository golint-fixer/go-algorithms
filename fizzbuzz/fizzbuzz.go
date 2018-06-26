package fizzbuzz

import "strconv"

func FizzBuzz(value int) string {
	if value%3 == 0 && value%5 == 0 {
		return "FizzBuzz"
	}

	if value%3 == 0 {
		return "Fizz"
	}

	if value%5 == 0 {
		return "Buzz"
	}

	return strconv.Itoa(value)
}
