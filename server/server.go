package server

import (
	"fmt"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"net/http"
	"stressClient/config"
	"stressClient/controller"
)

type Server struct {
	controller controller.ServerController
}

func (s *Server) StartServer(config config.Config, controller controller.ServerController) {
	s.controller = controller

	r := mux.NewRouter()
	r.HandleFunc("/device/{iot:[0-9]+}/metrics/func1", s.getInformationFromIotDeviceFunc1)
	r.HandleFunc("/device/{iot:[0-9]+}/metrics/func2", s.getInformationFromIotDeviceFunc2)
	r.HandleFunc("/device/{iot:[0-9]+}/metrics/func3", s.getInformationFromIotDeviceFunc3)
	r.HandleFunc("/device/{iot:[0-9]+}/metrics/func4", s.getInformationFromIotDeviceFunc4)
	r.HandleFunc("/device/{iot:[0-9]+}/metrics/func5", s.getInformationFromIotDeviceFunc5)

	fmt.Println("Server is listening... ", config.ProxyServerAddr)
	log.Fatal(http.ListenAndServe(config.ProxyServerAddr, r))
}
