package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"pigeye/common"
	"pigeye/model"
	"pigeye/web/core"
	"pigeye/web/repository"
)

func ApiRegister(writer http.ResponseWriter, request *http.Request) {
	switch request.Method {
	case common.HTTP_METHOD_GET:
		viewApiRegister(writer, request)
	case common.HTTP_METHOD_POST:
		doRegisterApi(writer, request)
	}
}

func viewApiRegister(writer http.ResponseWriter, request *http.Request) {
	serviceId, _ := strconv.ParseInt(request.FormValue("serviceId"), 10, 64)
	apiId, _ := strconv.ParseInt(request.FormValue("apiId"), 10, 64)

	empty := struct {
		ServiceId int64
		ApiId     int64
	}{
		ServiceId: serviceId,
		ApiId:     0}

	if apiId == 0 {
		core.Render(writer, "ApiRegister", empty)
		return
	}

	apiCard := repository.SelectApiCard(&apiId, &serviceId)
	if apiCard == nil {
		core.Render(writer, "ApiRegister", empty)
		return
	}

	core.Render(writer, "ApiRegister", apiCard)
}

func doRegisterApi(writer http.ResponseWriter, request *http.Request) {
	apiId, _ := strconv.ParseInt(request.FormValue("apiId"), 10, 64)
	serviceId, _ := strconv.ParseInt(request.FormValue("serviceId"), 10, 64)
	name := request.FormValue("name")
	description := request.FormValue("description")
	url := request.FormValue("url")
	userAgent := request.FormValue("userAgent")
	contentType := request.FormValue("contentType")
	method := request.FormValue("method")
	requestBody := request.FormValue("requestBody")
	status, _ := strconv.Atoi(request.FormValue("status"))
	responseBody := request.FormValue("responseBody")
	notificationScript := request.FormValue("notificationScript")

	apiCard := model.ApiCard{
		ServiceId:          serviceId,
		ApiId:              apiId,
		Name:               name,
		Description:        description,
		Method:             method,
		ContentType:        contentType,
		UserAgent:          userAgent,
		Url:                url,
		RequestBody:        requestBody,
		Status:             status,
		ResponseBody:       responseBody,
		NotificationScript: notificationScript,
	}

	if apiId > 0 {
		repository.UpdateApi(&apiCard)
	} else {
		repository.InsertApi(&apiCard)
	}

	redirectUrl := fmt.Sprintf("%s?serviceId=%d&apiId=%d", common.API_LIST_URL, serviceId, apiId)

	core.Redirect(writer, redirectUrl, request)
}
