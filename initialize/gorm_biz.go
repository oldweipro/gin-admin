package initialize

import (
	"github.com/oldweipro/gin-admin/pkg/app"
)

func bizModel() error {
	db := app.DBClient
	err := db.AutoMigrate()
	if err != nil {
		return err
	}
	return nil
}
