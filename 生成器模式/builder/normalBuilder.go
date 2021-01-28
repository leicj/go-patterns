package builder

type normalBuilder struct {
	WindowType string
	DoorType string
	Floor int
}

func newnormalBuilder() *normalBuilder {
	return &normalBuilder{}
}

func (b *normalBuilder) setWindowType() {
	b.WindowType = "Wooden Window"
}

func (b *normalBuilder) setDoorType() {
	b.DoorType = "Wooden Door"
}

func (b *normalBuilder) setNumFloor() {
	b.Floor =2
}

func (b *normalBuilder) getHouse() house {
	return house{
		DoorType: b.DoorType,
		WindowType: b.WindowType,
		Floor: b.Floor,
	}
}