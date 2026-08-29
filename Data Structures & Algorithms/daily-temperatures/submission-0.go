func dailyTemperatures(temperatures []int) []int {

	result := make([]int, len(temperatures))
	stack := []Item{}

	for i, t := range temperatures {

		if len(stack) > 0 {
			top := stack[len(stack) - 1]
			for top.temp < t {
				result[top.pos] = i - top.pos
				stack = stack[: len(stack) - 1]
				if len(stack) > 0 {
					top = stack[len(stack) - 1]
				}else{
					break
				}
			}
			stack = append(stack, Item{t, i})
		} else {
			stack = append(stack, Item{t, i})
		}
	}
	return result
}

type Item struct {
	temp int
	pos int
}