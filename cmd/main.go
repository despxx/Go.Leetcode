package main

import "fmt"

func main() {
	a := twoSum([]int{2, 7, 11, 15}, 9)
	fmt.Println(a)
}

func twoSum(nums []int, target int) []int {
	numMap := make(map[int]int)

	for i, num := range nums {
		complement := target - num
		if j, exists := numMap[complement]; exists {
			return []int{j, i}
		}
		numMap[num] = i
	}
	return nil
}
