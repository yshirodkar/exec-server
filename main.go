package main

import (
	"log"
	"net/http"
	"grpc-exec/common"
	gr "grpc-exec/router"
	"grpc-exec/controller"
)

func main() {
	healthController := common.GetHealthController()
	execController := controller.GetExecController()

	router := gr.NewRouter(healthController, execController)
	log.Fatal(http.ListenAndServe(":8080", router))
}
