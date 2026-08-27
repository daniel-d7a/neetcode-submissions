func evalRPN(tokens []string) int {
	stack := []int{}

	operationMap := map[string]bool{
		"+":true,
		"-":true,
		"*":true,
		"/":true,
	}

	for _, t := range tokens {
		x := string(t)
		
		if ok := operationMap[x]; ok {
			first := stack[len(stack) - 1]
			second := stack[len(stack) - 2]
			stack = stack[: len(stack) - 2]

			var result int

			if x == "+" {
				result = second + first
			} else if x == "-" {
				result = second - first
			} else if x == "*" {
				result = second * first
			} else if x == "/" {
				result = second / first
			}

			stack = append(stack, result)
		}else{
			val, err := strconv.Atoi(x)
			if err != nil {
				panic("error")
			}
			stack = append(stack, val)
		}
	}
	return stack[0]
}
