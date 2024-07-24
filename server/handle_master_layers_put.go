package server

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/dimfeld/httptreemux"
	"github.com/sobo-gen3v1/soboweb-tegola-cr/internal/log"
	"github.com/sobo-gen3v1/soboweb-tegola-cr/model"
	"gorm.io/gorm"
)

type ReqLayerInfo struct {
	AreaCode   string `json:"area_code"`
	Properties string `json:"properties"`
	Tags       string `json:"tags"`
}

type ReqLayer struct {
}

type HandlePutMasterLayers struct {
	Layers  []ReqLayerInfo `json:"layers"`
	Version model.VArea    `json:"version"`
}

func (req *HandlePutMasterLayers) parseParams(r *http.Request) error {
	params := httptreemux.ContextParams(r.Context())

	log.Debugf("[MasterLayers] [parseParams] params: %+v", params)

	err := json.NewDecoder(r.Body).Decode(req)
	if err != nil {
		log.Errorf("[MasterLayers] [parseParams] error: %+v", err)
		return fmt.Errorf("decode json struct failed: %+v", req)
	}

	return nil
}

func (req HandlePutMasterLayers) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Debug("[MasterLayers] [ServeHTTP] start")
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

	var areaCodes []string
	for _, item := range req.Layers {
		areaCodes = append(areaCodes, item.AreaCode)
	}

	var result *gorm.DB
	var mLayers []model.MArea

	result = db.Model(&model.MArea{}).Find(&mLayers, "area_code in ?", areaCodes)

	if result.Error != nil {
		log.Error(result.Error)
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
	}

	var dLayers []model.CArea
	for _, mArea := range mLayers {
		for _, reqInfo := range req.Layers {
			if reqInfo.AreaCode == mArea.AreaCode {
				dLayers = append(dLayers, model.CArea{
					OgcFid:        mArea.ID,
					ExtProperties: reqInfo.Properties,
					ExtTags:       reqInfo.Tags,
					CreatedAt:     req.Version.EndTime,
				})
			}
		}
	}

	result = db.Model(&model.DArea{}).Create(&dLayers)
	if result.Error != nil {
		log.Error(result.Error)
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
	}

	result = db.Model(&model.VArea{}).Create(&req.Version)
	if result.Error != nil {
		log.Error(result.Error)
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
	}

	log.Debug("[MasterLayers] [ServeHTTP] end")
	json.NewEncoder(w).Encode(map[string]any{
		"layers":  &dLayers,
		"version": &req.Version,
	})

}
