package service

import "vyshkat/internal/domain"

type ServiceCenter interface {
	Inspect(transport domain.Transport) bool
}

type DefaultServiceCenter struct{}

func (s DefaultServiceCenter) Inspect(transport domain.Transport) bool {
	return transport.InventoryNumber() != "" &&
		transport.Name() != "" &&
		transport.EaseOfUse() >= 1 &&
		transport.EaseOfUse() <= 10
}
