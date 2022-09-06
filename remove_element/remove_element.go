package remove_element

func removeElement(nums []int, val int) int {
	var index int
	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			nums[index] = nums[i]
			index++
		}
	}
	return index
}
