package main

import (
	"github.com/DmitryBugrov/log"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"httpshell/config"
	"httpshell/controller"
	"net/http"
	"strconv"
)

//ConfigFile - config file name
const ConfigFile = "./config.json"

var (
	c config.Config
	l log.Log
)

func main() {
	l.Init(log.LogLevelInfo, true, true, true)
	l.Print(log.LogLevelTrace, "Starting...")
	c.Init(ConfigFile)
	if err := c.Load(l); err != nil {
		l.Print(log.LogLevelError, err)
		panic(0)
	}
	l.LogLevel = c.LogLevel

	router := mux.NewRouter()
	router.HandleFunc("/", controller.Shell).Methods("POST")
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))
	n := negroni.Classic()
	n.UseHandler(router)

	http.ListenAndServe(":"+strconv.Itoa(c.Port), n)

}
