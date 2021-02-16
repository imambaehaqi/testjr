package main

import (
	"fmt"
	"sort"
	"strconv"
)

// create category structure
type category struct {
	code   int
	name   string
	unique string
}

// add category data
var categories = []category{
	{
		code:   6,
		name:   "Applied Sciences",
		unique: "(600)",
	},
	{
		code:   7,
		name:   "Arts",
		unique: "(700)",
	},
	{
		code:   0,
		name:   "General",
		unique: "(000)",
	},
	{
		code:   9,
		name:   "Geography, History",
		unique: "(900)",
	},
	{
		code:   4,
		name:   "Language",
		unique: "(400)",
	},
	{
		code:   8,
		name:   "Literature",
		unique: "(800)",
	},
	{
		code:   1,
		name:   "Philosophy, Psychology",
		unique: "(100)",
	},
	{
		code:   2,
		name:   "Religion",
		unique: "(200)",
	},
	{
		code:   5,
		name:   "Science, Math",
		unique: "(500)",
	},
	{
		code:   3,
		name:   "Social Sciences",
		unique: "(300)",
	},
}

// add inputs
var inputs = []string{"3A13", "5X19", "9Y20", "2C18", "1N20", "3N20", "1M21", "1F14", "9A21", "3N21", "0E13", "5G14", "8A23", "9E22", "3N14"}

var result []string
var value string

func main() {

	// loop data categories
	for i := 0; i < len(categories); i++ {

		var num []int
		// fmt.Println(categories[i])
		// fmt.Println("")

		// loop data inputs
		for x := 0; x < len(inputs); x++ {

			// took code category from inputs data
			code := inputs[x][:1]
			// took height from inputs
			height := inputs[x][2:]

			// fmt.Println(inputs[x])
			// fmt.Println(height)

			// parse code string to integer
			if codes, err := strconv.Atoi(code); err == nil {
				// parse height string to integer
				if heights, err := strconv.Atoi(height); err == nil {
					if categories[i].code == codes {
						num = append(num, heights)
					}
				}
			}
		}

		sort.Slice(num, func(a, b int) bool {
			return num[a] > num[b]
		})

		for _, v := range num {
			for a := 0; a < len(inputs); a++ {

				// took code category from inputs data
				code := inputs[a][:1]
				// took height from inputs
				height := inputs[a][2:]

				// parse code string to integer
				if codes, err := strconv.Atoi(code); err == nil {
					// parse height string to integer
					if heights, err := strconv.Atoi(height); err == nil {
						if categories[i].code == codes {
							if v == heights {
								// fmt.Printf("v = %v, a = %v # ", v, inputs[a])
								result = append(result, inputs[a])
							}
						}
					}
				}
			}
		}
	}

	// show result sorting category
	for _, res := range result {
		// fmt.Printf("%v ", res)

		value += res + " "
	}

	fmt.Println("unsorted category : ", inputs)
	fmt.Println("sorted category : ", value)
}
