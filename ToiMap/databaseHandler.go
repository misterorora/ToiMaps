package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

const (
	host     = "85.214.35.122"
	port     = 5432
	user     = "postgres"
	password = "toimap"
	dbname   = "toimapsdb"
)

func requestMarkers() ([]Marker, error) {
	var result []Marker

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, err
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	rows, err := db.Query("SELECT (Latitude, Longitude, Name, Art, Anzahl) FROM toimapsdb.public.toilette") //Latitude, Longitude, Name, Art, Anzahl
	if err != nil {
		log.Fatalf("Error: Unable to execute query: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var latitude string
		var longitude string
		var name string
		var art string
		var anzahl int
		rows.Scan(&latitude, &longitude, &name, &art, &anzahl)

		marker := Marker{Name: name, Latitude: latitude, Longitude: longitude, Art: art, Anzahl: anzahl}
		result = append(result, marker)

	}

	return result, nil
}

type DatabaseClient struct {
}
