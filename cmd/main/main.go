package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/sugrado/go-rest-api-template/internal/app"
	"github.com/sugrado/go-rest-api-template/internal/config"
	"github.com/sugrado/go-rest-api-template/internal/router"
	"github.com/sugrado/go-rest-api-template/internal/storage"
)

func main() {
	cfg := config.Load()

	db := storage.DBConn(cfg.Database)
	defer db.Close()

	d := storage.RegisterRepos(db)
	s := app.RegisterServices(d)

	srv := &http.Server{
		MaxHeaderBytes: 10, // 10 MB
		Addr:           ":" + cfg.Server.Port,
		IdleTimeout:    time.Second * 60,
		Handler:        router.New(s),
	}

	fmt.Println(fmt.Sprintf("Running on: '%s' \nEnvironment: '%s'", cfg.Server.Port, cfg.Server.Environment))
	if err := srv.ListenAndServe(); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
