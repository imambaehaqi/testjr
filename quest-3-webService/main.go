package main

import (
	"fmt"
	"net/http"
)

func testServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "total palindrome: %v ", isPalindrome(1, 100))
}

func isPalindrome(a, b int) (res int) {
	var totalPalindrome int
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
	fmt.Println("listen to localhost:8080")
	http.HandleFunc("/", testServer)
	http.ListenAndServe(":8080", nil)
}
