package main

import (
	"exec-server/common"
	"exec-server/controller"
	gr "exec-server/router"
	"log"
	"net/http"
)

func main() {
	healthController := common.GetHealthController()
	execController := controller.GetExecController()

	router := gr.NewRouter(healthController, execController)
	log.Fatal(http.ListenAndServe(":8080", router))
}
