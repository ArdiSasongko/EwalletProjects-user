package main

import (
	"github.com/ArdiSasongko/EwalletProjects-user/cmd/api"
)

func main() {
	// setup grpc
	go api.SetupGRPC()

	// setup http
	api.SetupHTTP()

}
