package decorator

type VeggeMania struct {
}

func (p *VeggeMania) GetPrice() int {
	return 15
}