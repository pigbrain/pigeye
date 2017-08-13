package handler

import (
	"net/http"

	"pigeye/model"
	"pigeye/web/core"
	"pigeye/web/repository"
)

func Index(writer http.ResponseWriter, _ *http.Request) {
	var cards []model.ServiceCard = repository.SelectServiceCardList()

	core.Render(writer, "Index", struct{ Cards []model.ServiceCard }{cards})
}

