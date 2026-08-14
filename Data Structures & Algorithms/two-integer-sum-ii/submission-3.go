func twoSum(numbers []int, target int) []int {
	i := 0
	j := len(numbers) - 1

	for i < j{
		total := numbers[i] + numbers[j]

		if total == target{
			return []int{i + 1, j + 1}
		}

		if total > target {
			j = j - 1
		}

		if total < target {
			i = i + 1
		}

	}
	return []int{}
}
