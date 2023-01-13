package nike

import interfaces "abstract_factory/interfaces"

type Nike struct{}

func (n *Nike) MakeShoe() interfaces.IShoe {
	return &nikeShoe{
		Shoe: interfaces.Shoe{
			Logo: "nike",
			Size: 14,
		},
	}
}

func (n *Nike) MakeShort() interfaces.IShort {
	return &nikeShort{
		Short: interfaces.Short{
			Logo: "nike",
			Size: 14,
		},
	}
}
