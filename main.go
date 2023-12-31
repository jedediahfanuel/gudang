package main

import (
	"log"
	"net/http"
	"os"

	"fmt"

	"gudang/database"
	"gudang/routes"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	db := database.GetDBConnection()
	// database.Migrate(db)

	r := mux.NewRouter()
	corsHandler := handlers.CORS(
		handlers.AllowedOrigins([]string{os.Getenv("FRONTEND_LINK")}),
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"}),
		handlers.AllowedHeaders([]string{"Content-Type"}),
	)(r)

	routes.RegisterRoutes(db, r)

	http.Handle("/", corsHandler)
	fmt.Println("Connected to port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
