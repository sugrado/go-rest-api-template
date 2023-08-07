package main

import (
	"fmt"
	"github.com/sugrado/tama-server/internal/config"
	"log"
	"net/http"
	"time"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatal("cannot load config: ", err)
	}

	srv := &http.Server{
		MaxHeaderBytes: 10, // 10 MB
		Addr:           ":" + cfg.Server.Port,
		IdleTimeout:    time.Second * 60,
	}

	fmt.Println(fmt.Sprintf("Running on: '%s' \nEnvironment: '%s'", cfg.Server.Port, cfg.Server.Environment))
	if err := srv.ListenAndServe(); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
