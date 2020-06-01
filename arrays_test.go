package gotest_test

import (
	"github.com/tclearned/gotest"
	"reflect"
	"testing"
)

func TestSliceFromArray(t *testing.T) {
	got := gotest.SliceFromArray([1]int{2})

	if !reflect.DeepEqual([]int{2}, got) {
		t.Errorf("SliceFromArray([2]) = %v; want [2]", got)
	}

	got = gotest.SliceFromArray([1]int{3})

	if !reflect.DeepEqual([]int{3}, got) {
		t.Errorf("SliceFromArray([3]) = %v; want [3]", got)
	}
}

func TestReverseSlice(t *testing.T) {
	got := gotest.ReverseSlice([]int{1, 2, 3})

	if !reflect.DeepEqual([]int{3, 2, 1}, got) {
		t.Errorf("ReverseSlice([1, 2, 3]) = %v; want [3, 2, 1]", got)
	}

	got = gotest.ReverseSlice([]int{4, 5, 6})

	if !reflect.DeepEqual([]int{6, 5, 4}, got) {
		t.Errorf("ReverseSlice([4, 5, 6]) = %v; want [6, 5, 4]", got)
	}
}
