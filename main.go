package main

import (
	"context"
	"database/sql"
	_ "github.com/lib/pq"
	"github.com/mohdaalam005/internal/database"
	"github.com/mohdaalam005/internal/endpoints"
	"github.com/mohdaalam005/internal/service"
	"github.com/mohdaalam005/internal/transport"
	"log"
	"net/http"
)

func main() {
	dsn := "dbname='go_lang' host='localhost' user='postgres' password='root' sslmode=disable"
	db, err := sql.Open("postgres", dsn)
	log.Println("application has been started")
	defer log.Println("application has been closed")
	if err != nil {
		log.Println(err)
	}
	ctx := context.Background()
	dao := database.NewEmployeeDao(*db)
	srv := service.NewEmployeeService(dao)
	end := endpoints.MakeEmployeeEndpoints(srv)

	handler := transport.CreateEmployeeHttpHandler(ctx, end)
	errs := make(chan error)

	//log.Info("listening on port", 8080)

	errs <- http.ListenAndServe(":8080", handler)
}
