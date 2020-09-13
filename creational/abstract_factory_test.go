package creational

import (
	"math"
	"strings"
	"testing"
)

func TestGetVehicleFactory_CarFactory(t *testing.T) {
	f, _ := GetVehicleFactory(Car)

	sb := &strings.Builder{}
	v := f.Manufacture(sb)
	v.Move()
	if sb.String() != "Move a protoCar" {
		t.Errorf("expected to get \"Move a protoCar\", got: %s", sb.String())
	}
}

func TestGetVehicleFactory_AircraftFactory(t *testing.T) {
	f, _ := GetVehicleFactory(Aircraft)

	sb := &strings.Builder{}
	v := f.Manufacture(sb)
	v.Move()
	if sb.String() != "Move a aircraft" {
		t.Errorf("expected to get \"Move a aircraft\", got: %s", sb.String())
	}
}

func TestGetVehicleFactory_Unknown(t *testing.T) {
	_, err := GetVehicleFactory(math.MaxInt32)

	if err == nil {
		t.Errorf("Expected error")
	}

}
