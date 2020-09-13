package creational

type (
	Beverage struct {
		name  string
		price int
	}

	VendingMachine struct {
		protoBeverage Beverage
	}
)

func NewVendingMachine(b Beverage) VendingMachine {
	return VendingMachine{b}
}

func (vm VendingMachine) MakeBeverage() Beverage {
	return vm.protoBeverage.Clone()
}

func (b Beverage) Clone() Beverage {
	return Beverage{
		name:  b.name,
		price: b.price,
	}
}
