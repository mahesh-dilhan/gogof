package af

type iphone struct {
}

func (a *iphone) makeChip() Chip {
	return &iphoneChip{
		chip: chip{
			powerby: "m1",
		},
	}
}
