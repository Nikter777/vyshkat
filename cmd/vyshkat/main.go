package main

import (
	"fmt"

	"vyshkat/internal/di"
	"vyshkat/internal/domain"
)

func main() {
	container := di.NewContainer()
	fleet := container.Fleet

	transports := []domain.Transport{
		domain.Scooter{
			ID: "SC-001", Model: "Ninebot Max", Energy: 1.5, Ease: 8,
		},
		domain.Bicycle{
			ID: "BC-001", Model: "City Bike", Ease: 9,
		},
		domain.ElectricBicycle{
			ID: "EB-001", Model: "E-Bike Pro", Energy: 0.8, Ease: 5,
		},
		domain.Scooter{
			ID: "", Model: "Broken Scooter", Energy: 2.0, Ease: 7,
		},
	}

	fmt.Println("=== Приём транспорта ===")

	for _, transport := range transports {
		err := fleet.AddTransport(transport)
		if err != nil {
			fmt.Printf("Отклонён: %s (%v)\n", transport.Name(), err)
			continue
		}

		fmt.Printf("Принят: %s\n", transport.Name())
	}

	fleet.AddItem(domain.Helmet{
		ID: "HL-001", Model: "Safety Helmet",
	})

	fleet.AddItem(domain.ChargingStation{
		ID: "CH-001", Model: "Charging Station", Energy: 3.0,
	})

	fmt.Println("\n=== Отчёт по парку ===")
	fmt.Printf("Транспорт: %d\n", fleet.TransportCount())
	fmt.Printf("Вещи: %d\n", fleet.ItemCount())
	fmt.Printf("Суммарное энергопотребление: %.2f кВт·ч/сутки\n", fleet.TotalEnergyConsumption())

	fmt.Println("\n=== Для новичков ===")

	for _, transport := range fleet.BeginnerFriendly() {
		fmt.Printf("%s (простота: %d)\n", transport.Name(), transport.EaseOfUse())
	}

	fmt.Println("\n=== Инвентаризация ===")

	for _, item := range fleet.Inventory() {
		fmt.Printf("%s — %s\n", item.InventoryNumber(), item.Name())
	}
}
