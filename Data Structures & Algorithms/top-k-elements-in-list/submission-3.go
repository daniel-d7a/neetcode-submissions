import ("slices")

func topKFrequent(nums []int, k int) []int {
	itemCounts := make(map[int]int)
	freqArray := make([][]int, len(nums)+1)

	for _, val := range nums {
		itemCounts[val]++
	}

	for val, freq:= range itemCounts {
		freqArray[freq] = append(freqArray[freq], val)
	}

	result := []int{}

	for _, val := range slices.Backward(freqArray) {
		result = append(result, val...)
		if len(result) >= k {
			break
		}
	}
	return result[:k]
}
