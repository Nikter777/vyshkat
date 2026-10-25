package domain

type Helmet struct {
	ID    string
	Model string
}

func (h Helmet) InventoryNumber() string {
	return h.ID
}

func (h Helmet) Name() string {
	return h.Model
}

type ChargingStation struct {
	ID     string
	Model  string
	Energy float64
}

func (c ChargingStation) InventoryNumber() string {
	return c.ID
}

func (c ChargingStation) Name() string {
	return c.Model
}

func (c ChargingStation) DailyEnergyConsumption() float64 {
	return c.Energy
}
