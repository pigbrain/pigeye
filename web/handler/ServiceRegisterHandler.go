package handler

import (
	"log"
	"net/http"

	"pigeye/common"
	"pigeye/web/core"
	"pigeye/web/repository"
)

func ServiceRegister(writer http.ResponseWriter, request *http.Request) {
	switch request.Method {
	case common.HTTP_METHOD_GET :
		viewServiceRegister(writer, request)
	case common.HTTP_METHOD_POST :
		doRegisterService(writer, request)
	}
}

func viewServiceRegister(writer http.ResponseWriter, request *http.Request) {
	core.Render(writer, "ServiceRegister", nil)
}

func doRegisterService(writer http.ResponseWriter, request *http.Request) {
	name := request.FormValue("name")
	description := request.FormValue("description")

	log.Println("# doRegister params")
	log.Println("name=" + name);

	repository.InsertService(&name, &description);

	core.Redirect(writer, common.INDEX_URL, request)
}


