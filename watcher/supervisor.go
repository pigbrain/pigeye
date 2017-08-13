package watcher

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"reflect"
	"time"

	"pigeye/common"
	"pigeye/model"
	"pigeye/web/repository"
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

	go func(tickChannel <-chan bool, quotaChannel chan<- model.Quota) {
		defer close(quotaChannel)

		var remainRefreshSecond = common.CACHE_REFRESH_SECOND

		apiCount := repository.SelectApiCount()

		for range tickChannel {
			remainRefreshSecond--
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
					Start: index,
					Count: common.API_QUOTA_PER_WORK,
				}

				index += common.API_QUOTA_PER_WORK
			}
		}
	}(tickChannel, quotaChannel)

	for i := 0; i < count; i++ {
		go worker(quotaChannel)
	}
}

func worker(quotaChannel <-chan model.Quota) {
	for quota := range quotaChannel {
		start := quota.Start
		count := quota.Count

		apiList := repository.SelectApiList(start, count)
		for _, api := range apiList {
			var client = &http.Client{
				Timeout: time.Second * 10,
			}

			request, err := http.NewRequest(api.Method, api.Url, bytes.NewBufferString(api.RequestBody))
			if err != nil {

			}

			if len(api.ContentType) > 0 {
				request.Header.Set("Content-Type", api.ContentType)
			}

			if len(api.UserAgent) > 0 {
				request.Header.Set("User-Agent", api.UserAgent)
			}

			response, err := client.Do(request)
			log.Println(response)

			if response == nil {
				log.Print("response is nil..")
				continue
			}

			if api.Status != response.StatusCode {
				log.Print("Status(", api.Status, ") !")
				repository.UpdateApiResult(&api.ApiId, &api.ServiceId, false)

				continue
			}

			if len(api.ResponseBody) == 0 {
				// success
				// don't compare any more
				return
			}

			body, err := ioutil.ReadAll(response.Body)
			if err != nil {

			}

			result, err := AreEqualJSON(string(body), api.ResponseBody)
			if err != nil {

			}

			if !result {

			}

			repository.UpdateApiResult(&api.ApiId, &api.ServiceId, true)
		}
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
