package service

import (
	"errors"

	"vyshkat/internal/domain"
)

var ErrInspectionFailed = errors.New("transport failed inspection")

type Fleet struct {
	serviceCenter ServiceCenter
	transports    []domain.Transport
	items         []domain.InventoryItem
}

func NewFleet(serviceCenter ServiceCenter) *Fleet {
	return &Fleet{
		serviceCenter: serviceCenter,
		transports:    make([]domain.Transport, 0),
		items:         make([]domain.InventoryItem, 0),
	}
}

func (f *Fleet) AddTransport(transport domain.Transport) error {
	if !f.serviceCenter.Inspect(transport) {
		return ErrInspectionFailed
	}

	f.transports = append(f.transports, transport)
	return nil
}

func (f *Fleet) AddItem(item domain.InventoryItem) {
	f.items = append(f.items, item)
}

func (f *Fleet) TransportCount() int {
	return len(f.transports)
}

func (f *Fleet) ItemCount() int {
	return len(f.items)
}

func (f *Fleet) TotalEnergyConsumption() float64 {
	var total float64

	for _, transport := range f.transports {
		if consumer, ok := transport.(domain.EnergyConsumer); ok {
			total += consumer.DailyEnergyConsumption()
		}
	}

	for _, item := range f.items {
		if consumer, ok := item.(domain.EnergyConsumer); ok {
			total += consumer.DailyEnergyConsumption()
		}
	}

	return total
}

func (f *Fleet) BeginnerFriendly() []domain.Transport {
	result := make([]domain.Transport, 0)

	for _, transport := range f.transports {
		if transport.EaseOfUse() >= 6 {
			result = append(result, transport)
		}
	}

	return result
}

func (f *Fleet) Inventory() []domain.InventoryItem {
	result := make([]domain.InventoryItem, 0, len(f.transports)+len(f.items))

	for _, transport := range f.transports {
		result = append(result, transport)
	}

	result = append(result, f.items...)

	return result
}
