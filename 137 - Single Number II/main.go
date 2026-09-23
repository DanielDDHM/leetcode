package main

func singleNumber(nums []int) int {
	var result int32 = 0
	for i := uint(0); i < 32; i++ {
		count := 0
		for _, num := range nums {
			n := int32(num)
			if (n>>i)&1 == 1 {
				count++
			}
		}
		if count%3 == 1 {
			result |= 1 << i
		}
	}
	return int(result)
}
