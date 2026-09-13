func findMin(nums []int) int {

	l, r:= 0, len(nums)
	mid := len(nums) / 2


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

	l, r = 0, len(nums) - 1
	if nums[l] > nums[r] {
		return nums[mid + 1];
	} else {
		return nums[0]
	}

	// l, r = 0, min(mid + 1, len(nums))
	
	// for l < r {
	// 	mid = (r + l) / 2

	// 	if nums[mid] == target {
	// 		return mid
	// 	} else if nums[mid] > target {
	// 		r = mid
	// 	} else {
	// 		l = mid + 1
	// 	}
	// } 

	// return -1

	// l, r := 0, len(nums) - 1

	// for l < r {
	// 	mid := (l + r) / 2
	// 	fmt.Println(l, mid, r)

	// 	if nums[mid] > nums[r] {
	// 		if nums [mid] > nums[l] {
	// 			l = mid + 1
	// 		} else {
	// 			r = mid
	// 		}
	// 	} else if nums[mid] > nums[l]{
	// 		l = mid + 1
	// 	} else {
	// 		return nums[mid]
	// 	}
	// }
	// return -1
}


func search(nums []int, target int) int {

	l, r:= 0, len(nums)
	mid := len(nums) / 2


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
	
	for l < r {
		mid = (r + l) / 2

		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			r = mid
		} else {
			l = mid + 1
		}
	} 

	return -1

}
