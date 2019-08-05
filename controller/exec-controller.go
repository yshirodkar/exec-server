package controller

import (
	"net/http"
	"bytes"
	"fmt"
	"log"
	"os/exec"
	"strings"
)

/*
	This class helps to determine service health
*/
type IExecController interface {
	GetDockerImage(w http.ResponseWriter, r *http.Request)
}

type execController struct {}

/*
	Return an implementation of IExecController
*/
func GetExecController() IExecController {
	return &execController{}
}

/*
	This returns a UTF-8 charset response: "Alive"
*/
func (this *execController) GetDockerImage(w http.ResponseWriter, r *http.Request) {

	cmd := exec.Command("docker", "ps")
	cmd.Stdin = strings.NewReader("some input")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}

	result := []byte(fmt.Sprintf("%q \n", out.String()))

	w.Header().Set("Content-Type", "charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}
