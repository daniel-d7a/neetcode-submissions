func maxProfit(prices []int) int {

	l := 0
	r := 1

	maxP := 0

	for r < len(prices) {
		if prices[r] > prices[l] {
			maxP = max(maxP, prices[r] - prices[l])
		} else {
			l = r
		}
		r++
	}

	return maxP;
}
