package api

import "os"

type Api struct {
	Address string
	Port    string
}

func NewApi() (*Api, error) {
	address := os.Getenv("ADDRESS")
	port := os.Getenv("PORT")

	return &Api{
		Address: address,
		Port:    port,
	}, nil
}
