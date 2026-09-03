func search(nums []int, target int) int {

	mid := len(nums) / 2
	l := 0
	r := len(nums) - 1

	for l <= r {

//		fmt.Println(l, " - ", r, " - ", mid)

		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			r = mid - 1
		} else {
			l = mid + 1
		}

		mid = (r + l) / 2
	} 

	return -1
}
