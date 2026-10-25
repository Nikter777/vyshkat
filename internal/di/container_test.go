package di

import (
	"testing"

	"vyshkat/internal/domain"
)

func TestNewContainer(t *testing.T) {
	container := NewContainer()

	if container == nil {
		t.Fatal("container must not be nil")
	}

	if container.ServiceCenter == nil {
		t.Fatal("service center must not be nil")
	}

	if container.Fleet == nil {
		t.Fatal("fleet must not be nil")
	}

	valid := domain.Scooter{
		ID:    "SC-001",
		Model: "Ninebot",
		Ease:  8,
	}

	if err := container.Fleet.AddTransport(valid); err != nil {
		t.Fatalf("valid transport should be accepted: %v", err)
	}

	invalid := domain.Scooter{
		ID:    "",
		Model: "Broken Scooter",
		Ease:  8,
	}

	if err := container.Fleet.AddTransport(invalid); err == nil {
		t.Fatal("invalid transport should be rejected")
	}

	if container.Fleet.TransportCount() != 1 {
		t.Fatalf(
			"expected 1 accepted transport, got %d",
			container.Fleet.TransportCount(),
		)
	}
}

type FakeServiceCenter struct {
	Approved bool
}

func (f FakeServiceCenter) Inspect(transport domain.Transport) bool {
	return f.Approved
}

func TestContainerWithFakeServiceCenter(t *testing.T) {
	container := NewContainerWithServiceCenter(FakeServiceCenter{
		Approved: false,
	})

	scooter := domain.Scooter{
		ID:    "SC-001",
		Model: "Ninebot",
		Ease:  8,
	}

	err := container.Fleet.AddTransport(scooter)

	if err == nil {
		t.Fatal("expected transport to be rejected by fake service center")
	}

	if container.Fleet.TransportCount() != 0 {
		t.Fatal("rejected transport must not enter the fleet")
	}
}
