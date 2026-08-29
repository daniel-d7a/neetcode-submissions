func dailyTemperatures(temperatures []int) []int {

	result := make([]int, len(temperatures))
	stack := []int{}

	for i, t := range temperatures {

		if len(stack) > 0 {
			top := stack[len(stack) - 1]
			for temperatures[top] < t {
				result[top] = i - top
				stack = stack[: len(stack) - 1]
				if len(stack) > 0 {
					top = stack[len(stack) - 1]
				}else{
					break
				}
			}
			stack = append(stack, i)
		} else {
			stack = append(stack, i)
		}
	}
	return result
}