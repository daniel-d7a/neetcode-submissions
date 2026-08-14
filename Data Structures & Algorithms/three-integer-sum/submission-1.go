import "slices"

func threeSum(nums []int) [][]int {

	slices.Sort(nums)
	res := [][]int{}

	for i, val := range nums {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		l := i + 1
		r := len(nums) - 1
		total := 0

		for l < r {

			// fmt.Println("i", i)
			// fmt.Println("l", l)
			// fmt.Println("r", r)
			// fmt.Println("-----------")

			total = val + nums[l] + nums[r]

			if total == 0 {
				res = append(res, []int{nums[i], nums[l], nums[r]})
				r = r - 1
				l = l + 1

				for l < r && nums[l] == nums [l - 1] {
					l = l + 1
				}


			} else if total > 0 {
				r = r - 1
			} else if total < 0 {
				l = l + 1
			}
		}
	}
	return res

}