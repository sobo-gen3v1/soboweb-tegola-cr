package model

import (
	"encoding/json"
	"errors"
	"testing"

	"gorm.io/gorm"
)

func TestModel(t *testing.T) {
	db, err := GetPgDb()
	if err != nil {
		t.Logf("Err: %+v", err)
		t.FailNow()
	}
	dbInstance, _ := db.DB()
	defer dbInstance.Close()

	var layers []MArea

	result := db.Find(&layers)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		t.Logf("Err: %+v", result.Error)
		t.FailNow()
	}

	dataJson, _ := json.MarshalIndent(layers, "", "    ")
	t.Logf("{%+v} \n", string(dataJson))
}
