package repository

import (
	"log"

	"database/sql"

	"pigeye/db"
	"pigeye/model"
)

func SelectApiCount() int {
	dbConnection := db.GetConnection()
	defer db.ReleaseConnection(dbConnection)

	stmtOut, err := dbConnection.Prepare("SELECT COUNT(*) FROM api")
	defer stmtOut.Close()

	if err != nil {
		panic(err.Error())
	}

	var (
		count int
	)

	err = stmtOut.QueryRow().Scan(&count)
	if err != nil {
		log.Fatal(err)
	}

	return count
}

func SelectApiCardList(serviceId *int64) []model.ApiCard {
	dbConnection := db.GetConnection()
	defer db.ReleaseConnection(dbConnection)

	stmtOut, err := dbConnection.Prepare("SELECT api_id, name, description, method, url, success FROM api WHERE service_id = ?")
	defer stmtOut.Close()

	if err != nil {
		panic(err.Error())
	}

	rows, err := stmtOut.Query(serviceId)
	if err == sql.ErrNoRows {
		return nil
	}

	if err != nil {
		panic(err.Error())
	}

	var cards []model.ApiCard
	var (
		apiId int64
		name string
		description string
		method string
		url string
		success int8
	)

	for rows.Next() {
		err := rows.Scan(&apiId, &name, &description, &method, &url, &success)
		if err != nil {
			log.Fatal(err)
		}

		cards = append(cards,
			model.ApiCard{
				ApiId           : apiId,
				Name            : name,
				Description     : description,
				Method          : method,
				Url             : url,
				Success         : success})
	}

	return cards
}

func SelectApiCard(apiId *int64, serviceId *int64) *model.ApiCard {
	dbConnection := db.GetConnection()
	defer db.ReleaseConnection(dbConnection)

	stmtOut, err := dbConnection.Prepare("SELECT name, description, url, content_type,  method, request_body, status, response_body FROM api WHERE api_id = ? and service_id = ?")
	defer stmtOut.Close()

	if err != nil {
		panic(err.Error())
	}

	var (
		name string
		description string
		url string
		contentType string
		method string
		requestBody string
		status int
		responseBody string
	)

	err = stmtOut.QueryRow(apiId, serviceId).Scan(&name, &description, &url, &contentType, &method, &requestBody, &status, &responseBody)

	if err == sql.ErrNoRows {
		return nil
	}

	if err != nil {
		panic(err.Error())
	}

	return &model.ApiCard{
		ServiceId: *serviceId,
		ApiId: *apiId,
		Name: name,
		Description: description,
		Method : method,
		ContentType:contentType,
		Url: url,
		RequestBody:requestBody,
		Status:status,
		ResponseBody:responseBody,
	}
}

func UpdateApi(apiCard *model.ApiCard) {
	dbConnection := db.GetConnection()
	defer db.ReleaseConnection(dbConnection)

	query := "UPDATE api SET name = ?, description = ?, url = ?, user_agent = ?, content_type = ?, method = ?, request_body = ?, status = ?, response_body = ?, creation_datetime = NOW(), updated_datetime = NOW() "
	query += "WHERE api_id = ? AND service_id = ?"

	stmtUpt, err := dbConnection.Prepare(query)
	defer stmtUpt.Close()

	if err != nil {
		panic(err.Error())
	}

	_, err = stmtUpt.Exec(apiCard.Name,
		apiCard.Description,
		apiCard.Url,
		apiCard.UserAgent,
		apiCard.ContentType,
		apiCard.Method,
		apiCard.RequestBody,
		apiCard.Status,
		apiCard.ResponseBody,
		apiCard.ApiId,
		apiCard.ServiceId)

	if err != nil {
		panic(err.Error())
	}
}

func InsertApi(apiCard *model.ApiCard) {
	dbConnection := db.GetConnection()
	defer db.ReleaseConnection(dbConnection)

	query := "INSERT INTO api(service_id, name, description, url, user_agent, content_type, method, request_body, status, response_body, creation_datetime, updated_datetime) "
	query += "VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, NOW(), NOW())"

	stmtIns, err := dbConnection.Prepare(query)
	defer stmtIns.Close()

	if err != nil {
		panic(err.Error())
	}

	_, err = stmtIns.Exec(
		apiCard.ServiceId,
		apiCard.Name,
		apiCard.Description,
		apiCard.Url,
		apiCard.UserAgent,
		apiCard.ContentType,
		apiCard.Method,
		apiCard.RequestBody,
		apiCard.Status,
		apiCard.ResponseBody)

	if err != nil {
		panic(err.Error())
	}
}

func UpdateApiResult(apiId *int64, serviceId *int64, success bool) {
	dbConnection := db.GetConnection()
	defer db.ReleaseConnection(dbConnection)

	stmtUpt, err := dbConnection.Prepare("UPDATE api SET success = ? WHERE api_id = ? AND service_id = ? ")
	defer stmtUpt.Close()

	if err != nil {
		panic(err.Error())
	}

	_, err = stmtUpt.Exec(success, apiId, *serviceId)

	if err != nil {
		panic(err.Error())
	}
}