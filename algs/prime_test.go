package algs

import "testing"

func TestIsPrimeSimpleTrue(t *testing.T) {
	primes := []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41}
	for _, n := range primes {
		if !IsPrimeSimple(n) {
			t.Error(n, "is prime but IsPrimeSimple returns false")
		}
	}
}

func TestIsPrimeSimpleFalse(t *testing.T) {
	nonPrimes := []int{4, 6, 8, 9, 10, 12, 14, 15, 16, 18, 20, 21, 22, 24, 25, 26, 27, 28}
	for _, n := range nonPrimes {
		if IsPrimeSimple(n) {
			t.Error(n, "is not prime but IsPrimeSimple returns true")
		}
	}
}
