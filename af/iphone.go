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

func (a *iphone) makeCamera() Camera {
	return &iPhoneCamera{
		camera: camera{
			pixel: "10",
		},
	}
}
