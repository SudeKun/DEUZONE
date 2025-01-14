package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/joho/godotenv"
)

var db *pgxpool.Pool

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

	// Bağlantı havuzunun yapılandırılması
	config, err := pgxpool.ParseConfig(connStr)
	if err != nil {
		log.Fatal("Error parsing config: ", err)
	}

	// Bağlantı havuzu boyutunu artırın (MaxConns)
	config.MaxConns = 20 // Burada 20'yi, ihtiyacınıza göre değiştirebilirsiniz

	// Bağlantı havuzunu oluşturma
	db, err = pgxpool.ConnectConfig(context.Background(), config)
	if err != nil {
		log.Fatal("Unable to connect to database: ", err)
	} else {
		log.Println("Successfully connected to the database!")
	}
	defer db.Close()
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

	APIRouter.Post("/users", createUser)                  // Yeni kullanıcı oluştur
	APIRouter.Get("/users", getUsers)                     // Tüm kullanıcıları listele
	APIRouter.Get("/users/{username}", getUserByUsername) // Kullanıcıyı Username'e göre getir
	APIRouter.Put("/users/{username}", updateUser)        // Kullanıcıyı güncelle
	APIRouter.Delete("/users/{username}", deleteUser)     // Kullanıcıyı sil

	APIRouter.Post("/customers", createCustomer)                  // Yeni kullanıcı oluştur
	APIRouter.Get("/customers", getCustomers)                     // Tüm kullanıcıları listele
	APIRouter.Get("/customers/{username}", getCustomerByUsername) // Kullanıcıyı Username'e göre getir
	APIRouter.Put("/customers/{username}", updateCustomer)        // Kullanıcıyı güncelle
	APIRouter.Delete("/customers/{username}", deleteCustomer)     // Kullanıcıyı sil

	APIRouter.Post("/markets", createMarket)        // Yeni kullanıcı oluştur
	APIRouter.Get("/markets", getMarkets)           // Tüm kullanıcıları listele
	APIRouter.Get("/markets/{id}", getMarketByID)   // Kullanıcıyı Username'e göre getir
	APIRouter.Put("/markets/{id}", updateMarket)    // Kullanıcıyı güncelle
	APIRouter.Delete("/markets/{id}", deleteMarket) // Kullanıcıyı sil

	APIRouter.Post("/login", approveLogin) // Kullanıcı girişini onaylama
	APIRouter.Post("/logout", logoutUser)  // Kullanıcı çıkışını gerçekleştirme

	APIRouter.Post("/products", createProduct)
	APIRouter.Get("/products", getProducts)
	APIRouter.Post("/products/{id}", getProductByID)
	APIRouter.Put("/products/{id}", updateProduct)
	APIRouter.Delete("/products/{id}", deleteProduct)

	APIRouter.Post("/carts", createCart)        // Yeni sepet oluştur
	APIRouter.Get("/carts", getAllCarts)        // Tüm sepetleri listele
	APIRouter.Get("/carts/{id}", getCartByID)   // Sepeti ID ile getir
	APIRouter.Put("/carts/{id}", updateCart)    // Sepeti güncelle
	APIRouter.Delete("/carts/{id}", deleteCart) // Sepeti sil

	APIRouter.Post("/orders", createOrder)        // Yeni sipariş oluştur
	APIRouter.Get("/orders", getAllOrders)        // Tüm siparişleri listele
	APIRouter.Get("/orders/{id}", getOrderByID)   // Siparişi ID ile getir
	APIRouter.Put("/orders/{id}", updateOrder)    // Siparişi güncelle
	APIRouter.Delete("/orders/{id}", deleteOrder) // Siparişi sil

	APIRouter.Post("/prices", createPrice)        // Yeni fiyat oluştur
	APIRouter.Get("/prices", getPrices)           // Tüm fiyatları listele
	APIRouter.Get("/prices/{id}", getPriceByID)   // Fiyatı ID ile getir
	APIRouter.Put("/prices/{id}", updatePrice)    // Fiyatı güncelle
	APIRouter.Delete("/prices/{id}", deletePrice) // Fiyatı sil

	APIRouter.Post("/categories", createCategory)        // Yeni kategori oluştur
	APIRouter.Get("/categories", getCategories)          // Tüm kategorileri listele
	APIRouter.Get("/categories/{id}", getCategoryByID)   // Kategoriyi ID ile getir
	APIRouter.Put("/categories/{id}", updateCategory)    // Kategoriyi güncelle
	APIRouter.Delete("/categories/{id}", deleteCategory) // Kategoriyi sil

	APIRouter.Post("/colors", createColor)        // Yeni renk oluştur
	APIRouter.Get("/colors", getColors)           // Tüm renkleri listele
	APIRouter.Get("/colors/{id}", getColorByID)   // Rengi ID ile getir
	APIRouter.Put("/colors/{id}", updateColor)    // Rengi güncelle
	APIRouter.Delete("/colors/{id}", deleteColor) // Rengi sil

	APIRouter.Post("/cartitems", createCartItem)        // Yeni sepet öğesi oluştur
	APIRouter.Get("/cartitems", getCartItems)           // Tüm sepet öğelerini listele
	APIRouter.Get("/cartitems/{id}", getCartItemByID)   // Sepet öğesini ID ile getir
	APIRouter.Put("/cartitems/{id}", updateCartItem)    // Sepet öğesini güncelle
	APIRouter.Delete("/cartitems/{id}", deleteCartItem) // Sepet öğesini sil

	APIRouter.Post("/orderitems", createOrderItem)        // Yeni sipariş öğesi oluştur
	APIRouter.Get("/orderitems", getOrderItems)           // Tüm sipariş öğelerini listele
	APIRouter.Get("/orderitems/{id}", getOrderItemByID)   // Sipariş öğesini ID ile getir
	APIRouter.Put("/orderitems/{id}", updateOrderItem)    // Sipariş öğesini güncelle
	APIRouter.Delete("/orderitems/{id}", deleteOrderItem) // Sipariş öğesini sil

	APIRouter.Post("/marketcomments", createMarketComment)        // Yeni yorum oluştur
	APIRouter.Get("/marketcomments", getMarketComments)           // Tüm yorumları listele
	APIRouter.Get("/marketcomments/{id}", getMarketCommentByID)   // Yorum ID ile getir
	APIRouter.Put("/marketcomments/{id}", updateMarketComment)    // Yorum güncelle
	APIRouter.Delete("/marketcomments/{id}", deleteMarketComment) // Yorum sil

	APIRouter.Post("/productcomments", createProductComment)        // Yeni yorum oluştur
	APIRouter.Get("/productcomments", getProductComments)           // Tüm yorumları listele
	APIRouter.Get("/productcomments/{id}", getProductCommentByID)   // Yorum ID ile getir
	APIRouter.Put("/productcomments/{id}", updateProductComment)    // Yorum güncelle
	APIRouter.Delete("/productcomments/{id}", deleteProductComment) // Yorum sil

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
