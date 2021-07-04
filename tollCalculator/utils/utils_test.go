package utils

import (
	"testing"
)

func TestIsfeeFreeDay(t *testing.T) {

	free := IsfeeFreeDay("May", 16)

	if free == false {
		t.Fatalf("Result was incorrect, got: %d, want: %d.", free, true)
	}

	free := IsfeeFreeDay("Dec", 31)

	if free == false {
		t.Fatalf("Result was incorrect, got: %d, want: %d.", free, true)
	}

	free := IsfeeFreeDay("Dec", 30)

	if free == true {
		t.Fatalf("Result was incorrect, got: %d, want: %d.", free, false)
	}
}
