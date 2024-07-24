package model

import (
	"time"

	"gorm.io/gorm"
)

type VArea struct {
	Version   string         `gorm:"primarykey;column:version" json:"version"`
	LayerName string         `gorm:"null;column:layer_name" json:"layer_name,omitempty"`
	StartTime time.Time      `gorm:"column:start_time" json:"start_time"`
	EndTime   time.Time      `gorm:"column:end_time" json:"end_time"`
	CreatedAt time.Time      `gorm:"column:created_at" json:"created_at,omitempty"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime:milli;column:updated_at" json:"updated_at,omitempty"`
	DeletedAt gorm.DeletedAt `gorm:"null;index;column:deleted_at" json:"deleted_at,omitempty"`
}
