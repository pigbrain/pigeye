package core

import (
	"fmt"
	"log"
	"net/http"
	"html/template"

	"pigeye/common"
)

var funcMap = template.FuncMap{
	"inc": func(i int) int {
		return i + 1
	},

	"mod": func(i, j int) bool {
		return i % j == 0
	},
}

func Render(writer http.ResponseWriter, view string, params interface{}) {
	//writer.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
	//writer.Header().Set("Pragma", "no-cache")
	//writer.Header().Set("Expires", "0")

	templates := [] string{
		common.TEMPLATE_FILE_PATH + "template.html",
		fmt.Sprintf(common.TEMPLATE_FILE_PATH + "%s.html", view)}

	t := template.Must(template.New("template.html").Funcs(funcMap).ParseFiles(templates...))

	err := t.Execute(writer, params)
	if err != nil {
		log.Print("template executing error: ", err)
	}
}

func Redirect(writer http.ResponseWriter, url string, request *http.Request) {
	http.Redirect(writer, request, url, 301)
}

