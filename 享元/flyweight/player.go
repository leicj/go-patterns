package flyweight

type player struct {
    dress      dress
    playerType string
    lat        int
    long       int
}

func NewPlayer(playerType, dressType string) *player {
    dress, _ := GetDressFactorySingleInstance().GetDressByType(dressType)
    return &player{
        playerType: playerType,
        dress:      dress,
    }
}

func (p *player) NewLocation(lat, long int) {
    p.lat = lat
    p.long = long
}