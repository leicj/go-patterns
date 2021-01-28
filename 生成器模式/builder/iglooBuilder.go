package builder

type iglooBuilder struct {
	WindowType string
	DoorType string
	Floor int
}

func newIglooBuilder() *iglooBuilder {
	return &iglooBuilder{}
}

func (b *iglooBuilder) setWindowType() {
    b.WindowType = "Snow Window"
}

func (b *iglooBuilder) setDoorType() {
    b.DoorType = "Snow Door"
}

func (b *iglooBuilder) setNumFloor() {
    b.Floor = 1
}

func (b *iglooBuilder) getHouse() house {
    return house{
        DoorType:   b.DoorType,
        WindowType: b.WindowType,
        Floor:      b.Floor,
    }
}