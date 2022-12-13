package main

import (
	"fmt"
	"log"
)

func CheckDatabase() {
	errCon := ConnectToDatabase(false)
	if errCon != nil {
		log.Fatal(errCon.Error())
	}

	dbName := GetDatabase()

	var result []interface{}
	errExists := Con.Raw("SELECT SCHEMA_NAME FROM INFORMATION_SCHEMA.SCHEMATA WHERE SCHEMA_NAME = ?;", dbName).Scan(&result)
	if errExists.Error != nil {
		log.Fatal("Error check database exists: " + errExists.Error.Error())
	}
	if len(result) <= 0 {
		fmt.Println("database not exits")
		fmt.Println("create database")
		errCreateDb := Con.Exec("CREATE SCHEMA `" + dbName + "` DEFAULT CHARACTER SET utf8 COLLATE utf8_bin;")
		if errCreateDb.Error != nil {
			log.Fatal("Error create database: " + errCreateDb.Error.Error())
		}

		CloseConnection()
		errCon := ConnectToDatabase(true)
		if errCon != nil {
			log.Fatal(errCon.Error())
		}

		errCreateMgHistory := Con.Exec("CREATE TABLE `migration_histories` (`id` VARCHAR(100) NOT NULL,`runner_at` DATETIME NOT NULL,PRIMARY KEY (`id`));")
		if errCreateMgHistory.Error != nil {
			log.Fatal("Error table migration_histories: " + errCreateMgHistory.Error.Error())
		}

	} else {
		CloseConnection()
		errCon := ConnectToDatabase(true)
		if errCon != nil {
			log.Fatal(errCon.Error())
		}
	}
}
