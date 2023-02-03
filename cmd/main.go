package main

import (
	"fmt"
	"net/http"

	"session.com/config"
	"session.com/package/controller"

	"github.com/gorilla/mux"
	"github.com/tkanos/gonfig"
)

func main() {

	configuration := config.Configuration{}
	err := gonfig.GetConf("config.yml", &configuration)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Active:", configuration.Active)
	fmt.Println("Port:", configuration.Server.Port)

	router := mux.NewRouter()

	router.HandleFunc("/session", controller.CreateSession).Methods("POST")
	router.HandleFunc("/session", controller.FindSessions).Methods("GET")
	router.HandleFunc("/session/:id", controller.FindSessionById)
	router.HandleFunc("/session/:id", controller.DeleteSession).Methods("DELETE")

	http.ListenAndServe(":8088", router)

}
