package creational

import (
	"strings"
	"testing"
)

func TestNewCheapHouse(t *testing.T) {
	h := NewCheapHouse()

	if !strings.Contains(h.String(), "cheap construction supplies") ||
		!strings.Contains(h.String(), "cheap furniture") {
		t.Errorf("A cheap house should be built using `NewCheapHouse`. Got: %s", h.String())
	}
}

func TestNewLuxuryHouse(t *testing.T) {
	h := NewLuxuryHouse()

	if !strings.Contains(h.String(), "luxury construction supplies") ||
		!strings.Contains(h.String(), "luxury furniture") {
		t.Errorf("A luxury house should be built using `NewLuxuryHouse`. Got: %s", h.String())
	}
}
