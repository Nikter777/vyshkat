package domain

type Bicycle struct {
	ID    string
	Model string
	Ease  int
}

func (b Bicycle) InventoryNumber() string {
	return b.ID
}

func (b Bicycle) Name() string {
	return b.Model
}

func (b Bicycle) EaseOfUse() int {
	return b.Ease
}

type ElectricBicycle struct {
	ID     string
	Model  string
	Energy float64
	Ease   int
}

func (b ElectricBicycle) InventoryNumber() string {
	return b.ID
}

func (b ElectricBicycle) Name() string {
	return b.Model
}

func (b ElectricBicycle) DailyEnergyConsumption() float64 {
	return b.Energy
}

func (b ElectricBicycle) EaseOfUse() int {
	return b.Ease
}
