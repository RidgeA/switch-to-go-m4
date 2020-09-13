package creational

import "testing"

func TestCounter_Inc(t *testing.T) {
	Counter.Inc()

	if Counter.Value() != 1 {
		t.Errorf("Counter should return 1 after one Inc() call, got %d instead", Counter.Value())
	}
}
