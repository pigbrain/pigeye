package model

type ApiCard struct {
	ServiceId    int64
	ApiId        int64
	Name         string
	Description  string
	Method       string
	ContentType  string
	UserAgent	 string
	Url          string
	RequestBody  string
	Status       int
	ResponseBody string
	Success      int8
}