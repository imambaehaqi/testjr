package main

import "fmt"

// set variable for get 2 number input
var number1 = 21
var number2 = 31

// set variable total palindrome
var totalPalindrome int

func isPalindrome(a, b int) (res int) {
	// check if first number is < than second number
	if a < b {
		// looping numbers based on user input
		for i := a; i < b; i++ {
			// set value for oriNumber and revNumber
			oriNumber := i
			revNumber := 0

			num := oriNumber

			for num > 0 {
				var remainder = num % 10
				revNumber = (revNumber * 10) + remainder
				num = num / 10
			}

			// the process of checking the palindrome number
			if oriNumber == revNumber {
				totalPalindrome++
			}
		}

		res = totalPalindrome

	} // else {
	// 	fmt.Printf("cannot process palindrome numbers between %d and %d", a, b)
	// 	fmt.Print("\n")
	// }

	return res
}

func main() {

	// set input value for number1
	fmt.Print("input number1: \n")
	fmt.Scan(&number1)

	// set input value for number2
	fmt.Print("input number2: \n")
	fmt.Scan(&number2)

	// showing total palindrome number between a and b
	fmt.Printf("%v < %v", number1, number2)
	fmt.Println(" ")
	fmt.Println("total palindrome: ", isPalindrome(number1, number2))
}
