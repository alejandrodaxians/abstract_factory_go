package adidas

import interfaces "abstract_factory/interfaces"

type Adidas struct {}

func (a *Adidas) MakeShoe() interfaces.IShoe {
	return &adidasShoe {
		Shoe: interfaces.Shoe {
			Logo: "adidas",
			Size: 14,
		},
	}
}

func (a *Adidas) MakeShort() interfaces.IShort {
	return &adidasShort {
		Short: interfaces.Short {
			Logo: "adidas",
			Size: 14,
		},
	}
}