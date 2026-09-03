func searchMatrix(matrix [][]int, target int) bool {

	mid := len(matrix) / 2
	l := 0
	r := len(matrix) - 1

	for l <= r {
		if matrix[mid][0] <= target && matrix[mid][len(matrix[mid]) - 1] >= target {
			return binarySearch(matrix[mid], target)
		} else if matrix[mid][0] > target {
			r = mid - 1
		} else {
			l = mid + 1
		}

		mid = (r + l) / 2
	} 

	return false

}

func binarySearch(nums []int, target int) bool {

	mid := len(nums) / 2
	l := 0
	r := len(nums) - 1

	for l <= r {
		if nums[mid] == target {
			return true
		} else if nums[mid] > target {
			r = mid - 1
		} else {
			l = mid + 1
		}

		mid = (r + l) / 2
	} 

	return false
}
