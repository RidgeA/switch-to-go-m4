package behavioral

import "testing"

func TestEvaluate(t *testing.T) {
	v := Evaluate("2 2 +")

	if v != 4 {
		t.Errorf("Expected 4, got %d", v)
	}
}
