package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/jackc/pgx/v4"
	"github.com/joho/godotenv"
)

var db *pgx.Conn

func main() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	connPort := os.Getenv("PORT")
	dbUser := os.Getenv("DB_USER")
	dbName := os.Getenv("DB_NAME")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	if connPort == "" || dbUser == "" || dbName == "" || dbHost == "" || dbPort == "" {
		log.Fatal("One or more required environment variables are missing")
	}

	connStr := fmt.Sprintf("postgres://%s@%s:%s/%s?sslmode=disable", dbUser, dbHost, dbPort, dbName)

	connConfig, err := pgx.ParseConfig(connStr)
	if err != nil {
		log.Fatal("Error parsing config: ", err)
	}

	db, err = pgx.ConnectConfig(context.Background(), connConfig)
	if err != nil {
		log.Fatal("Unable to connect to database: ", err)
	} else {
		log.Println("Successfully connected to the database!")
	}

	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	APIRouter := chi.NewRouter()

	APIRouter.Post("/users", createUser)                          // Yeni kullanıcı oluştur
	APIRouter.Get("/users", getUsers)                             // Tüm kullanıcıları listele
	APIRouter.Get("/users/{username}", getUserByUsername)         // Kullanıcıyı Username'e göre getir
	APIRouter.Put("/users/{username}", updateUser)                // Kullanıcıyı güncelle
	APIRouter.Delete("/users/{username}", deleteUser)             // Kullanıcıyı sil
	APIRouter.Post("/customers", createCustomer)                  // Yeni kullanıcı oluştur
	APIRouter.Get("/customers", getCustomers)                     // Tüm kullanıcıları listele
	APIRouter.Get("/customers/{username}", getCustomerByUsername) // Kullanıcıyı Username'e göre getir
	APIRouter.Put("/customers/{username}", updateCustomer)        // Kullanıcıyı güncelle
	APIRouter.Delete("/customers/{username}", deleteCustomer)     // Kullanıcıyı sil
	APIRouter.Post("/markets", createMarket)                      // Yeni kullanıcı oluştur
	APIRouter.Get("/markets", getMarkets)                         // Tüm kullanıcıları listele
	APIRouter.Get("/markets/{id}", getMarketByID)                 // Kullanıcıyı Username'e göre getir
	APIRouter.Put("/markets/{id}", updateMarket)                  // Kullanıcıyı güncelle
	APIRouter.Delete("/markets/{id}", deleteMarket)               // Kullanıcıyı sil
	APIRouter.Post("/login", approveLogin)                        // Kullanıcı girişini onaylama
	APIRouter.Post("/logout", logoutUser)                         // Kullanıcı çıkışını gerçekleştirme

	router.Mount("/api", APIRouter)

	srv := &http.Server{
		Handler: router,
		Addr:    ":" + connPort,
	}

	log.Printf("Server starting on port %v", connPort)

	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}

}
