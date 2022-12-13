package main

import (
	"fmt"
	"io/fs"
	"log"
	"os"

	"github.com/joho/godotenv"
)

var err error

func main() {
	if os.Getenv("ENVIRONMENT") == "" {
		errEnv := godotenv.Load(".env")
		if errEnv != nil {
			log.Fatal(errEnv.Error())
		}
	}

	CheckDatabase()

	var migrationHistory []MigrationHistory
	result := Con.Find(&migrationHistory)
	if result.Error != nil {
		log.Fatal("Error get migration history: " + result.Error.Error())
	}

	var filesNotExecuted []fs.FileInfo
	files := GetFilesSQL()
	for _, file := range files {
		executed := false
		for _, migrateFile := range migrationHistory {
			if migrateFile.Id == file.Name() {
				executed = true
			}
		}
		if !executed {
			filesNotExecuted = append(filesNotExecuted, file)
		}
	}

	if len(filesNotExecuted) > 0 {
		fmt.Println("------------------------------------")
		fmt.Println("Files for processing:")
		for _, file := range filesNotExecuted {
			fmt.Println("- " + file.Name())
		}
		fmt.Println("------------------------------------")
	} else {
		fmt.Println("------------------------------------")
		fmt.Println("No files processing:")
		fmt.Println("------------------------------------")
		os.Exit(0)
	}

	for _, file := range filesNotExecuted {
		fmt.Println("------------------------------------")
		fmt.Println("Exec file: " + file.Name())
		sql := LoadFileSql(file.Name())
		err := Con.Exec(sql)
		if err.Error != nil {
			fmt.Println("Error exec SQL file: " + file.Name())
			log.Fatalln(err.Error.Error())
		}
		result := Con.Create(&MigrationHistory{Id: file.Name()})
		if result.Error != nil {
			fmt.Println("Error insert migration_histories: " + file.Name())
			log.Fatalln(result.Error.Error())
		}
		fmt.Println("------------------------------------")
	}

	os.Exit(0)
}
