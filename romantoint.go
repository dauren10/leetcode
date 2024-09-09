package main

import "fmt"

func main() {
	//res := romanToInt("MCMXCIV")

	res := romanToInt("III")

	fmt.Println(res)
}

// MCMXCIV
// 5-1+100-10+1000-100+1000

// Input: s = "MCMXCIV"
// Output: 1994
func romanToInt(s string) int {
	translate := make(map[string]int)
	translate["I"] = 1
	translate["V"] = 5
	translate["X"] = 10
	translate["L"] = 50
	translate["C"] = 100
	translate["D"] = 500
	translate["M"] = 1000

	translate["CM"] = 900
	translate["XC"] = 90
	translate["IV"] = 4
	var res int
	var sum string

	for i := 0; i < len(s); i++ {

		current := string(s[i])
		if i < len(s)-1 {
			next := string(s[i+1])
			sum = current + next
		}

		val, ok := translate[sum]
		if ok {
			res += val
			i++
		} else {
			res += translate[current]
		}

	}

	return res

}

func romanToInt2(s string) int {
	if current < prev {
		total -= current
	} else {
		total += current
	}
}
