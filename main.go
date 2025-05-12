package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	dotenv "github.com/dotenv-org/godotenvvault"
	"github.com/graphql-go/handler"
)

var (
	port                           int    = 8080
	host                           string = "0.0.0.0"
	carrier_graphql_endpoint       string
	organizations_graphql_endpoint string
)

func init() {
	if err := dotenv.Load(".env"); err != nil {
		log.Fatal(err)
		return
	}

	carrier_graphql_endpoint = os.Getenv("CARRIERS_GRAPHQL_ENDPOINT")
	organizations_graphql_endpoint = os.Getenv("ORGANIZATIONS_GRAPHQL_ENDPOINT")
	portStr := os.Getenv("PORT")
	if portStr != "" {
		parsedPort, err := strconv.Atoi(portStr)
		if err != nil {
			log.Fatal(err)
		}
		port = parsedPort
	}
	host = os.Getenv("host")
}

func main() {
	flag.IntVar(&port, "p", 8080, "Port to run application on")
	flag.StringVar(&host, "h", "0.0.0.0", "Host address to run application on")
	fmt.Println("Starting application...")

	fmt.Println("Setting up handlers and routes")
	schema, err := NewSchema()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Registered mutations and queries successfully")
	h := handler.New(&handler.Config{
		Schema:   &schema,
		Pretty:   true,
		GraphiQL: true,
	})

	server := http.NewServeMux()
	server.Handle("/graphql", h)
	server.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("{\"status\":\"ok\"}"))
		w.Header().Set("ContentType", "application/json")
	})
	fmt.Println("Routes and Handlers setup successfully")
	hostname := fmt.Sprintf("%v:%d", host, port)
	fmt.Printf("Start server at %s", hostname)
	http.ListenAndServe(hostname, server)
}
