package db

import (
	"log"
	"time"
	"container/list"
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

var connectionManageChannel = make(chan *sql.DB)
var connectionRequestChannel = make(chan bool)
var connectionResponseChannel = make(chan *sql.DB)

var connectionList = list.New()

type DBConfig struct {
	Id       string
	Password string
	Name     string
}

func (dbConfig *DBConfig) CreationConnection() *sql.DB {
	db, err := sql.Open("mysql", dbConfig.Id + ":" + dbConfig.Password + "@/" + dbConfig.Name)
	if err != nil {
		log.Fatal(err)
		return nil
	} else {
		return db
	}
}

var dbConfig DBConfig = DBConfig{}

func Create(id string, password string, name string, count int) {
	dbConfig.Id = id
	dbConfig.Password = password
	dbConfig.Name = name

	go manageConnection(connectionManageChannel)

	for i := 0; i < count; i++ {
		db := dbConfig.CreationConnection();

		connectionManageChannel <- db
	}
}

func manageConnection(connectoinManageChannel chan *sql.DB) {
	defer func() {
		for element := connectionList.Front(); element != nil; element = element.Next() {
			db := element.Value.(*sql.DB)
			db.Close()
		}
	}()

	for {
		select {
		case db := <-connectoinManageChannel:
			connectionList.PushBack(db)
			log.Println("db create");

		case <-connectionRequestChannel:
			if connectionList.Len() > 0 {
				connectionResponseChannel <- connectionList.Back().Value.(*sql.DB)
			} else {
				connectionResponseChannel <- dbConfig.CreationConnection()
			}

		case <-time.After(time.Second * 2):
			if (connectionList.Len() > 0) {
				db := connectionList.Front().Value.(*sql.DB);
				err := db.Ping()
				if err != nil {
					log.Fatal(err)
					db.Close()
				} else {
					connectionList.PushBack(db)
				}
			}
		}
	}
}

func GetConnection() *sql.DB {
	connectionRequestChannel <- true

	return <-connectionResponseChannel
}

func ReleaseConnection(db *sql.DB) {
	connectionManageChannel <- db
}

