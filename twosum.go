package main

import "fmt"

func main() {

	ar := []int{2, 7, 11, 15}
	res := twoSum(ar, 9)
	fmt.Println(res)
}

func twoSum(nums []int, target int) []int {
	obj := make(map[int]int)
	res := []int{}

	for i, num := range nums {
		diff := target - num

		if index, ok := obj[diff]; ok && index != i {

			res = append(res, obj[diff], i)
			fmt.Println("Work append", diff, index, i)
		} else {
			fmt.Println("index", index, diff)
		}
		obj[num] = i
	}

	return res
}
