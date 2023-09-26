package main

import "testing"

func TestSearch(t *testing.T) {
	nums := []int{1, 0, 1, 1, 1, 1}
	target := 0
	found := search(nums, target)
	if !found {
		t.Errorf("get %v", found)
	}

}
