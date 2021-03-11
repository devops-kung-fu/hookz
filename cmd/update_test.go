package cmd

import (
	"testing"
)

func TestUpdateExecutables(t *testing.T) {
	err := updateExecutables()
	if err != nil {
		t.Fail()
	}
}