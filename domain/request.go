package domain

type Request struct {
	Authorization string
	Text          string `form:"text"`
	Folder        string `param:"folder"`
}
