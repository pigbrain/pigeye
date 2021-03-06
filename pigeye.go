package main

import (
	"flag"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"gopkg.in/yaml.v2"

	"pigeye/common"
	"pigeye/db"
	"pigeye/model"
	"pigeye/watcher"
	"pigeye/web/core"
	"pigeye/web/handler"
)

var (
	Version = "1.0.0"
	Build   = "2017-08-12"
)

func main() {
	fpLog, err := os.OpenFile("logfile.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	defer fpLog.Close()
	core.Create(fpLog)
	core.GetLogger().Print("Adsfasdfasfas")

	configFlag := flag.String("config", "config.yml", "Please specify a config file (-config or --config)")
	flag.Parse()

	log.Print(*configFlag)

	configFile, err := ioutil.ReadFile(*configFlag)
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

	port := ":" + config.Http.Port
	err = http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

}
