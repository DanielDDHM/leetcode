package main

func maximumGap(nums []int) int {
	if len(nums) < 2 {
		return 0
	}

	minNum, maxNum := nums[0], nums[0]
	for _, num := range nums {
		if num < minNum {
			minNum = num
		}
		if num > maxNum {
			maxNum = num
		}
	}

	if minNum == maxNum {
		return 0
	}

	gap := maxNum - minNum
	bucketSize := gap / len(nums)
	if bucketSize == 0 {
		bucketSize = 1
	}

	bucketCount := (gap + bucketSize - 1) / bucketSize

	type bucket struct {
		min int
		max int
	}
	buckets := make([]bucket, bucketCount)
	for i := range buckets {
		buckets[i].min = -1
		buckets[i].max = -1
	}

	for _, num := range nums {
		idx := (num - minNum) / bucketSize
		if idx >= bucketCount {
			idx = bucketCount - 1
		}
		if buckets[idx].min == -1 {
			buckets[idx].min = num
			buckets[idx].max = num
		} else {
			if num < buckets[idx].min {
				buckets[idx].min = num
			}
			if num > buckets[idx].max {
				buckets[idx].max = num
			}
		}
	}

	maxGap := 0
	prevMax := minNum
	for i := 0; i < bucketCount; i++ {
		if buckets[i].min == -1 {
			continue
		}
		if buckets[i].min > prevMax {
			gap := buckets[i].min - prevMax
			if gap > maxGap {
				maxGap = gap
			}
		}
		prevMax = buckets[i].max
	}

	return maxGap
}
