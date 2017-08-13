package handler

import (
	"net/http"
	"strconv"

	"pigeye/common"
	"pigeye/web/core"
	"pigeye/model"
	"pigeye/web/repository"
)

func ApiList(writer http.ResponseWriter, request *http.Request) {
	switch request.Method {
	case common.HTTP_METHOD_GET :
		viewApiList(writer, request)
	}
}

func viewApiList(writer http.ResponseWriter, request *http.Request) {
	serviceId, _ := strconv.ParseInt(request.FormValue("serviceId"), 10, 64)

	var cards []model.ApiCard = repository.SelectApiCardList(&serviceId)

	core.Render(writer, "ApiList", struct{ ServiceId int64; Cards []model.ApiCard }{serviceId, cards })
}



