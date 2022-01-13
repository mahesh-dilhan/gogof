package main

type regularBuilder struct {
	windowType string
	doorType   string
	floor      int
}

func newRegularBuilder() *regularBuilder {
	return &regularBuilder{}
}

func (b *regularBuilder) setWindowType() {
	b.windowType = "Wooden Window"
}

func (b *regularBuilder) setDoorType() {
	b.doorType = "Wooden Door"
}
func (b *regularBuilder) setNumFloor() {
	b.floor = 2
}
func (b *regularBuilder) getHouse() house {
	return house{
		doorType:   b.doorType,
		windowType: b.windowType,
		floor:      b.floor,
	}
}
