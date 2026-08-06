func twoSum(nums []int, target int) []int {
    
	valMap := make(map[int]int)

	for i, val := range nums {
		valMap[target - val] = i
	}

	for i, val := range nums {
		if idx, prs := valMap[val]; prs && i != idx {
			if i < idx {
				return []int{i, idx}
			}
			return []int{idx, i}
		}
	}
	return []int{}
}
