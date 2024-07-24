package model

import (
	"time"

	"gorm.io/gorm"
)

type GeometricType string

const (
	Point              GeometricType = "Point"
	Polygon            GeometricType = "Polygon"
	PolygonWithHole    GeometricType = "PolygonWithHole"
	Collection         GeometricType = "Collection"
	Linestring         GeometricType = "Linestring"
	MultiPoint         GeometricType = "MultiPoint"
	MultiLineString    GeometricType = "MultiLineString"
	MultiPolygon       GeometricType = "MultiPolygon"
	GeometryCollection GeometricType = "GeometryCollection"
)

type MArea struct {
	ID            uint64         `gorm:"primarykey;autoIncrement;column:ogc_fid" json:"id"`
	Geometry      string         `gorm:"column:wkb_geometry;type:geometry" json:"geometry"`
	GeometricType GeometricType  `gorm:"column:geometric_type" json:"geometric_type"`
	AreaCode      string         `gorm:"column:area_code" json:"area_code"`
	Properties    string         `gorm:"null;column:properties" json:"properties,omitempty"`
	Tags          string         `gorm:"null;column:tags" json:"tags,omitempty"`
	CreatedAt     time.Time      `gorm:"column:created_at" json:"created_at,omitempty"`
	UpdatedAt     time.Time      `gorm:"autoUpdateTime:milli;column:updated_at" json:"updated_at,omitempty"`
	DeletedAt     gorm.DeletedAt `gorm:"null;index;column:deleted_at" json:"deleted_at,omitempty"`
}

func (MArea) TableName() string {
	return "public.m_areas"
}
