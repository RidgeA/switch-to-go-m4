package creational

import "testing"

func TestNewVendingMachine(t *testing.T) {
	b := Beverage{
		name:  "cola",
		price: 42,
	}

	vm := NewVendingMachine(b)

	newBeverage := vm.MakeBeverage()

	if b.name != newBeverage.name {
		t.Error("both beverages should have the same name")
	}

	if b.price != newBeverage.price {
		t.Error("both beverages should have the same price")
	}

	if &b == &newBeverage {
		t.Error("it should be two different structures of the same type")
	}
}
