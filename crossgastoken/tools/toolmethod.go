package tools

import "poly-bridge/models"

func migrateTable(){
	err := db.Debug().AutoMigrate(
		&models.LockTokenStatistic{},
	)
	checkError(err, "Creating tables")
}
