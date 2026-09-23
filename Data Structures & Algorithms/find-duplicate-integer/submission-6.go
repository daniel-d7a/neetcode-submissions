func findDuplicate(nums []int) int {
    
	slow := 0
	fast := 0

	for {
		// fmt.Println("slow: ", nums[slow])
		// fmt.Println("fast: ", nums[fast])

		slow = nums[slow]
		fast = nums[nums[fast]]

		if slow == fast {


			fmt.Println("break")

			newSlow := 0

			for {
				slow = nums[slow]
				newSlow = nums[newSlow]

				if slow == newSlow {
					return slow
				}
			}

			return nums[0]
		}
	}
	return nums[0]
}
