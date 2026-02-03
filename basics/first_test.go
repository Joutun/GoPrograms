package startingGoJourney

import "testing"

// Test change value works
func TestChangeValue(t *testing.T) {
	var knowledge Knowledge
	knowledge = Knowledge{value: 12, identifier: "banana"}

	if knowledge.value != 12 {
		t.Error("Value wasn't 12!")
	}

	knowledge.changeValue(100)

	if knowledge.value != 100 {
		t.Error("Value wasn't 100!")
	}
}
