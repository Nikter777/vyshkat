package domain

type InventoryItem interface {
	InventoryNumber() string
	Name() string
}

type EnergyConsumer interface {
	DailyEnergyConsumption() float64
}

type Transport interface {
	InventoryItem
	EaseOfUse() int
}

type Scooter struct {
	ID     string
	Model  string
	Energy float64
	Ease   int
}

func (s Scooter) InventoryNumber() string {
	return s.ID
}

func (s Scooter) Name() string {
	return s.Model
}

func (s Scooter) DailyEnergyConsumption() float64 {
	return s.Energy
}

func (s Scooter) EaseOfUse() int {
	return s.Ease
}
