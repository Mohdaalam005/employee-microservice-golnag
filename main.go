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
	"fmt"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	database2 "github.com/mohdaalam005/go-common/database"
	"github.com/mohdaalam005/internal/database"
	"github.com/mohdaalam005/internal/endpoints"
	"github.com/mohdaalam005/internal/service"
	"github.com/mohdaalam005/internal/transport"
	"log"
	"net/http"
)

func main() {

	dbConfig := database2.DbConfig{
		Host:   "localhost",
		Port:   5432,
		User:   "postgres",
		Pass:   "root",
		DbName: "go_lang",
	}
	dbConn, err := database2.InitDatabase(dbConfig)

	defer log.Println("application has been closed")
	if err != nil {
		log.Println(err)
	}
	route := mux.NewRouter()
	fs := http.FileServer(http.Dir("./swagger-ui"))
	route.PathPrefix("/swagger-ui").
		Handler(http.StripPrefix("/swagger-ui", fs))

	ctx := context.Background()
	dao := database.NewEmployeeDao(dbConn)
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
