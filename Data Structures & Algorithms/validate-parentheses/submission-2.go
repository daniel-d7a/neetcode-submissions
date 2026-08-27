func isValid(s string) bool {
	stack := []string{}

	for _, c := range s {
		
		x := string(c)
		
		if x == "(" || x == "{" || x == "[" {
			stack = append(stack, x)
			continue
		} else if len(stack) > 0 {
			top := stack[len(stack) - 1]

			if (top == "{" && x == "}") || (top == "(" && x == ")") || (top == "[" && x == "]") {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		} else {
			return false
		}
	}

	return len(stack) == 0
}
