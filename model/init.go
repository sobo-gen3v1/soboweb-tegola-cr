package model

import (
	"fmt"
	"sync"

	"github.com/sobo-gen3v1/soboweb-tegola-cr/internal/ttools"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var PgConn *PgDb

type PgDb struct {
	mtx      *sync.Mutex
	Instance *gorm.DB
}

func NewPgDB(dsn string) (*PgDb, error) {
	p := &PgDb{
		mtx:      new(sync.Mutex),
		Instance: nil,
	}
	p.mtx.Lock()
	defer p.mtx.Unlock()

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, err
	}
	p.Instance = db

	return p, nil
}

func GetPgDb() (*gorm.DB, error) {
	if PgConn == nil {
		dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
			ttools.GetEnvDefault("DB_HOST", "localhost"),
			ttools.GetEnvDefault("DB_USER", "postgres"),
			ttools.GetEnvDefault("DB_PASSWORD", "postgres"),
			ttools.GetEnvDefault("DB_NAME", "sobo"),
			ttools.GetEnvDefault("DB_PORT", "5432"),
			ttools.GetEnvDefault("TEGOLA_POSTGIS_SSL", "disable"),
		)

		db, err := NewPgDB(dsn)

		if err != nil {
			return nil, err
		}

		// migrate table
		// db.AutoMigrate(&TLayer{})
		PgConn = db
	}

	return PgConn.Instance, nil
}
