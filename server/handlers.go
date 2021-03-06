package server

import (
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

func (s *Server) getInformationFromIotDeviceFunc1(w http.ResponseWriter, r *http.Request) {
	log.Println("handler getInformationFromIotDevice")
	defer r.Body.Close()

	vars := mux.Vars(r)
	iotId, isExist := vars["iot"]
	if !isExist {
		http.Error(w, "set iot id", http.StatusBadRequest)
		return
	}

	atoiId, err := strconv.Atoi(iotId)
	if err != nil {
		log.Errorln(err)
		http.Error(w, "iot id must be int", http.StatusBadRequest)
		return
	}

	inf, err := s.controller.GetInformationIotDevFunc1(atoiId)
	if err != nil {
		log.Errorln(err)
		http.Error(w, "internal error", http.StatusInternalServerError)
		return
	}

	_, err = w.Write([]byte(inf))
	if err != nil {
		log.Errorln(err)
		http.Error(w, "internal error", http.StatusInternalServerError)
		return
	}
}

func (s *Server) getInformationFromIotDeviceFunc2(w http.ResponseWriter, r *http.Request) {
	log.Println("handler getInformationFromIotDevice")
	defer r.Body.Close()

	vars := mux.Vars(r)
	iotId, isExist := vars["iot"]
	if !isExist {
		http.Error(w, "set iot id", http.StatusBadRequest)
		return
	}

	atoiId, err := strconv.Atoi(iotId)
	if err != nil {
		log.Errorln(err)
		http.Error(w, "iot id must be int", http.StatusBadRequest)
		return
	}

	inf, err := s.controller.GetInformationIotDevFunc2(atoiId)
	if err != nil {
		log.Errorln(err)
		http.Error(w, "internal error", http.StatusInternalServerError)
		return
	}

	_, err = w.Write([]byte(inf))
	if err != nil {
		log.Errorln(err)
		http.Error(w, "internal error", http.StatusInternalServerError)
		return
	}
}

func (s *Server) getInformationFromIotDeviceFunc3(w http.ResponseWriter, r *http.Request) {
	log.Println("handler getInformationFromIotDevice")
	defer r.Body.Close()

	vars := mux.Vars(r)
	iotId, isExist := vars["iot"]
	if !isExist {
		http.Error(w, "set iot id", http.StatusBadRequest)
		return
	}

	atoiId, err := strconv.Atoi(iotId)
	if err != nil {
		log.Errorln(err)
		http.Error(w, "iot id must be int", http.StatusBadRequest)
		return
	}

	inf, err := s.controller.GetInformationIotDevFunc3(atoiId)
	if err != nil {
		log.Errorln(err)
		http.Error(w, "internal error", http.StatusInternalServerError)
		return
	}

	_, err = w.Write([]byte(inf))
	if err != nil {
		log.Errorln(err)
		http.Error(w, "internal error", http.StatusInternalServerError)
		return
	}
}

func (s *Server) getInformationFromIotDeviceFunc4(w http.ResponseWriter, r *http.Request) {
	log.Println("handler getInformationFromIotDevice")
	defer r.Body.Close()

	vars := mux.Vars(r)
	iotId, isExist := vars["iot"]
	if !isExist {
		http.Error(w, "set iot id", http.StatusBadRequest)
		return
	}

	atoiId, err := strconv.Atoi(iotId)
	if err != nil {
		log.Errorln(err)
		http.Error(w, "iot id must be int", http.StatusBadRequest)
		return
	}

	inf, err := s.controller.GetInformationIotDevFunc4(atoiId)
	if err != nil {
		log.Errorln(err)
		http.Error(w, "internal error", http.StatusInternalServerError)
		return
	}

	_, err = w.Write([]byte(inf))
	if err != nil {
		log.Errorln(err)
		http.Error(w, "internal error", http.StatusInternalServerError)
		return
	}
}

func (s *Server) getInformationFromIotDeviceFunc5(w http.ResponseWriter, r *http.Request) {
	log.Println("handler getInformationFromIotDevice")
	defer r.Body.Close()

	vars := mux.Vars(r)
	iotId, isExist := vars["iot"]
	if !isExist {
		http.Error(w, "set iot id", http.StatusBadRequest)
		return
	}

	atoiId, err := strconv.Atoi(iotId)
	if err != nil {
		log.Errorln(err)
		http.Error(w, "iot id must be int", http.StatusBadRequest)
		return
	}

	inf, err := s.controller.GetInformationIotDevFunc5(atoiId)
	if err != nil {
		log.Errorln(err)
		http.Error(w, "internal error", http.StatusInternalServerError)
		return
	}

	_, err = w.Write([]byte(inf))
	if err != nil {
		log.Errorln(err)
		http.Error(w, "internal error", http.StatusInternalServerError)
		return
	}
}
