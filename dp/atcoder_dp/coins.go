package dp

// https://atcoder.jp/contests/dp/tasks/dp_i
func probabilityWithMoreHeads(n int, probabilities []float64) float64 {
	dp := make([]float64, n+1)
	dp[0] = 1.0
	for coins := 0; coins < n; coins++ {
		for j := coins + 1; j >= 0; j-- {
			if j == 0 {
				dp[j] = 0.0 + dp[j]*(1-probabilities[coins])
			} else {
				dp[j] = dp[j-1]*probabilities[coins] + dp[j]*(1-probabilities[coins])
			}
		}
	}
	answer := 0.0
	for heads := 0; heads <= n; heads++ {
		tails := n - heads
		if heads > tails {
			answer += dp[heads]
		}
	}
	return answer
}
