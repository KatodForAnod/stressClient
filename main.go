package main

import (
	log "github.com/sirupsen/logrus"
	"stressClient/config"
	"stressClient/controller"
	"stressClient/iot"
	"stressClient/server"
)

func main() {
	loadConfig, err := config.LoadConfig()
	if err != nil {
		log.Error(err)
		return
	}

	servcontroller := controller.Controller{}
	servcontroller.Init()

	for _, device := range loadConfig.IoTsDevices {
		iotCommon := iot.CommonIotDevice{}
		iotCommon.Init(device)

		err = servcontroller.AddIot(&iotCommon)
		if err != nil {
			log.Fatal(err)
			return
		}
	}

	server := server.Server{}
	server.StartServer(loadConfig, &servcontroller)
}
