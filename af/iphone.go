package af

type iphone struct {
}

func (a *iphone) makeChip() Chip {
	return &iphoneChip{
		shoe: shoe{
			logo: "iphone",
			size: 14,
		},
	}
}

func (a *iphone) makeCamera() Camera {
	return &adidasShort{
		short: short{
			logo: "iphone",
			size: 14,
		},
	}
}
