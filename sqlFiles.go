package main

import (
	"io/fs"
	"io/ioutil"
	"log"
)

func getPathSql() string {
	path := "./sql"
	if path == "" {
		log.Fatal("DB_HOST not found")
	}
	return path
}

func GetFilesSQL() []fs.FileInfo {
	path := getPathSql()

	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}

	return files
}

func LoadFileSql(fileName string) string {
	pathFile := getPathSql() + "/" + fileName
	fileByte, err := ioutil.ReadFile(pathFile)
	if err != nil {
		log.Fatalln("Error load file SQL '" + pathFile + "': " + err.Error())
	}
	return string(fileByte)
}
