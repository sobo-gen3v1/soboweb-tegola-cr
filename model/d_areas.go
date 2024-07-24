package model

import (
	"time"

	"gorm.io/gorm"
)

type DArea struct {
	ID            uint64         `gorm:"primarykey;autoIncrement;column:ogc_fid" json:"id"`
	Geometry      string         `gorm:"column:wkb_geometry;type:geometry" json:"geometry"`
	GeometricType GeometricType  `gorm:"column:geometric_type" json:"geometric_type"`
	AreaCode      string         `gorm:"column:area_code" json:"area_code"`
	GisDataId     string         `gorm:"column:gis_data_id" json:"gis_data_id"`
	Properties    string         `gorm:"null;column:properties" json:"properties,omitempty"`
	Tags          string         `gorm:"null;column:tags" json:"tags,omitempty"`
	CreatedAt     time.Time      `gorm:"column:created_at" json:"created_at,omitempty"`
	UpdatedAt     time.Time      `gorm:"autoUpdateTime:milli;column:updated_at" json:"updated_at,omitempty"`
	DeletedAt     gorm.DeletedAt `gorm:"null;index;column:deleted_at" json:"deleted_at,omitempty"`
}

func (DArea) TableName() string {
	return "public.d_areas"
}
