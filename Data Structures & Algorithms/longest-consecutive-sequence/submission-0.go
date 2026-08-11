	func longestConsecutive(nums []int) int {

		hash := make(map[int]bool)

		for _, val := range nums {
			hash[val] = true
		}

		longestSequence := 0
		for num := range hash {

			if hash[num - 1] {
				continue
			}

			lenSequence := 1
			next := num + 1
			for hash[next] {
				lenSequence = lenSequence + 1
				next = next + 1
			}

			if lenSequence > longestSequence {
				longestSequence = lenSequence
			}
		}

		return longestSequence

	}
