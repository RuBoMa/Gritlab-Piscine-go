package piscine

func Abort(a, b, c, d, e int) int {
	nums := []int{a, b, c, d, e}

	for i := 0; i < len(nums)-1; i++ {
		for j := 0; j < len(nums)-i-1; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
	}

	middleIndex := len(nums) / 2

	if len(nums)%2 == 1 {
		return nums[middleIndex]
	} else {
		return (nums[middleIndex-1] + nums[middleIndex]) / 2
	}
}
