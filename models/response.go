package models

type Responses interface {
	IsOk() bool
	KnowStatus() int
}

func (r Response) IsOk() bool {
	if r.Message == "OK" {
		return true
	}
	return false

}

func (r Error) IsOk() bool {
	if r.Message == "OK" {
		return true
	}
	return false
}

func (r Response) KnowStatus() int {
	return r.Status
}

func (r Error) KnowStatus() int {
	return r.Status
}

type Error struct {
	Message string `json:"Message"`
	Status  int    `json:"ErrorCode"`
}

type Response struct {
	Message string      `json:"Message"`
	Body    interface{} `json:"Body"`
	Status  int         `json:"Status"`
}
