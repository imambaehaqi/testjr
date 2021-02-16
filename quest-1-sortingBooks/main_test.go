package main

import (
	"fmt"
	"testing"
)

func TestMain(t *testing.T) {
	main()

	var expect string = "0E13 9E22 9A21 9Y20 8A23 1M21 1N20 1F14 2C18 5X19 5G14 3N21 3N20 3A13"

	if value == expect {
		fmt.Println(value)
	} else {
		t.Logf("result : %v", value)
		t.Logf("expect : %v", expect)

		t.Fail()
	}
}
