package main

import (
	"test/alterstay/config"
	"test/alterstay/factory"
	"test/alterstay/migration"
	"test/alterstay/routes"
)

func main() {
	dbConn := config.InitDB()
	migration.InitMigrate(dbConn)
	presenter := factory.InitFactory(dbConn)
	e := routes.New(presenter)

	e.Logger.Fatal(e.Start(":8000"))
}
