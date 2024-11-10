package main

func twoSum(nums []int, target int) []int {
	compMap := make(map[int]int)
	for i, num := range nums {
		comp := target - num
		if index, found := compMap[comp]; found {
			return []int{index, i}
		}
		compMap[num] = i
	}
	return nil
}
