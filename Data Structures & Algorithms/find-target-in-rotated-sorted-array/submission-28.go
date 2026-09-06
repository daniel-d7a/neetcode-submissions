func search(nums []int, target int) int {

	l, r:= 0, len(nums)
	mid := len(nums) / 2

	if nums[l] > nums[r - 1]{

		for l < r {
			if mid > 0 && mid < len(nums) - 1 && nums[mid] > nums[mid - 1] && nums[mid] > nums[mid + 1] {
				break
			} else if nums[mid] > nums[l] {
				l = mid + 1
			} else {
				r = mid
			}
			mid = (r + l) / 2
		} 
		l, r = 0, len(nums)

		if target >= nums[l] {
			r = min(mid + 1, len(nums))
		} else {
			l = mid
		}
	} else {
		r = len(nums)
	}
	fmt.Println("mid: ", mid)
	fmt.Println("l: ", l)
	fmt.Println("r: ", r)

	if nums[mid] == target {
		return mid
	} 
	for l < r {
		mid = ((r + l) / 2)

		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			r = mid
		} else {
			l = mid + 1
		}
	} 

	if nums[mid] == target {
		return mid
	}
	return -1

}
