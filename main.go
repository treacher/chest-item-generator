package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"chest-item-generator/model"
	"chest-item-generator/service"
)

func getenv(key, fallback string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return fallback
	}
	return value
}

func main() {
	file, e := ioutil.ReadFile("./items.json")

	if e != nil {
		fmt.Printf("File error: %v\n", e)
		os.Exit(1)
	}

	itemMap := make(map[string][]model.Item)
	json.Unmarshal(file, &itemMap)

	config := &service.Config{
		ItemMap: itemMap,
	}

	rootRouter := service.Routes(config)

	server := &http.Server{
		Addr:    ":8080",
		Handler: rootRouter,
	}

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		err := server.ListenAndServe()
		if err != nil {
		}
	}()

	<-interrupt
	deadline := time.Now().Add(3 * time.Second)
	ctx, cancel := context.WithDeadline(context.Background(), deadline)
	defer cancel()
	err := server.Shutdown(ctx)
	if err != nil {
	}

}
