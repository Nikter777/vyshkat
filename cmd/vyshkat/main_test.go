package main

import (
	"bytes"
	"strings"
	"testing"
)

func TestRun(t *testing.T) {
	var output bytes.Buffer

	run(&output)

	result := output.String()

	expected := []string{
		"Принят: Ninebot Max",
		"Принят: City Bike",
		"Принят: E-Bike Pro",
		"Отклонён: Broken Scooter",
		"Транспорт: 3",
		"Вещи: 2",
		"Суммарное энергопотребление: 5.30",
		"Ninebot Max (простота: 8)",
		"City Bike (простота: 9)",
		"SC-001 — Ninebot Max",
		"BC-001 — City Bike",
		"EB-001 — E-Bike Pro",
		"HL-001 — Safety Helmet",
		"CH-001 — Charging Station",
	}

	for _, value := range expected {
		if !strings.Contains(result, value) {
			t.Errorf("expected output to contain %q", value)
		}
	}

	if strings.Contains(result, "Принят: Broken Scooter") {
		t.Error("rejected transport must not be accepted")
	}

	if strings.Contains(result, "Broken Scooter (простота:") {
		t.Error("rejected transport must not appear in beginner-friendly list")
	}
}
