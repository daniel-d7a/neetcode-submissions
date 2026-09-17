import "slices"

func minEatingSpeed(piles []int, h int) int {

	maxVal := slices.Max(piles)

	l, r := 1, maxVal

	mid := (l + r) / 2;
	result := 0

	for l <= r {
		if ok := check(piles, h, mid); ok {
			result = mid
			r = mid - 1
		} else {
			l = mid + 1
		}
		mid = (l + r) / 2;
	}

	return result
}

func check(piles []int, h int, k int) bool {
	result := 0
	for _, p := range piles {
		result = result + int(math.Ceil(float64(p) / float64(k))) 
	}

	return result <= h
}