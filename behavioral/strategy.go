package behavioral

import "math"

type (
	DiscountStrategy interface {
		Apply(float64) float64
	}

	VIPDiscount struct {
	}

	PromoDiscount struct{}

	NoDiscount struct{}
)

func (p PromoDiscount) Apply(price float64) float64 {
	return price - math.Min(price*0.3, 100.)
}

func (V VIPDiscount) Apply(price float64) float64 {
	return price * 0.7
}

func CalculatePrice(itemPrice float64, quantity int, discount DiscountStrategy) float64 {
	return discount.Apply(itemPrice * float64(quantity))
}
