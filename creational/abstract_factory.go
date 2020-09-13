package creational

import (
	"fmt"
	"io"
)

type (
	Vehicle interface {
		Move()
	}

	VehicleFactory interface {
		Manufacture(w io.Writer) Vehicle
	}

	VehicleType int

	car struct {
		w io.Writer
	}

	aircraft struct {
		w io.Writer
	}

	carFactory      struct{}
	aircraftFactory struct{}
)

const (
	Car VehicleType = iota
	Aircraft
)

func (c car) Move() {
	_, _ = fmt.Fprint(c.w, "Move a protoCar")
}

func (a aircraft) Move() {
	_, _ = fmt.Fprint(a.w, "Move a aircraft")
}

func (cf carFactory) Manufacture(w io.Writer) Vehicle {
	return car{w}
}

func (pf aircraftFactory) Manufacture(w io.Writer) Vehicle {
	return aircraft{w}
}

func GetVehicleFactory(t VehicleType) (VehicleFactory, error) {
	switch t {
	default:
		return nil, fmt.Errorf("unknown vehicle type %d", t)
	case Car:
		return carFactory{}, nil
	case Aircraft:
		return aircraftFactory{}, nil
	}
}
