package utils

import (
	"testing"
)

func Test_URL_ShiftPath(t *testing.T) {
	paths := []string{
		"", "", "/",
		"/", "", "/",
		"/scoutfish", "scoutfish", "/",
		"/scoutfish/", "scoutfish", "/",
		"/scoutfish/1", "scoutfish", "/1",
	}

	for i := 0; i < len(paths); i += 3 {
		head, tail := ShiftPath(paths[i])
		if head != paths[i+1] {
			t.Error("path", paths[i], "expected head", paths[i+1], "received", head)
		}
		if tail != paths[i+2] {
			t.Error("path", paths[i], "expected tail", paths[i+2], "received", tail)
		}
	}
}
