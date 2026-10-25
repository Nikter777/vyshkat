package service

import (
	"errors"
	"testing"

	"vyshkat/internal/domain"
)

type FakeServiceCenter struct {
	Approved bool
	Calls    int
}

func (f *FakeServiceCenter) Inspect(transport domain.Transport) bool {
	f.Calls++
	return f.Approved
}

func TestAddTransportApproved(t *testing.T) {
	center := &FakeServiceCenter{Approved: true}
	fleet := NewFleet(center)

	scooter := domain.Scooter{
		ID: "SC-001", Model: "Ninebot", Energy: 1.5, Ease: 8,
	}

	err := fleet.AddTransport(scooter)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if fleet.TransportCount() != 1 {
		t.Errorf("expected 1 transport, got %d", fleet.TransportCount())
	}

	if center.Calls != 1 {
		t.Errorf("expected 1 inspection, got %d", center.Calls)
	}
}

func TestAddTransportRejected(t *testing.T) {
	center := &FakeServiceCenter{Approved: false}
	fleet := NewFleet(center)

	scooter := domain.Scooter{
		ID: "SC-002", Model: "Broken Scooter", Ease: 8,
	}

	err := fleet.AddTransport(scooter)

	if !errors.Is(err, ErrInspectionFailed) {
		t.Errorf("expected ErrInspectionFailed, got %v", err)
	}

	if fleet.TransportCount() != 0 {
		t.Errorf("rejected transport must not enter the fleet")
	}

	if center.Calls != 1 {
		t.Errorf("expected 1 inspection, got %d", center.Calls)
	}
}

func TestBeginnerFriendly(t *testing.T) {
	fleet := NewFleet(&FakeServiceCenter{Approved: true})

	transports := []domain.Transport{
		domain.Scooter{ID: "SC-001", Ease: 5},
		domain.Scooter{ID: "SC-002", Ease: 6},
		domain.Bicycle{ID: "BC-001", Ease: 9},
	}

	for _, transport := range transports {
		if err := fleet.AddTransport(transport); err != nil {
			t.Fatal(err)
		}
	}

	result := fleet.BeginnerFriendly()

	if len(result) != 2 {
		t.Fatalf("expected 2 beginner-friendly transports, got %d", len(result))
	}

	if result[0].InventoryNumber() != "SC-002" {
		t.Errorf("unexpected first transport: %s", result[0].InventoryNumber())
	}

	if result[1].InventoryNumber() != "BC-001" {
		t.Errorf("unexpected second transport: %s", result[1].InventoryNumber())
	}
}

func TestTotalEnergyConsumption(t *testing.T) {
	fleet := NewFleet(&FakeServiceCenter{Approved: true})

	transports := []domain.Transport{
		domain.Scooter{ID: "SC-001", Energy: 1.5},
		domain.Bicycle{ID: "BC-001"},
		domain.ElectricBicycle{ID: "EB-001", Energy: 0.8},
	}

	for _, transport := range transports {
		if err := fleet.AddTransport(transport); err != nil {
			t.Fatal(err)
		}
	}

	fleet.AddItem(domain.Helmet{ID: "HL-001"})
	fleet.AddItem(domain.ChargingStation{ID: "CH-001", Energy: 3.0})

	expected := 5.3
	actual := fleet.TotalEnergyConsumption()

	if actual < expected-0.000001 || actual > expected+0.000001 {
		t.Errorf("expected energy %.2f, got %.2f", expected, actual)
	}
}

func TestInventory(t *testing.T) {
	fleet := NewFleet(&FakeServiceCenter{Approved: true})

	if err := fleet.AddTransport(domain.Scooter{ID: "SC-001"}); err != nil {
		t.Fatal(err)
	}

	fleet.AddItem(domain.Helmet{ID: "HL-001"})

	if fleet.ItemCount() != 1 {
		t.Errorf("expected 1 item, got %d", fleet.ItemCount())
	}

	inventory := fleet.Inventory()

	if len(inventory) != 2 {
		t.Fatalf("expected 2 inventory objects, got %d", len(inventory))
	}

	if inventory[0].InventoryNumber() != "SC-001" {
		t.Errorf("unexpected transport inventory number")
	}

	if inventory[1].InventoryNumber() != "HL-001" {
		t.Errorf("unexpected item inventory number")
	}
}

func TestDefaultServiceCenter(t *testing.T) {
	center := DefaultServiceCenter{}

	tests := []struct {
		name     string
		scooter  domain.Scooter
		expected bool
	}{
		{
			name:     "valid transport",
			scooter:  domain.Scooter{ID: "SC-001", Model: "Ninebot", Ease: 6},
			expected: true,
		},
		{
			name:     "missing inventory number",
			scooter:  domain.Scooter{Model: "Ninebot", Ease: 6},
			expected: false,
		},
		{
			name:     "missing model",
			scooter:  domain.Scooter{ID: "SC-001", Ease: 6},
			expected: false,
		},
		{
			name:     "ease below minimum",
			scooter:  domain.Scooter{ID: "SC-001", Model: "Ninebot", Ease: 0},
			expected: false,
		},
		{
			name:     "ease above maximum",
			scooter:  domain.Scooter{ID: "SC-001", Model: "Ninebot", Ease: 11},
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := center.Inspect(tt.scooter)

			if actual != tt.expected {
				t.Errorf("expected %t, got %t", tt.expected, actual)
			}
		})
	}
}
