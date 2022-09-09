package algorithms

func searchRange(nums []int, target int) []int {
	l, r := 0, len(nums)-1
	for l <= r {
		m := (l + r) / 2
		if nums[m] < target {
			l = m + 1
		} else if nums[m] > target {
			r = m - 1
		} else { // narrow down the region
			if nums[l] < target {
				l++
			}
			if nums[r] > target {
				r--
			}
			if nums[l] == nums[r] {
				return []int{l, r}
			}
		}
	}

	return []int{-1, -1}
}
