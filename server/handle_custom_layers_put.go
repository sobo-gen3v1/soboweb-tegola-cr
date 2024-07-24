package server

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/dimfeld/httptreemux"
	"github.com/sobo-gen3v1/soboweb-tegola-cr/internal/log"
	"github.com/sobo-gen3v1/soboweb-tegola-cr/model"
)

type HandlePutCustomLayers struct {
	Layers  []model.DArea `json:"layers"`
	Version model.VArea   `json:"version"`
}

func (req *HandlePutCustomLayers) parseParams(r *http.Request) error {
	params := httptreemux.ContextParams(r.Context())

	log.Debugf("[CustomLayers] [parseParams] params: %+v", params)

	err := json.NewDecoder(r.Body).Decode(req)
	if err != nil {
		log.Errorf("[CustomLayers] [parseParams] error: %+v", err)
		return fmt.Errorf("decode json struct failed: %+v", req)
	}

	return nil
}

func (req HandlePutCustomLayers) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Debug("[CustomLayers] [ServeHTTP] start")
	// parse body
	if err := req.parseParams(r); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	db, err := model.GetPgDb()
	if err != nil {
		log.Error(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	result := db.Model(&model.DArea{}).Create(&req.Layers)
	if result.Error != nil {
		log.Error(result.Error)
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
	}

	result = db.Model(&model.VArea{}).Create(&req.Version)
	if result.Error != nil {
		log.Error(result.Error)
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
	}

	log.Debug("[CustomLayers] [ServeHTTP] end")
	json.NewEncoder(w).Encode(map[string]any{
		"layers":  &req.Layers,
		"version": &req.Version,
	})

}
