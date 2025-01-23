package main

import (
	"github.com/ArdiSasongko/EwalletProjects-notification/cmd/api"
)

func main() {
	// setup grpc
	api.SetupGRPC()

	// setup http
	//api.SetupHTTP()

}
