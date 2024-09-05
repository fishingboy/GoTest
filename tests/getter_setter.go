package tests

type Player struct {
	PlayerID int
	Name     string
	Age      int
	money    int
}

func (this Player) GetName() string {
	return this.Name
}

// *Player 可以視為傳址呼叫，改的值才會有用
func (this *Player) SetName(name string) error {
	this.Name = name
	return nil
}

func (this Player) GetMoney() int {
	return this.money
}

func (this *Player) SetMoney(money int) error {
	this.money = money
	return nil
}
