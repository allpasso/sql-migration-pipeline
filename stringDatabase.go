package main

import (
	"log"
	"os"
)

func GetDatabase() string {
	database := os.Getenv("DB_DATABASE")
	if database == "" {
		log.Fatal("DB_DATABASE not found")
	}
	return database
}

func GetStringDatabase(withDatabse bool) string {
	host := os.Getenv("DB_HOST")
	if host == "" {
		log.Fatal("DB_HOST not found")
	}
	port := os.Getenv("DB_PORT")
	if port == "" {
		log.Fatal("DB_PORT not found")
	}
	user := os.Getenv("DB_USER")
	if user == "" {
		log.Fatal("DB_USER not found")
	}
	pass := os.Getenv("DB_PASS")
	if pass == "" {
		log.Fatal("DB_PASS not found")
	}
	database := os.Getenv("DB_DATABASE")
	if database == "" {
		log.Fatal("DB_DATABASE not found")
	}
	if withDatabse {
		return user + ":" + pass + "@tcp(" + host + ":" + port + ")/" + database + "?charset=utf8&parseTime=True&loc=Local"
	} else {
		return user + ":" + pass + "@tcp(" + host + ":" + port + ")/?charset=utf8&parseTime=True&loc=Local"
	}

}
