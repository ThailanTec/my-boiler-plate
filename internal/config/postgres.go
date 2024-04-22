package config

import (
	"context"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Postgres struct {
	db *gorm.DB
}

func NewPostgresRepository(ctx context.Context, dsn string) (*Postgres, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return &Postgres{db: db}, nil
}

func (r *Postgres) GetDB() *gorm.DB {
	return r.db
}
