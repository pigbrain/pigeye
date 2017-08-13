package handler

import (
	"net/http"
	"io"
	"time"
	"pigeye/common"
)

func Static(writer http.ResponseWriter, request *http.Request) {
	staticFile := request.URL.Path[len(common.STATIC_URL):]

	if len(staticFile) != 0 {
		file, err := http.Dir(common.STATIC_FILE_PATH).Open(staticFile)
		if err == nil {
			content := io.ReadSeeker(file)
			http.ServeContent(writer, request, staticFile, time.Now(), content)
			return
		}
	}

	http.NotFound(writer, request)
}
