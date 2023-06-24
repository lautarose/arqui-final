package dtos

type ResponseDto struct {
	NumFound      int  `json:"numFound"`
	Start         int  `json:"start"`
	NumFoundExact bool `json:"numFoundExact"`
	Docs          ItemsDto
}

type ResponsesDto []ResponseDto
