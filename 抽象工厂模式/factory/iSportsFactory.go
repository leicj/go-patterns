package factory

import "fmt"

type iSportsFactory interface {
	MakeShoe() IShoe
	MakeShirt() IShirt
}

func GetSportsFactory(brand string) (iSportsFactory, error) {
	if brand == "adidas" {
		return &adidas{}, nil
	}

	if brand == "nike" {
		return &nike{}, nil
	}

	return nil, fmt.Errorf("Wrong brand type passed")
}
