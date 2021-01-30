package flyweight

type counterTerroristDress struct {
    color string
}

func (c *counterTerroristDress) GetColor() string {
    return c.color
}

func NewCounterTerroristDress() *counterTerroristDress {
    return &counterTerroristDress{color: "green"}
}