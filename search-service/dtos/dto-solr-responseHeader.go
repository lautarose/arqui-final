package dtos

type ResponseHeaderDto struct {
	Status int `json:"status"`
	QTime  int `json:"QTime"`
}

type ResponsesHeadersDto []ResponseHeaderDto
