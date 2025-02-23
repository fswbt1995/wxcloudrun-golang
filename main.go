package main

import (
	"fmt"
	"log"
	"net/http"
	"wxcloudrun-golang/config"
	"wxcloudrun-golang/db"
	"wxcloudrun-golang/service"
)

func main() {
	if err := db.Init(); err != nil {
		panic(fmt.Sprintf("mysql init failed with %+v", err))
	}

	mtrService := service.NewMTRService(config.MTR_API_BASE_URL)
	
	http.HandleFunc("/", service.IndexHandler)
	http.HandleFunc("/api/count", service.CounterHandler)
	http.HandleFunc("/mtr/schedule", mtrService.HandleMTRSchedule)

	log.Printf("服务启动在%s端口...\n", config.SERVER_PORT)
	log.Fatal(http.ListenAndServe(config.SERVER_PORT, nil))
}
