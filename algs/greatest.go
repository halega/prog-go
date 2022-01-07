package algs

func GreatestNumber(nums []int) int {
	greatest := nums[0]
	for _, n := range nums {
		if n > greatest {
			greatest = n
		}
	}
	return greatest
}
