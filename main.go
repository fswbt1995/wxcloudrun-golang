package main

import (
	"fmt"
	"log"
	"net/http"
	"wxcloudrun-golang/db"
	"wxcloudrun-golang/service"
)

func main() {
	if err := db.Init(); err != nil {
		panic(fmt.Sprintf("mysql init failed with %+v", err))
	}

	mtrService := service.NewMTRService()
	
	http.HandleFunc("/", service.IndexHandler)
	http.HandleFunc("/api/count", service.CounterHandler)
	http.HandleFunc("/mtr/schedule", mtrService.HandleMTRSchedule)

	log.Fatal(http.ListenAndServe(":80", nil))
}
