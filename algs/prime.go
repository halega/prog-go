package algs

// IsPrimeSimple checks if a number is a prime number.
//
// It implements straightforward algorithm. Time complexity: O(N).
func IsPrimeSimple(n int) bool {
	for i := 2; i < n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}
