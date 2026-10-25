package di

import "vyshkat/internal/service"

type Container struct {
	ServiceCenter service.ServiceCenter
	Fleet         *service.Fleet
}

func NewContainer() *Container {
	serviceCenter := service.DefaultServiceCenter{}

	return &Container{
		ServiceCenter: serviceCenter,
		Fleet:         service.NewFleet(serviceCenter),
	}
}
