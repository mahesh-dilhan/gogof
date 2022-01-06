package main

type iphone struct {
}

func (a *iphone) makeChip() Chip {
	return &iphoneChip{
		chip: chip{
			manfacture: "m1",
		},
	}
}
