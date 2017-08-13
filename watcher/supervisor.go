package watcher

import (
	"log"
	"time"
	"fmt"
	"reflect"
	"net/http"
	"io/ioutil"
	"encoding/json"

	"pigeye/db"
	"pigeye/common"
	"pigeye/web/repository"
	"pigeye/model"
)

var tickChannel = make(chan bool)
var quotaChannel = make(chan model.Quota, 10)

func Create(count int) {
	ticker := time.NewTicker(time.Second * 1)
	go func(ticker *time.Ticker) {
		defer ticker.Stop()
		defer close(tickChannel)

		for range ticker.C {
			tickChannel <- true
		}
	}(ticker)

	go func(tickChannel <- chan bool, quotaChannel chan <- model.Quota) {
		defer close(quotaChannel)

		var remainRefreshSecond = common.CACHE_REFRESH_SECOND;

		apiCount := repository.SelectApiCount()

		for range tickChannel {
			remainRefreshSecond--;
			if remainRefreshSecond > 0 {
				continue
			}

			remainRefreshSecond = common.CACHE_REFRESH_SECOND

			apiCount = repository.SelectApiCount()
			if apiCount <= 0 {
				continue
			}

			index := 0
			for index < apiCount {
				quotaChannel <- model.Quota{
					Start : index,
					Count : common.API_QUOTA_PER_WORK,
				}

				index += common.API_QUOTA_PER_WORK;
			}
		}
	}(tickChannel, quotaChannel)

	for i := 0; i < count; i++ {
		go worker(quotaChannel)
	}
}

func worker(quotaChannel <-chan model.Quota) {
	for quota := range quotaChannel {
		dbConnection := db.GetConnection()
		stmtOut, err := dbConnection.Prepare("SELECT api_id, service_id, url, user_agent, content_type,  method, request_body, status, response_body FROM api LIMIT ?, ?")

		start := quota.Start
		count := quota.Count

		var (
			apiId int64
			serviceId int64
			url string
			userAgent string
			contentType string
			method string
			requestBody string
			status int
			responseBody string
		)

		err = stmtOut.QueryRow(start, count).Scan(&apiId, &serviceId, &url, &userAgent, &contentType, &method, &requestBody, &status, &responseBody)

		db.ReleaseConnection(dbConnection)
		stmtOut.Close()

		if err != nil {
			panic(err.Error())
		}

		var client = &http.Client{
			Timeout: time.Second * 10,
		}

		request, err := http.NewRequest(method, url, nil)
		if err != nil {

		}

		if len(contentType) > 0 {
			request.Header.Set("Content-Type", contentType)
		}

		if len(userAgent) > 0 {
			request.Header.Set("User-Agent", userAgent)
		}

		response, err := client.Do(request)
		log.Println(response)

		if (response == nil) {
			log.Print("response is nil..")
			continue
		}

		if (status != response.StatusCode) {
			log.Print("Status(", status, ") !")
			repository.UpdateApiResult(&apiId, &serviceId, false)

			continue
		}

		if (len(responseBody) == 0) {
			// success
			// don't compare any more
			return
		}

		body, err := ioutil.ReadAll(response.Body)
		if err != nil {

		}

		result, err := AreEqualJSON(string(body), responseBody)
		if err != nil {

		}

		if !result {

		}

		repository.UpdateApiResult(&apiId, &serviceId, true)
	}
}

func AreEqualJSON(s1, s2 string) (bool, error) {
	var o1 interface{}
	var o2 interface{}

	var err error
	err = json.Unmarshal([]byte(s1), &o1)
	if err != nil {
		return false, fmt.Errorf("Error mashalling string 1 :: %s", err.Error())
	}
	err = json.Unmarshal([]byte(s2), &o2)
	if err != nil {
		return false, fmt.Errorf("Error mashalling string 2 :: %s", err.Error())
	}

	return reflect.DeepEqual(o1, o2), nil
}