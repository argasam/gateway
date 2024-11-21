package model

type JSONResponse struct {
	Code   int         `json:"Code"`
	Result interface{} `json:"Result"`
}
