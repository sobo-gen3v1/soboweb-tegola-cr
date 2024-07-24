package model

import (
	"time"

	"gorm.io/gorm"
)

type CArea struct {
	ID            uint64         `gorm:"primaryKey;autoIncrement;column:geom_id" json:"geom_id,omitempty"`
	OgcFid        uint64         `gorm:"column:ogc_fid" json:"ogc_fid"`
	ExtProperties string         `gorm:"null;column:ext_properties" json:"ext_properties,omitempty"`
	ExtTags       string         `gorm:"null;column:ext_tags" json:"ext_tags,omitempty"`
	CreatedAt     time.Time      `gorm:"column:created_at" json:"created_at,omitempty"`
	UpdatedAt     time.Time      `gorm:"autoUpdateTime:milli;column:updated_at" json:"updated_at"`
	DeletedAt     gorm.DeletedAt `gorm:"null;index;column:deleted_at" json:"deleted_at,omitempty"`
}

func (CArea) TableName() string {
	return "public.c_areas"
}
