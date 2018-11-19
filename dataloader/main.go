package main

import (
	"../databasemodels"
	"../jsonutil"
	"../logging"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func migrateSchema(db *gorm.DB) {
	logging.LogInfo("Migrating schema")
	db.AutoMigrate(
		&databasemodels.BusinessLine{},
		&databasemodels.Environment{},
		&databasemodels.AppCache{},
		&databasemodels.Log{},
		&databasemodels.IisSite{},
		&databasemodels.IisApplication{},
		&databasemodels.IisApplicationPool{},
		&databasemodels.ScheduledTask{},
		&databasemodels.WindowsService{},
		&databasemodels.Resource{},
		&databasemodels.Server{})
}

func flushData(db *gorm.DB) {
	logging.LogInfo("Flushing data")
	db.DropTable(&databasemodels.BusinessLine{})
	db.DropTable(&databasemodels.Environment{})
	db.DropTable(&databasemodels.AppCache{})
	db.DropTable(&databasemodels.IisSite{})
	db.DropTable(&databasemodels.IisApplication{})
	db.DropTable(&databasemodels.IisApplicationPool{})
	db.DropTable(&databasemodels.ScheduledTask{})
	db.DropTable(&databasemodels.WindowsService{})
	db.DropTable(&databasemodels.Resource{})
	db.DropTable(&databasemodels.Server{})
}

func loadData(db *gorm.DB) {
	logging.LogInfo("Loading data")
	// id, _ := uuid.NewV4()

}

func main() {
	logging.LogHeader("Data Loader")
	logging.CreateLog()

	var c Configuration
	jsonutil.DecodeJsonFromFile("./appsettings.json", &c)

	connStr := c.GetPostgresConnectionString()
	db, err := gorm.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	db.LogMode(true)

	flushData(db)
	migrateSchema(db)
	loadData(db)
}
