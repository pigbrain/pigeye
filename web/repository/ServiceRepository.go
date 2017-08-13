package repository

import (
	"pigeye/db"
	"pigeye/model"
	"log"
	"database/sql"
)

func SelectServiceCardList() []model.ServiceCard {
	dbConnection := db.GetConnection()
	defer db.ReleaseConnection(dbConnection)

	stmtOut, err := dbConnection.Prepare("SELECT service_id, name, description FROM service")
	defer stmtOut.Close()

	if err != nil {
		panic(err.Error())
	}

	var (
		serviceId int64
		name string
		description string
	)

	rows, err := stmtOut.Query()
	if err == sql.ErrNoRows {
		return nil
	}

	if err != nil {
		panic(err.Error())
	}

	var cards []model.ServiceCard
	for rows.Next() {
		err := rows.Scan(&serviceId, &name, &description)
		if err != nil {
			log.Fatal(err)
		}

		cards = append(cards, model.ServiceCard{
			ServiceId : serviceId,
			Name : name,
			Description : description})
	}

	return cards
}

func InsertService(name *string, description*string) {
	dbConnection := db.GetConnection()
	defer db.ReleaseConnection(dbConnection)

	stmtIns, err := dbConnection.Prepare("INSERT INTO service(name, description, creation_datetime, updated_datetime) VALUES( ?, ?, NOW(), NOW())")
	defer stmtIns.Close()

	if err != nil {
		panic(err.Error())
	}

	_, err = stmtIns.Exec(name, description)
	if err != nil {
		panic(err.Error())
	}
}

