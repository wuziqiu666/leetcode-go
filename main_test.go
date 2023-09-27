package main

import (
	"testing"
)

func TestDeleteDuplicatess(t *testing.T) {
	cases := [][]int{
		{1, 2, 3, 3, 4, 4, 5},
		{1, 1},
	}
	for _, c := range cases {
		head := NewListNode(c)
		head = deleteDuplicates(head)
		got := head.String()
		want := "[1, 2, 5]"
		if got != want {
			t.Fatalf("got %q, want %q", got, want)
		}
	}
}
