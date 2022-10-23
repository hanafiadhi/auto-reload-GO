package main

import (
	"net/http"
	"time"

	"github.com/hanafiadhi/auto-reload-GO/handler"
	"github.com/hanafiadhi/auto-reload-GO/helper"
	"github.com/hanafiadhi/auto-reload-GO/service"
)

func main() {
	go autoUpdate()
	http.HandleFunc("/", handler.Handler)

	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		panic(err)
	}
}

func autoUpdate() {
	disaster := helper.GetDisaster()
	for range time.Tick(15 * time.Second) {
		service.UpdateJSON(disaster)
	}
}
