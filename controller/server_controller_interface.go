package controller

import "stressClient/iot"

type ServerController interface {
	AddIot(iotDevice iot.IoT) error
	RemoveIot(id int) error
	GetInformationIotDevFunc1(id int) (string, error)
	GetInformationIotDevFunc2(id int) (string, error)
	GetInformationIotDevFunc3(id int) (string, error)
	GetInformationIotDevFunc4(id int) (string, error)
	GetInformationIotDevFunc5(id int) (string, error)
}
