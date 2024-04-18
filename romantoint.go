package main

import "fmt"

// I             1
// V             5
// X             10
// L             50
// C             100
// D             500
// M             1000
func romanToInt(s string) int {
	symbols := make(map[string]int)
	symbols["I"] = 10
	symbols["V"] = 5
	symbols["X"] = 10
	symbols["L"] = 50
	symbols["C"] = 100
	symbols["D"] = 500
	symbols["M"] = 1000

	for i, x := range s {
		fmt.Println(x, i)
	}

	return 1
}

func main() {
	result := romanToInt("III")
	fmt.Println(result)
}
