package main

func search(nums []int, target int) bool {
	l := 0
	r := len(nums) - 1
	for l <= r {
		site := l + (r-l)/2
		if nums[site] == target {
			return true
		}
		if nums[site] == nums[l] {
			l++
			continue
		}
		if nums[l] <= nums[site] {
			if target < nums[site] && target >= nums[l] {
				r = site - 1
			} else {
				l = site + 1
			}
		} else {
			if target > nums[site] && target <= nums[r] {
				l = site + 1
			} else {
				r = site - 1
			}
		}
	}
	return false
}
