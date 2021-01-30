package flyweight

type terroristDress struct {
    color string
}

func (t *terroristDress) GetColor() string {
    return t.color
}

func NewTerroristDress() *terroristDress {
    return &terroristDress{color: "red"}
}