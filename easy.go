package goleetcode

/*
1. Two Sum
https://leetcode.com/problems/two-sum/
*/

func twoSum(nums []int, target int) []int {
	numMap := make(map[int]int)

	for i, num := range nums {
		compl := target - num
		if j, exists := numMap[compl]; exists {
			return []int{j, i}
		}
		numMap[num] = i
	}
	return nil
}
