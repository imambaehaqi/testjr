package main

import (
	"fmt"
	"testing"
)

func TestInput1(t *testing.T) {
	input := isPalindrome(1, 10)

	if input == 9 {
		fmt.Println("unit test 1 with result :", input)
	} else {
		t.Logf("Got %d expect %d", input, 9)
		t.Fail()
	}
}

func TestInput2(t *testing.T) {
	input := isPalindrome(99, 100)

	if input == 1 {
		fmt.Println("unit test 2 with result :", input)
	} else {
		t.Logf("Got %d expect %d", input, 1)
		t.Fail()
	}
}

func TestInput3(t *testing.T) {
	input1 := isPalindrome(21, 31)

	if input1 == 1 {
		fmt.Println("unit test 3 with result :", input1)
	} else {
		t.Logf("Got %d expect %d", input1, 1)
		t.Fail()
	}
}
