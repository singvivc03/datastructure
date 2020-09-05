package dp

// Fibonacci function calculates the nth fibonacci number
// Time complexity ø(n) and space complexity O(n^2)
func Fibonacci(n int) int64 {
	memo := make([]int64, n+1)
	for index := range memo {
		memo[index] = int64(-1)
	}
	return calculateFibonacci(n, memo)
}

func calculateFibonacci(n int, memo []int64) int64 {
	if memo[n] == -1 {
		res := int64(0)
		if n == 0 || n == 1 {
			res = int64(n)
		} else {
			res = calculateFibonacci(n-1, memo) + calculateFibonacci(n-2, memo)
		}
		memo[n] = res
	}
	return memo[n]
}

// FibonacciTabular func calculates the fibonacci of nth number in a bottom up approach
// Time complexity ø(n) and space complexity is ø(n)
func FibonacciTabular(n int) int64 {
	if n == 0 || n == 1 {
		return int64(n)
	}
	memo := make([]int64, n+1)
	memo[0], memo[1] = int64(0), int64(1)
	for i := 2; i <= n; i++ {
		memo[i] = memo[i-1] + memo[i-2]
	}
	return memo[n]
}

// FibonacciWithConstantSpace func calculates the fibonacci of nth number
// Time complexity ø(n) and space complexity ø(1)
func FibonacciWithConstantSpace(n int) int64 {
	if n == 0 || n == 1 {
		return int64(n)
	}
	f0, f1, res := int64(0), int64(1), int64(0)
	for i := 2; i <= n; i++ {
		res = f0 + f1
		f0 = f1
		f1 = res
	}
	return res
}
