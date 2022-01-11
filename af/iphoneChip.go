package main

type iphoneChip struct {
	chip
}

func (i *iphoneChip) powerby(manfacture string) string {
	return manfacture
}
