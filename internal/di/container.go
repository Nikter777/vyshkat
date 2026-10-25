package di

import "vyshkat/internal/service"

type Container struct {
	ServiceCenter service.ServiceCenter
	Fleet         *service.Fleet
}

func NewContainer() *Container {
	return NewContainerWithServiceCenter(service.DefaultServiceCenter{})
}

func NewContainerWithServiceCenter(center service.ServiceCenter) *Container {
	return &Container{
		ServiceCenter: center,
		Fleet:         service.NewFleet(center),
	}
}
