package leetcode

func search704(nums []int, target int) int {
	min, max := 0, len(nums)-1 // 错误点1: 忘记减1

	for min <= max {
		mid := (min + max) / 2 // 可以优化成 mid = min + (max- min) << 1
		if nums[mid] > target {
			max = mid - 1
		} else if nums[mid] < target {
			min = mid + 1
		} else {
			return mid
		}
	}
	return -1
}
