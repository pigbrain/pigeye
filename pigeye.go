package main

import (
	"io/ioutil"
	"log"
	"net/http"

	"gopkg.in/yaml.v2"

	"pigeye/common"
	"pigeye/db"
	"pigeye/model"
	"pigeye/watcher"
	"pigeye/web/handler"
)

var (
	Version = "1.0.0"
	Build   = "2017-08-12"
)

func main() {
	configFile, err := ioutil.ReadFile("config.yml")
	config := model.Config{}

	err = yaml.Unmarshal(configFile, &config)
	log.Print(config)

	db.Create(config.DB.Ip, config.DB.Port, config.DB.Id, config.DB.Password, config.DB.Name, config.DB.PoolSize)
	watcher.Create(config.Monitor.PoolSize)

	http.HandleFunc(common.INDEX_URL, handler.Index)
	http.HandleFunc(common.SERVICE_REGISTER_URL, handler.ServiceRegister)
	http.HandleFunc(common.API_REGISTER_URL, handler.ApiRegister)
	http.HandleFunc(common.API_LIST_URL, handler.ApiList)
	http.HandleFunc(common.STATIC_URL, handler.Static)

	err = http.ListenAndServe(common.HTTP_PORT, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

}
