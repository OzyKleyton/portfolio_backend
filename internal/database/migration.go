package database

import (
	"context"

	"github.com/ozykt4/portfolio_backend/config"
	"github.com/ozykt4/portfolio_backend/config/db"
	"github.com/ozykt4/portfolio_backend/internal/model"
)

func InitMigrate() error {

	db, err := db.ConnectDB(config.GetConfig().DBURL)
	if err != nil {
		return err
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	db = db.WithContext(ctx)

	if err := db.AutoMigrate(
		&model.Project{},
	); err != nil {
		return err
	}

	return err
}
