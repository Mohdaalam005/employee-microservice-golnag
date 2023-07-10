// Package portal-service API.
//
// # Endpoints for employees
//
// employee-Service:
//
//	Schemes: http
//	Host: localhost:8080
//	BasePath: /
//	Version: 0.0.1
//
// Consumes:
// - application/json
//
// Produces:
// - application/json
//
// swagger:meta
package main

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/gorilla/mux"
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
	route := mux.NewRouter()
	fs := http.FileServer(http.Dir("./swagger-ui"))
	route.PathPrefix("/swagger-ui").
		Handler(http.StripPrefix("/swagger-ui", fs))

	ctx := context.Background()
	dao := database.NewEmployeeDao(*db)
	srv := service.NewEmployeeService(dao)
	end := endpoints.MakeEmployeeEndpoints(srv)

	transport.EmployeeHttpHandler(ctx, end, route)

	startServer(route)

}

func startServer(r *mux.Router) {
	serverPort := fmt.Sprintf(":%s", "8080")
	log.Println("starting server on ", serverPort)
	server := &http.Server{
		Addr:    serverPort,
		Handler: r,
	}
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal("failed to start server")
	}
}
