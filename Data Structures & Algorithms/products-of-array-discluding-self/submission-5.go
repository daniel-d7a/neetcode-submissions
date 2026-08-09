import ("slices")

func productExceptSelf(nums []int) []int {

	left := make([]int, len(nums))
	right := make([]int, len(nums))

	for i := range nums{
		if i == 0{
			left[i] = 1
		}else {
			left[i] = left[i-1] * nums[i-1]
		}
	}

	for i := range slices.Backward(nums){
		if i == len(nums)-1{
			right[i] = 1
		}else {
			right[i] = right[i+1] * nums[i+1]
		}
	}

	result := []int{}

	for i := range nums {
		result = append(result, left[i] * right[i])
	}
	return result
}
