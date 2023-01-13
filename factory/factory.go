package factory

import (
	interfaces "abstract_factory/interfaces"
	nike "abstract_factory/nike"
	adidas "abstract_factory/adidas"
	"fmt"
)

func GetSportsFactory(brand string) (interfaces.ISportsFactory, error) {
	if brand == "adidas" {
		return &adidas.Adidas{}, nil
	}
	if brand == "nike" {
		return &nike.Nike{}, nil
	}
	return nil, fmt.Errorf("wrong type of brand passed")
}