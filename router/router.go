package router

import (
	"exec-server/common"
	"exec-server/controller"
	"github.com/gorilla/mux"
	"net/http"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

func NewRouter(
	healthController common.IHealthController,
	execController controller.IExecController,
) *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	DefineHealthRoutes(router, healthController)
	DefineRoutes(router, execController)
	return router
}

func DefineHealthRoutes(router *mux.Router, healthController common.IHealthController) {
	router.Methods("GET").Path("/health").HandlerFunc(healthController.GetHealth)
}

func DefineRoutes(router *mux.Router, execController controller.IExecController) {
	router.Methods("GET").Path("/pull-image").HandlerFunc(execController.GetDockerImage)
	router.Methods("GET").Path("/ps-image").HandlerFunc(execController.GetDockerImage)
}
