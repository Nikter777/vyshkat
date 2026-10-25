package main

import (
	"fmt"
	"io"
	"os"

	"vyshkat/internal/di"
	"vyshkat/internal/domain"
)

func main() {
	run(os.Stdout)
}

func run(out io.Writer) {
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

	fmt.Fprintln(out, "=== Приём транспорта ===")

	for _, transport := range transports {
		err := fleet.AddTransport(transport)
		if err != nil {
			fmt.Fprintf(out, "Отклонён: %s (%v)\n", transport.Name(), err)
			continue
		}

		fmt.Fprintf(out, "Принят: %s\n", transport.Name())
	}

	fleet.AddItem(domain.Helmet{
		ID: "HL-001", Model: "Safety Helmet",
	})

	fleet.AddItem(domain.ChargingStation{
		ID: "CH-001", Model: "Charging Station", Energy: 3.0,
	})

	fmt.Fprintln(out, "\n=== Отчёт по парку ===")
	fmt.Fprintf(out, "Транспорт: %d\n", fleet.TransportCount())
	fmt.Fprintf(out, "Вещи: %d\n", fleet.ItemCount())
	fmt.Fprintf(out, "Суммарное энергопотребление: %.2f кВт·ч/сутки\n", fleet.TotalEnergyConsumption())

	fmt.Fprintln(out, "\n=== Для новичков ===")

	for _, transport := range fleet.BeginnerFriendly() {
		fmt.Fprintf(out, "%s (простота: %d)\n", transport.Name(), transport.EaseOfUse())
	}

	fmt.Fprintln(out, "\n=== Инвентаризация ===")

	for _, item := range fleet.Inventory() {
		fmt.Fprintf(out, "%s — %s\n", item.InventoryNumber(), item.Name())
	}
}
