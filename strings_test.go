package gotest_test

import (
	"github.com/tclearned/gotest"
	"testing"
)

func TestRemoveSubstring(t *testing.T) {
	got := gotest.RemoveSubstring("abaaba", "aa")

	if got != "abba" {
		t.Errorf("RemoveSubstring(\"abaaba\") = \"%s\"; want \"abba\"", got)
	}

	got = gotest.RemoveSubstring("haystackneedlehaystack", "needle")

	if got != "haystackhaystack" {
		t.Errorf("RemoveSubstring(\"haystackneedlehaystack\") = \"%s\"; want \"haystackhaystack\"")
	}
}
