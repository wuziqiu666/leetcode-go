package main

import (
	"testing"
)

func TestDeleteDuplicatess(t *testing.T) {
	cases := []struct {
		list []int
		want string
	}{
		{list: []int{1, 2, 3, 3, 4, 4, 5}, want: "[1, 2, 5]"},
		{list: []int{1, 1}, want: "[]"},
		{list: []int{0, 1, 2, 2, 3, 4}, want: "[0, 1, 3, 4]"},
	}
	for _, c := range cases {
		head := NewListNode(c.list)
		head = deleteDuplicates(head)
		got := head.String()
		want := c.want
		if got != want {
			t.Fatalf("got %q, want %q", got, want)
		}
	}
}
