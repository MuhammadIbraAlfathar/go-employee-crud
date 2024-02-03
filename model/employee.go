package model

type Employee struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	Address string `json:"address"`
	NPWP    string `json:"npwp"`
}
