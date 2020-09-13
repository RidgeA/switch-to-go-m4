package creational

import "fmt"

type (
	House struct {
		furniture            string
		constructionSupplies string
	}
)

func (h House) String() string {
	return fmt.Sprintf("This house built using %s and it has %s", h.constructionSupplies, h.furniture)
}

func NewCheapHouse() House {
	return House{
		furniture:            "cheap furniture",
		constructionSupplies: "cheap construction supplies",
	}
}

func NewLuxuryHouse() House {
	return House{
		furniture:            "luxury furniture",
		constructionSupplies: "luxury construction supplies",
	}
}
