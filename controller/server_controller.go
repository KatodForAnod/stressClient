package controller

import (
	"errors"
	log "github.com/sirupsen/logrus"
	"stressClient/iot"
)

type Controller struct {
	iotDevices map[int]iot.IoT
}

func (c *Controller) Init() {
	c.iotDevices = make(map[int]iot.IoT)
}

func (c *Controller) AddIot(iotDevice iot.IoT) error {
	log.Infoln("Controller AddIot", iotDevice.GetId())

	if _, isExist := c.iotDevices[iotDevice.GetId()]; isExist {
		return errors.New("iot device already exist")
	}

	c.iotDevices[iotDevice.GetId()] = iotDevice
	return nil
}

func (c *Controller) RemoveIot(id int) error {
	log.Infoln("Controller RemoveIot", id)

	if _, isExist := c.iotDevices[id]; !isExist {
		return errors.New("iot device not exist")
	}

	delete(c.iotDevices, id)
	return nil
}

func (c *Controller) GetInformationIotDevFunc1(id int) (string, error) {
	log.Infoln("Controller getInformationIotDevFunc1")
	iotDevice, isExist := c.iotDevices[id]
	if !isExist {
		return "", errors.New("iot device not exist")
	}

	return iotDevice.GetCurrentStateFunc1()
}

func (c *Controller) GetInformationIotDevFunc2(id int) (string, error) {
	log.Infoln("Controller getInformationIotDevFunc2")
	iotDevice, isExist := c.iotDevices[id]
	if !isExist {
		return "", errors.New("iot device not exist")
	}

	return iotDevice.GetCurrentStateFunc2()
}

func (c *Controller) GetInformationIotDevFunc3(id int) (string, error) {
	log.Infoln("Controller getInformationIotDevFunc3")
	iotDevice, isExist := c.iotDevices[id]
	if !isExist {
		return "", errors.New("iot device not exist")
	}

	return iotDevice.GetCurrentStateFunc3()
}

func (c *Controller) GetInformationIotDevFunc4(id int) (string, error) {
	log.Infoln("Controller getInformationIotDevFunc4")
	iotDevice, isExist := c.iotDevices[id]
	if !isExist {
		return "", errors.New("iot device not exist")
	}

	return iotDevice.GetCurrentStateFunc4()
}

func (c *Controller) GetInformationIotDevFunc5(id int) (string, error) {
	log.Infoln("Controller getInformationIotDevFunc5")
	iotDevice, isExist := c.iotDevices[id]
	if !isExist {
		return "", errors.New("iot device not exist")
	}

	return iotDevice.GetCurrentStateFunc5()
}
