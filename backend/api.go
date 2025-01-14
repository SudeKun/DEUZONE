package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/google/uuid"
)

type User struct {
	UserID   uuid.UUID `json:"user_id"`
	Email    string    `json:"email"`
	Password string    `json:"password"`
	Username string    `json:"username"`
	Status   bool      `json:"status"`
}

type Customer struct {
	CustomerID      uuid.UUID `json:"customer_id"`
	UserID          uuid.UUID `json:"user_id"`
	CustomerName    string    `json:"customer_name"`
	CustomerAddress string    `json:"customer_address"`
	CustomerPhone   string    `json:"customer_phone"`
}

type Market struct {
	MarketID      uuid.UUID `json:"market_id"`
	UserID        uuid.UUID `json:"user_id"`
	MarketName    string    `json:"market_name"`
	MarketAddress string    `json:"market_address"`
	MarketPhone   string    `json:"market_phone"`
}

type Product struct {
	ProductID    uuid.UUID `json:"product_id"`
	MarketID     uuid.UUID `json:"market_id"`
	CategoryID   uuid.UUID `json:"category_id"`
	ProductName  string    `json:"product_name"`
	ProductImage []byte    `json:"product_image"`
	Keyword      string    `json:"keyword"`
	Description  string    `json:"description"`
}

type Cart struct {
	CartID     uuid.UUID `json:"cart_id"`
	CustomerID uuid.UUID `json:"customer_id"`
}

type Order struct {
	OrderID    uuid.UUID `json:"order_id"`
	CustomerID uuid.UUID `json:"customer_id"`
	CartID     uuid.UUID `json:"cart_id"`
	Status     string    `json:"status"`
	DateOrder  time.Time `json:"date_order"` // Could be time.Time depending on your needs
	TotalPrice float64   `json:"total_price"`
}

type Price struct {
	PriceID   uuid.UUID `json:"price_id"`
	ProductID uuid.UUID `json:"product_id"`
	Price     float64   `json:"price"`
	Stock     bool      `json:"stock"`
}

type Category struct {
	CategoryID uuid.UUID `json:"category_id"`
	Name       string    `json:"name"`
}

type Color struct {
	ColorID uuid.UUID `json:"color_id"`
	Name    string    `json:"name"`
}

type CartItem struct {
	CartID   uuid.UUID `json:"cart_id"`
	PriceID  uuid.UUID `json:"price_id"`
	Quantity int       `json:"quantity"`
	Status   string    `json:"status"`
}

type OrderItem struct {
	OrderID         uuid.UUID `json:"order_id"`
	PriceID         uuid.UUID `json:"price_id"`
	Quantity        int       `json:"quantity"`
	PriceAtPurchase float64   `json:"price_at_purchase"`
}

type MarketComment struct {
	MCommentID uuid.UUID `json:"mcomment_id"`
	CustomerID uuid.UUID `json:"customer_id"`
	MarketID   uuid.UUID `json:"market_id"`
	Star       float64   `json:"star"`
	Comment    string    `json:"comment"`
	Date       time.Time `json:"date"`
}

type ProductComment struct {
	PCommentID uuid.UUID `json:"pcomment_id"`
	CustomerID uuid.UUID `json:"customer_id"`
	ProductID  uuid.UUID `json:"product_id"`
	Star       float64   `json:"star"`
	Comment    string    `json:"comment"`
	Date       time.Time `json:"date"`
}

func createUser(w http.ResponseWriter, r *http.Request) {
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	sql := `INSERT INTO "User" ("user_id", "email", "password", "username", "status") 
	VALUES ($1, $2, $3, $4, $5) RETURNING "user_id"`

	userID := uuid.New()
	err = db.QueryRow(context.Background(), sql,
		userID, user.Email, user.Password, user.Username, user.Status).Scan(&userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	user.UserID = userID
	respondWithJSON(w, 200, user)
}

func getUsers(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query(context.Background(),
		`SELECT "user_id", "email", "password", "username", "status"
		FROM "User"`)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var user User
		if err := rows.Scan(
			&user.UserID,
			&user.Email,
			&user.Password,
			&user.Username,
			&user.Status,
		); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		users = append(users, user)
	}

	if len(users) == 0 {
		respondWithError(w, 404, "No materials found!")
		return
	}

	respondWithJSON(w, 200, users)
}

func getUserByUsername(w http.ResponseWriter, r *http.Request) {
	var user User
	// URL parametresini almak için chi kullanıyoruz
	username := chi.URLParam(r, "username")
	sql := `SELECT "user_id", "email", "password", "username", "status"
		FROM "User" WHERE username = $1`
	err := db.QueryRow(context.Background(), sql, username).Scan(
		&user.UserID, &user.Email, &user.Password, &user.Username, &user.Status)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respondWithJSON(w, 200, user)
}

func updateUser(w http.ResponseWriter, r *http.Request) {
	var user User
	// URL parametresinden username'i alıyoruz
	username := chi.URLParam(r, "username")

	// Request'ten gelen JSON verisini `User` struct'ına parse ediyoruz
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid JSON body", http.StatusBadRequest)
		return
	}

	// SQL sorgusunu hazırlıyoruz (Kullanıcıyı güncellemek için)
	sql := `UPDATE "User" 
			SET "email" = $1, "password" = $2, "username" = $3, "status" = $4
			WHERE "username" = $5 RETURNING "user_id", "email", "password", "username", "status"`

	// Veritabanı sorgusunu çalıştırıyoruz
	err = db.QueryRow(context.Background(), sql, user.Email, user.Password, user.Username, user.Status, username).
		Scan(&user.UserID, &user.Email, &user.Password, &user.Username, &user.Status)
	if err != nil {
		http.Error(w, "Error updating user", http.StatusInternalServerError)
		log.Printf("Error updating user: %v", err)
		return
	}

	// Güncellenmiş kullanıcı bilgilerini JSON olarak döndürüyoruz
	respondWithJSON(w, 200, user)
}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	var user User
	username := chi.URLParam(r, "username")
	// SQL sorgusunu hazırlıyoruz (Kullanıcıyı güncellemek için)
	sql := `DELETE FROM "User" WHERE "username" = $1 RETURNING "user_id", "email", "password", "username", "status"`
	err := db.QueryRow(context.Background(), sql, username).Scan(&user.UserID, &user.Email,
		&user.Password, &user.Username, &user.Status)
	if err != nil {
		// Hata durumunda 500 hatası dönüyoruz
		if err.Error() == "no rows in result set" {
			// Eğer kullanıcı bulunamazsa
			http.Error(w, "User not found", http.StatusNotFound)
		} else {
			http.Error(w, "Error deleting user", http.StatusInternalServerError)
		}
		log.Printf("Error deleting user: %v", err)
		return
	}

	// Güncellenmiş kullanıcı bilgilerini JSON olarak döndürüyoruz
	respondWithJSON(w, 200, user)
}

func createCustomer(w http.ResponseWriter, r *http.Request) {
	var customer Customer
	err := json.NewDecoder(r.Body).Decode(&customer)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	sql := `INSERT INTO "Customer" ("customer_id", "user_id", "customer_name", "customer_address", "customer_phone")
	VALUES ($1, $2, $3, $4, $5) RETURNING "customer_id"`

	customerID := uuid.New()
	err = db.QueryRow(context.Background(), sql,
		customerID, customer.UserID, customer.CustomerName, customer.CustomerAddress, customer.CustomerPhone).
		Scan(&customerID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	customer.CustomerID = customerID
	respondWithJSON(w, 200, customer)
}

func getCustomers(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query(context.Background(),
		`SELECT "customer_id", "user_id", "customer_name", "customer_address", "customer_phone"
		FROM "Customer"`)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var customers []Customer
	for rows.Next() {
		var customer Customer
		if err := rows.Scan(
			&customer.CustomerID,
			&customer.UserID,
			&customer.CustomerName,
			&customer.CustomerAddress,
			&customer.CustomerPhone,
		); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		customers = append(customers, customer)
	}

	if len(customers) == 0 {
		respondWithError(w, 404, "No customers found!")
		return
	}

	respondWithJSON(w, 200, customers)
}

func getCustomerByUsername(w http.ResponseWriter, r *http.Request) {
	var customer Customer
	// URL parametresini almak için chi kullanıyoruz
	username := chi.URLParam(r, "username")

	// Kullanıcıya ait customer verisini çekmek için SQL sorgusu
	sql := `SELECT c."customer_id", c."user_id", c."customer_name", c."customer_address", c."customer_phone"
			FROM "Customer" c
			JOIN "User" u ON c."user_id" = u."user_id"
			WHERE u."username" = $1`

	// SQL sorgusunu çalıştırıyoruz
	err := db.QueryRow(context.Background(), sql, username).Scan(
		&customer.CustomerID, &customer.UserID, &customer.CustomerName, &customer.CustomerAddress, &customer.CustomerPhone)
	if err != nil {
		if err.Error() == "no rows in result set" {
			http.Error(w, "Customer not found", http.StatusNotFound)
		} else {
			http.Error(w, "Error fetching customer", http.StatusInternalServerError)
		}
		log.Printf("Error fetching customer: %v", err)
		return
	}

	// JSON olarak kullanıcıyı döndürüyoruz
	respondWithJSON(w, 200, customer)
}

func updateCustomer(w http.ResponseWriter, r *http.Request) {
	var customer Customer
	username := chi.URLParam(r, "username")

	// Gelen JSON verisini `Customer` struct'ına parse ediyoruz
	err := json.NewDecoder(r.Body).Decode(&customer)
	if err != nil {
		http.Error(w, "Invalid JSON body", http.StatusBadRequest)
		return
	}

	// SQL sorgusunu hazırlıyoruz
	sql := `UPDATE "Customer" 
			SET "customer_name" = $1, "customer_address" = $2, "customer_phone" = $3
			FROM "User" u
			WHERE u."username" = $4 AND "Customer"."user_id" = u."user_id"
			RETURNING "customer_id", "user_id", "customer_name", "customer_address", "customer_phone"`

	// Veritabanı sorgusunu çalıştırıyoruz
	err = db.QueryRow(context.Background(), sql, customer.CustomerName, customer.CustomerAddress, customer.CustomerPhone, username).
		Scan(&customer.CustomerID, &customer.UserID, &customer.CustomerName, &customer.CustomerAddress, &customer.CustomerPhone)
	if err != nil {
		http.Error(w, "Error updating customer", http.StatusInternalServerError)
		log.Printf("Error updating customer: %v", err)
		return
	}

	// Güncellenmiş kullanıcı bilgilerini JSON olarak döndürüyoruz
	respondWithJSON(w, 200, customer)
}

func deleteCustomer(w http.ResponseWriter, r *http.Request) {
	var customer Customer
	username := chi.URLParam(r, "username")

	// SQL sorgusunu hazırlıyoruz
	sql := `DELETE FROM "Customer" 
			WHERE "user_id" = (SELECT "user_id" FROM "User" WHERE "username" = $1) 
			RETURNING "customer_id", "user_id", "customer_name", "customer_address", "customer_phone"`

	// Veritabanı sorgusunu çalıştırıyoruz
	err := db.QueryRow(context.Background(), sql, username).Scan(&customer.CustomerID, &customer.UserID,
		&customer.CustomerName, &customer.CustomerAddress, &customer.CustomerPhone)
	if err != nil {
		if err.Error() == "no rows in result set" {
			http.Error(w, "Customer not found", http.StatusNotFound)
		} else {
			http.Error(w, "Error deleting customer", http.StatusInternalServerError)
		}
		log.Printf("Error deleting customer: %v", err)
		return
	}

	// Silinen kullanıcı bilgilerini JSON olarak döndürüyoruz
	respondWithJSON(w, 200, customer)
}

func createMarket(w http.ResponseWriter, r *http.Request) {
	var market Market
	err := json.NewDecoder(r.Body).Decode(&market)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	sql := `INSERT INTO "Market" ("market_id", "user_id", "market_name", "market_address", "market_phone")
	VALUES ($1, $2, $3, $4, $5) RETURNING "market_id"`

	marketID := uuid.New()
	err = db.QueryRow(context.Background(), sql,
		marketID, market.UserID, market.MarketName, market.MarketAddress, market.MarketPhone).
		Scan(&marketID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	market.MarketID = marketID
	respondWithJSON(w, 200, market)
}

func getMarkets(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query(context.Background(),
		`SELECT "market_id", "user_id", "market_name", "market_address", "market_phone"
		FROM "Market"`)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var markets []Market
	for rows.Next() {
		var market Market
		if err := rows.Scan(
			&market.MarketID,
			&market.UserID,
			&market.MarketName,
			&market.MarketAddress,
			&market.MarketPhone,
		); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		markets = append(markets, market)
	}

	if len(markets) == 0 {
		respondWithError(w, 404, "No markets found!")
		return
	}

	respondWithJSON(w, 200, markets)
}

func getMarketByID(w http.ResponseWriter, r *http.Request) {
	var market Market
	// URL parametresini almak için chi kullanıyoruz
	marketID := chi.URLParam(r, "id")

	// Market verisini almak için SQL sorgusu
	sql := `SELECT "market_id", "user_id", "market_name", "market_address", "market_phone"
			FROM "Market" WHERE "market_id" = $1`

	// SQL sorgusunu çalıştırıyoruz
	err := db.QueryRow(context.Background(), sql, marketID).Scan(
		&market.MarketID, &market.UserID, &market.MarketName, &market.MarketAddress, &market.MarketPhone)
	if err != nil {
		if err.Error() == "no rows in result set" {
			http.Error(w, "Market not found", http.StatusNotFound)
		} else {
			http.Error(w, "Error fetching market", http.StatusInternalServerError)
		}
		log.Printf("Error fetching market: %v", err)
		return
	}

	// JSON olarak market bilgisini döndürüyoruz
	respondWithJSON(w, 200, market)
}

func updateMarket(w http.ResponseWriter, r *http.Request) {
	var market Market
	marketID := chi.URLParam(r, "id")

	// Gelen JSON verisini `Market` struct'ına parse ediyoruz
	err := json.NewDecoder(r.Body).Decode(&market)
	if err != nil {
		http.Error(w, "Invalid JSON body", http.StatusBadRequest)
		return
	}

	// SQL sorgusunu hazırlıyoruz
	sql := `UPDATE "Market"
			SET "market_name" = $1, "market_address" = $2, "market_phone" = $3
			WHERE "market_id" = $4 RETURNING "market_id", "user_id", "market_name", "market_address", "market_phone"`

	// Veritabanı sorgusunu çalıştırıyoruz
	err = db.QueryRow(context.Background(), sql, market.MarketName, market.MarketAddress, market.MarketPhone, marketID).
		Scan(&market.MarketID, &market.UserID, &market.MarketName, &market.MarketAddress, &market.MarketPhone)
	if err != nil {
		http.Error(w, "Error updating market", http.StatusInternalServerError)
		log.Printf("Error updating market: %v", err)
		return
	}

	// Güncellenmiş market bilgilerini JSON olarak döndürüyoruz
	respondWithJSON(w, 200, market)
}

func deleteMarket(w http.ResponseWriter, r *http.Request) {
	var market Market
	marketID := chi.URLParam(r, "id")

	// SQL sorgusunu hazırlıyoruz
	sql := `DELETE FROM "Market" WHERE "market_id" = $1 RETURNING "market_id", "user_id", "market_name", "market_address", "market_phone"`

	// Veritabanı sorgusunu çalıştırıyoruz
	err := db.QueryRow(context.Background(), sql, marketID).Scan(&market.MarketID, &market.UserID,
		&market.MarketName, &market.MarketAddress, &market.MarketPhone)
	if err != nil {
		if err.Error() == "no rows in result set" {
			http.Error(w, "Market not found", http.StatusNotFound)
		} else {
			http.Error(w, "Error deleting market", http.StatusInternalServerError)
		}
		log.Printf("Error deleting market: %v", err)
		return
	}

	// Silinen market bilgisini JSON olarak döndürüyoruz
	respondWithJSON(w, 200, market)
}

func approveLoginHelper(field string, w http.ResponseWriter) string {
	// İlk önce username ile sorgulama yapıyoruz
	sql := `SELECT "user_id", "email", "password", "username", "status"
			FROM "User" WHERE "username" = $1`
	var user User
	err := db.QueryRow(context.Background(), sql, field).Scan(&user.UserID, &user.Email, &user.Password, &user.Username, &user.Status)
	if err == nil {
		// Eğer username ile sonuç bulduysak, "username" döndür
		return "username"
	}

	// Eğer username ile sonuç bulamadıysak, email ile sorgulama yapıyoruz
	sql = `SELECT "user_id", "email", "password", "username", "status"
			FROM "User" WHERE "email" = $1`
	err = db.QueryRow(context.Background(), sql, field).Scan(&user.UserID, &user.Email, &user.Password, &user.Username, &user.Status)
	if err == nil {
		// Eğer email ile sonuç bulduysak, "email" döndür
		return "email"
	}

	// Eğer her iki sorgu da başarısız olduysa, hata döndürüyoruz
	respondWithError(w, 403, "This field is invalid !!")
	return ""
}

func approveLogin(w http.ResponseWriter, r *http.Request) {
	var user User
	var sql string
	var updateSQL string
	var jsonData map[string]interface{}
	err := json.NewDecoder(r.Body).Decode(&jsonData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	field, ok := jsonData["username_or_email_field"].(string)
	if !ok {
		http.Error(w, "Please enter an username or email !!", http.StatusBadRequest)
		return
	}
	password, ok := jsonData["password"].(string)
	if !ok || password == "" {
		http.Error(w, "Password is required", http.StatusBadRequest)
		return
	}
	username_or_email := approveLoginHelper(field, w)
	switch username_or_email {
	case "username":
		// Eğer field 'username' ise, username ile sorgulama yapıyoruz
		sql = `SELECT "user_id", "email", "password", "username", "status" 
			FROM "User" WHERE username = $1 AND password = $2`
		err = db.QueryRow(context.Background(), sql, field, password).Scan(
			&user.UserID, &user.Email, &user.Password, &user.Username, &user.Status)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	case "email":
		// Eğer field 'email' ise, email ile sorgulama yapıyoruz
		sql = `SELECT "user_id", "email", "password", "username", "status" 
			FROM "User" WHERE email = $1`
		rows, err := db.Query(context.Background(), sql, field)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer rows.Close()

		// Şifreyi kontrol et
		var foundUser bool
		for rows.Next() {
			var userInDb User
			if err := rows.Scan(&userInDb.UserID, &userInDb.Email, &userInDb.Password, &userInDb.Username, &userInDb.Status); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			// Şifreyi kontrol et
			if userInDb.Password == password {
				user = userInDb
				foundUser = true
				break
			}
		}

		if !foundUser {
			http.Error(w, "Invalid password", http.StatusUnauthorized)
			return
		}
	}

	if user.Status {
		respondWithError(w, 403, "This user is already logged in !!")
		return
	}
	updateSQL = `UPDATE "User" SET "status" = true WHERE username = $1`
	_, err = db.Exec(context.Background(), updateSQL, user.Username)
	if err != nil {
		http.Error(w, "Error updating user status", http.StatusInternalServerError)
		log.Printf("Error updating user status: %v", err)
		return
	}
	user.Status = true
	respondWithJSON(w, 200, user)
}

func logoutUser(w http.ResponseWriter, r *http.Request) {
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	sql := `SELECT "user_id", "email", "password", "username", "status"
		FROM "User" WHERE username = $1`
	err = db.QueryRow(context.Background(), sql, user.Username).Scan(
		&user.UserID, &user.Email, &user.Password, &user.Username, &user.Status)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if !user.Status {
		respondWithError(w, 403, "This user is already logged out !!")
		return
	}
	updateSQL := `UPDATE "User" SET "status" = false WHERE username = $1`
	_, err = db.Exec(context.Background(), updateSQL, user.Username)
	if err != nil {
		http.Error(w, "Error updating status", http.StatusInternalServerError)
		log.Printf("Error updating user status: %v", err)
		return
	}
	user.Status = false
	respondWithJSON(w, 200, user)
}

func createProduct(w http.ResponseWriter, r *http.Request) {
	var product Product
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	sql := `INSERT INTO "Products" ("product_id", "market_id", "category_id", "product_name", 
		"product_image", "keyword", "description") VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING "product_id"`

	productID := uuid.New()
	err = db.QueryRow(context.Background(), sql,
		productID, product.MarketID, product.CategoryID, product.ProductName,
		product.ProductImage, product.Keyword, product.Description).Scan(&productID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	product.ProductID = productID
	respondWithJSON(w, 200, product)
}

func getProducts(w http.ResponseWriter, r *http.Request) {
	sql := `SELECT "product_id", "market_id", "category_id", "product_name", "product_image", "keyword", "description" FROM "Products"`

	rows, err := db.Query(context.Background(), sql)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var products []Product
	for rows.Next() {
		var product Product
		err := rows.Scan(&product.ProductID, &product.MarketID, &product.CategoryID, &product.ProductName, &product.ProductImage, &product.Keyword, &product.Description)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		products = append(products, product)
	}

	if err := rows.Err(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respondWithJSON(w, 200, products)
}

func getProductByID(w http.ResponseWriter, r *http.Request) {
	productID := chi.URLParam(r, "id") // URL parametresinden ID'yi alıyoruz

	sql := `SELECT "product_id", "market_id", "category_id", "product_name", "product_image", "keyword", "description" 
			FROM "Products" WHERE "product_id" = $1`

	var product Product
	err := db.QueryRow(context.Background(), sql, productID).Scan(&product.ProductID, &product.MarketID, &product.CategoryID, &product.ProductName, &product.ProductImage, &product.Keyword, &product.Description)
	if err != nil {
		if err.Error() == "no rows in result set" {
			http.Error(w, "Product not found", http.StatusNotFound)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	respondWithJSON(w, 200, product)
}

func updateProduct(w http.ResponseWriter, r *http.Request) {
	// URL parametresinden product_id'yi alıyoruz
	productID := chi.URLParam(r, "id")

	var product Product
	// Request'ten gelen JSON verisini `Product` struct'ına parse ediyoruz
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		http.Error(w, "Invalid JSON body", http.StatusBadRequest)
		return
	}

	// SQL sorgusunu hazırlıyoruz (Ürünü güncellemek için)
	sql := `UPDATE "Products" 
			SET "market_id" = $1, "category_id" = $2, "product_name" = $3, 
				"product_image" = $4, "keyword" = $5, "description" = $6
			WHERE "product_id" = $7 
			RETURNING "product_id", "market_id", "category_id", "product_name", 
				"product_image", "keyword", "description"`

	// Veritabanı sorgusunu çalıştırıyoruz
	err = db.QueryRow(context.Background(), sql,
		product.MarketID, product.CategoryID, product.ProductName,
		product.ProductImage, product.Keyword, product.Description, productID).
		Scan(&product.ProductID, &product.MarketID, &product.CategoryID, &product.ProductName,
			&product.ProductImage, &product.Keyword, &product.Description)

	if err != nil {
		http.Error(w, "Error updating product", http.StatusInternalServerError)
		log.Printf("Error updating product: %v", err)
		return
	}

	// Güncellenmiş ürün bilgilerini JSON olarak döndürüyoruz
	respondWithJSON(w, 200, product)
}

func deleteProduct(w http.ResponseWriter, r *http.Request) {
	// URL parametresinden product_id'yi alıyoruz
	productID := chi.URLParam(r, "id")

	var product Product
	// SQL sorgusunu hazırlıyoruz (Ürünü silmek için)
	sql := `DELETE FROM "Products" WHERE "product_id" = $1 
			RETURNING "product_id", "market_id", "category_id", "product_name", 
				"product_image", "keyword", "description"`

	// Veritabanı sorgusunu çalıştırıyoruz
	err := db.QueryRow(context.Background(), sql, productID).
		Scan(&product.ProductID, &product.MarketID, &product.CategoryID, &product.ProductName,
			&product.ProductImage, &product.Keyword, &product.Description)

	if err != nil {
		// Hata durumunda 404 ya da 500 hatası dönüyoruz
		if err.Error() == "no rows in result set" {
			// Eğer ürün bulunamazsa
			http.Error(w, "Product not found", http.StatusNotFound)
		} else {
			http.Error(w, "Error deleting product", http.StatusInternalServerError)
		}
		log.Printf("Error deleting product: %v", err)
		return
	}

	// Silinen ürün bilgilerini JSON olarak döndürüyoruz
	respondWithJSON(w, 200, product)
}

func createCart(w http.ResponseWriter, r *http.Request) {
	var cart Cart
	err := json.NewDecoder(r.Body).Decode(&cart)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// SQL sorgusu ile yeni sepeti ekliyoruz
	sql := `INSERT INTO "Cart" ("cart_id", "customer_id") VALUES ($1, $2) RETURNING "cart_id"`

	// Yeni UUID oluşturuyoruz
	cartID := uuid.New()

	// Veritabanı sorgusunu çalıştırıyoruz
	err = db.QueryRow(context.Background(), sql, cartID, cart.CustomerID).Scan(&cart.CartID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Sepet bilgilerini JSON olarak döndürüyoruz
	cart.CartID = cartID
	respondWithJSON(w, 200, cart)
}

func getAllCarts(w http.ResponseWriter, r *http.Request) {
	sql := `SELECT "cart_id", "customer_id" FROM "Cart"`

	rows, err := db.Query(context.Background(), sql)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var carts []Cart
	for rows.Next() {
		var cart Cart
		err := rows.Scan(&cart.CartID, &cart.CustomerID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		carts = append(carts, cart)
	}

	if err := rows.Err(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respondWithJSON(w, 200, carts)
}

func getCartByID(w http.ResponseWriter, r *http.Request) {
	cartID := chi.URLParam(r, "id") // URL parametresinden `id` alıyoruz

	sql := `SELECT "cart_id", "customer_id" FROM "Cart" WHERE "cart_id" = $1`

	var cart Cart
	err := db.QueryRow(context.Background(), sql, cartID).Scan(&cart.CartID, &cart.CustomerID)
	if err != nil {
		if err.Error() == "no rows in result set" {
			http.Error(w, "Cart not found", http.StatusNotFound)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	respondWithJSON(w, 200, cart)
}

func updateCart(w http.ResponseWriter, r *http.Request) {
	cartID := chi.URLParam(r, "id") // URL parametresinden `cart_id` alıyoruz

	var cart Cart
	err := json.NewDecoder(r.Body).Decode(&cart)
	if err != nil {
		http.Error(w, "Invalid JSON body", http.StatusBadRequest)
		return
	}

	// SQL sorgusu ile sepeti güncelliyoruz
	sql := `UPDATE "Cart" 
			SET "customer_id" = $1
			WHERE "cart_id" = $2
			RETURNING "cart_id", "customer_id"`

	// Veritabanı sorgusunu çalıştırıyoruz
	err = db.QueryRow(context.Background(), sql, cart.CustomerID, cartID).
		Scan(&cart.CartID, &cart.CustomerID)
	if err != nil {
		http.Error(w, "Error updating cart", http.StatusInternalServerError)
		log.Printf("Error updating cart: %v", err)
		return
	}

	// Güncellenmiş sepet bilgisini JSON olarak döndürüyoruz
	respondWithJSON(w, 200, cart)
}

func deleteCart(w http.ResponseWriter, r *http.Request) {
	cartID := chi.URLParam(r, "id") // URL parametresinden `cart_id` alıyoruz

	var cart Cart
	// SQL sorgusu ile sepeti siliyoruz
	sql := `DELETE FROM "Cart" WHERE "cart_id" = $1 
			RETURNING "cart_id", "customer_id"`

	// Veritabanı sorgusunu çalıştırıyoruz
	err := db.QueryRow(context.Background(), sql, cartID).
		Scan(&cart.CartID, &cart.CustomerID)
	if err != nil {
		// Hata durumunda 404 ya da 500 hatası dönüyoruz
		if err.Error() == "no rows in result set" {
			// Eğer sepet bulunamazsa
			http.Error(w, "Cart not found", http.StatusNotFound)
		} else {
			http.Error(w, "Error deleting cart", http.StatusInternalServerError)
		}
		log.Printf("Error deleting cart: %v", err)
		return
	}

	// Silinen sepet bilgisini JSON olarak döndürüyoruz
	respondWithJSON(w, 200, cart)
}

func createOrder(w http.ResponseWriter, r *http.Request) {
	var order Order
	err := json.NewDecoder(r.Body).Decode(&order)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// SQL sorgusu ile yeni siparişi ekliyoruz
	sql := `INSERT INTO "Orders" ("order_id", "customer_id", "cart_id", "status", "date_order", "total_price")
			VALUES ($1, $2, $3, $4, $5, $6) RETURNING "order_id"`

	// Yeni UUID oluşturuyoruz
	orderID := uuid.New()
	order.DateOrder = time.Now() // Sipariş oluşturulma tarihi, şu anki zaman
	err = db.QueryRow(context.Background(), sql, orderID, order.CustomerID, order.CartID, order.Status, order.DateOrder, order.TotalPrice).
		Scan(&order.OrderID)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Sipariş bilgilerini JSON olarak döndürüyoruz
	order.OrderID = orderID
	respondWithJSON(w, 200, order)
}

func getAllOrders(w http.ResponseWriter, r *http.Request) {
	sql := `SELECT "order_id", "customer_id", "cart_id", "status", "date_order", "total_price" FROM "Orders"`

	rows, err := db.Query(context.Background(), sql)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var orders []Order
	for rows.Next() {
		var order Order
		err := rows.Scan(&order.OrderID, &order.CustomerID, &order.CartID, &order.Status, &order.DateOrder, &order.TotalPrice)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		orders = append(orders, order)
	}

	if err := rows.Err(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respondWithJSON(w, 200, orders)
}

func getOrderByID(w http.ResponseWriter, r *http.Request) {
	orderID := chi.URLParam(r, "id") // URL parametresinden `id` alıyoruz

	sql := `SELECT "order_id", "customer_id", "cart_id", "status", "date_order", "total_price" 
			FROM "Orders" WHERE "order_id" = $1`

	var order Order
	err := db.QueryRow(context.Background(), sql, orderID).Scan(&order.OrderID, &order.CustomerID, &order.CartID, &order.Status, &order.DateOrder, &order.TotalPrice)
	if err != nil {
		if err.Error() == "no rows in result set" {
			http.Error(w, "Order not found", http.StatusNotFound)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	respondWithJSON(w, 200, order)
}

func updateOrder(w http.ResponseWriter, r *http.Request) {
	orderID := chi.URLParam(r, "id") // URL parametresinden `order_id` alıyoruz

	var order Order
	// Request'ten gelen JSON verisini `Order` struct'ına parse ediyoruz
	err := json.NewDecoder(r.Body).Decode(&order)
	if err != nil {
		http.Error(w, "Invalid JSON body", http.StatusBadRequest)
		return
	}

	// SQL sorgusunu hazırlıyoruz (Siparişi güncellemek için)
	sql := `UPDATE "Orders" 
			SET "customer_id" = $1, "cart_id" = $2, "status" = $3, "date_order" = $4, "total_price" = $5
			WHERE "order_id" = $6 
			RETURNING "order_id", "customer_id", "cart_id", "status", "date_order", "total_price"`

	// Veritabanı sorgusunu çalıştırıyoruz
	err = db.QueryRow(context.Background(), sql,
		order.CustomerID, order.CartID, order.Status, order.DateOrder, order.TotalPrice, orderID).
		Scan(&order.OrderID, &order.CustomerID, &order.CartID, &order.Status, &order.DateOrder, &order.TotalPrice)

	if err != nil {
		http.Error(w, "Error updating order", http.StatusInternalServerError)
		log.Printf("Error updating order: %v", err)
		return
	}

	// Güncellenmiş sipariş bilgisini JSON olarak döndürüyoruz
	respondWithJSON(w, 200, order)
}

func deleteOrder(w http.ResponseWriter, r *http.Request) {
	orderID := chi.URLParam(r, "id") // URL parametresinden `order_id` alıyoruz

	var order Order
	// SQL sorgusu ile siparişi siliyoruz
	sql := `DELETE FROM "Orders" WHERE "order_id" = $1 
			RETURNING "order_id", "customer_id", "cart_id", "status", "date_order", "total_price"`

	// Veritabanı sorgusunu çalıştırıyoruz
	err := db.QueryRow(context.Background(), sql, orderID).
		Scan(&order.OrderID, &order.CustomerID, &order.CartID, &order.Status, &order.DateOrder, &order.TotalPrice)

	if err != nil {
		// Hata durumunda 404 ya da 500 hatası dönüyoruz
		if err.Error() == "no rows in result set" {
			// Eğer sipariş bulunamazsa
			http.Error(w, "Order not found", http.StatusNotFound)
		} else {
			http.Error(w, "Error deleting order", http.StatusInternalServerError)
		}
		log.Printf("Error deleting order: %v", err)
		return
	}

	// Silinen sipariş bilgisini JSON olarak döndürüyoruz
	respondWithJSON(w, 200, order)
}

func createPrice(w http.ResponseWriter, r *http.Request) {
	var price Price
	err := json.NewDecoder(r.Body).Decode(&price)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// SQL sorgusu ile yeni fiyatı ekliyoruz
	sql := `INSERT INTO "Prices" ("price_id", "product_id", "price", "stock") 
			VALUES ($1, $2, $3, $4) RETURNING "price_id"`

	// Yeni UUID oluşturuyoruz
	priceID := uuid.New()
	err = db.QueryRow(context.Background(), sql, priceID, price.ProductID, price.Price, price.Stock).
		Scan(&priceID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	price.PriceID = priceID
	respondWithJSON(w, 200, price)
}

func getPrices(w http.ResponseWriter, r *http.Request) {
	sql := `SELECT "price_id", "product_id", "price", "stock" FROM "Prices"`

	rows, err := db.Query(context.Background(), sql)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var prices []Price
	for rows.Next() {
		var price Price
		err := rows.Scan(&price.PriceID, &price.ProductID, &price.Price, &price.Stock)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		prices = append(prices, price)
	}

	if err := rows.Err(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respondWithJSON(w, 200, prices)
}

func getPriceByID(w http.ResponseWriter, r *http.Request) {
	priceID := chi.URLParam(r, "id") // URL parametresinden `id` alıyoruz

	sql := `SELECT "price_id", "product_id", "price", "stock" 
			FROM "Prices" WHERE "price_id" = $1`

	var price Price
	err := db.QueryRow(context.Background(), sql, priceID).Scan(&price.PriceID, &price.ProductID, &price.Price, &price.Stock)
	if err != nil {
		if err.Error() == "no rows in result set" {
			http.Error(w, "Price not found", http.StatusNotFound)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	respondWithJSON(w, 200, price)
}

func updatePrice(w http.ResponseWriter, r *http.Request) {
	// URL parametresinden `price_id`yi alıyoruz
	priceID := chi.URLParam(r, "id")

	var price Price
	// Request'ten gelen JSON verisini `Price` struct'ına parse ediyoruz
	err := json.NewDecoder(r.Body).Decode(&price)
	if err != nil {
		http.Error(w, "Invalid JSON body", http.StatusBadRequest)
		return
	}

	// SQL sorgusunu hazırlıyoruz (Fiyatı güncellemek için)
	sql := `UPDATE "Prices" 
			SET "product_id" = $1, "price" = $2, "stock" = $3
			WHERE "price_id" = $4 
			RETURNING "price_id", "product_id", "price", "stock"`

	// Veritabanı sorgusunu çalıştırıyoruz
	err = db.QueryRow(context.Background(), sql,
		price.ProductID, price.Price, price.Stock, priceID).
		Scan(&price.PriceID, &price.ProductID, &price.Price, &price.Stock)

	if err != nil {
		http.Error(w, "Error updating price", http.StatusInternalServerError)
		log.Printf("Error updating price: %v", err)
		return
	}

	// Güncellenmiş fiyat bilgisini JSON olarak döndürüyoruz
	respondWithJSON(w, 200, price)
}

func deletePrice(w http.ResponseWriter, r *http.Request) {
	// URL parametresinden `price_id`yi alıyoruz
	priceID := chi.URLParam(r, "id")

	var price Price
	// SQL sorgusu ile fiyatı siliyoruz
	sql := `DELETE FROM "Prices" WHERE "price_id" = $1 
			RETURNING "price_id", "product_id", "price", "stock"`

	// Veritabanı sorgusunu çalıştırıyoruz
	err := db.QueryRow(context.Background(), sql, priceID).
		Scan(&price.PriceID, &price.ProductID, &price.Price, &price.Stock)

	if err != nil {
		// Hata durumunda 404 ya da 500 hatası dönüyoruz
		if err.Error() == "no rows in result set" {
			// Eğer fiyat bulunamazsa
			http.Error(w, "Price not found", http.StatusNotFound)
		} else {
			http.Error(w, "Error deleting price", http.StatusInternalServerError)
		}
		log.Printf("Error deleting price: %v", err)
		return
	}

	// Silinen fiyat bilgisini JSON olarak döndürüyoruz
	respondWithJSON(w, 200, price)
}

func createCategory(w http.ResponseWriter, r *http.Request) {
	var category Category
	err := json.NewDecoder(r.Body).Decode(&category)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// SQL sorgusu ile yeni kategoriyi ekliyoruz
	sql := `INSERT INTO "Categories" ("category_id", "name") 
			VALUES ($1, $2) RETURNING "category_id"`

	// Yeni UUID oluşturuyoruz
	categoryID := uuid.New()
	err = db.QueryRow(context.Background(), sql, categoryID, category.Name).
		Scan(&categoryID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	category.CategoryID = categoryID
	respondWithJSON(w, 200, category)
}

func getCategories(w http.ResponseWriter, r *http.Request) {
	sql := `SELECT "category_id", "name" FROM "Categories"`

	rows, err := db.Query(context.Background(), sql)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var categories []Category
	for rows.Next() {
		var category Category
		err := rows.Scan(&category.CategoryID, &category.Name)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		categories = append(categories, category)
	}

	if err := rows.Err(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respondWithJSON(w, 200, categories)
}

func getCategoryByID(w http.ResponseWriter, r *http.Request) {
	categoryID := chi.URLParam(r, "id") // URL parametresinden ID'yi alıyoruz

	sql := `SELECT "category_id", "name" 
			FROM "Categories" WHERE "category_id" = $1`

	var category Category
	err := db.QueryRow(context.Background(), sql, categoryID).Scan(&category.CategoryID, &category.Name)
	if err != nil {
		if err.Error() == "no rows in result set" {
			http.Error(w, "Category not found", http.StatusNotFound)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	respondWithJSON(w, 200, category)
}

func updateCategory(w http.ResponseWriter, r *http.Request) {
	// URL parametresinden category_id'yi alıyoruz
	categoryID := chi.URLParam(r, "id")

	var category Category
	// Request'ten gelen JSON verisini Category struct'ına parse ediyoruz
	err := json.NewDecoder(r.Body).Decode(&category)
	if err != nil {
		http.Error(w, "Invalid JSON body", http.StatusBadRequest)
		return
	}

	// SQL sorgusunu hazırlıyoruz (Kategoriyi güncellemek için)
	sql := `UPDATE "Categories" 
			SET "name" = $1
			WHERE "category_id" = $2 
			RETURNING "category_id", "name"`

	// Veritabanı sorgusunu çalıştırıyoruz
	err = db.QueryRow(context.Background(), sql, category.Name, categoryID).
		Scan(&category.CategoryID, &category.Name)

	if err != nil {
		http.Error(w, "Error updating category", http.StatusInternalServerError)
		log.Printf("Error updating category: %v", err)
		return
	}

	// Güncellenmiş kategori bilgisini JSON olarak döndürüyoruz
	respondWithJSON(w, 200, category)
}

func deleteCategory(w http.ResponseWriter, r *http.Request) {
	// URL parametresinden category_id'yi alıyoruz
	categoryID := chi.URLParam(r, "id")

	var category Category
	// SQL sorgusu ile kategoriyi siliyoruz
	sql := `DELETE FROM "Categories" WHERE "category_id" = $1 
			RETURNING "category_id", "name"`

	// Veritabanı sorgusunu çalıştırıyoruz
	err := db.QueryRow(context.Background(), sql, categoryID).
		Scan(&category.CategoryID, &category.Name)

	if err != nil {
		// Hata durumunda 404 ya da 500 hatası dönüyoruz
		if err.Error() == "no rows in result set" {
			// Eğer kategori bulunamazsa
			http.Error(w, "Category not found", http.StatusNotFound)
		} else {
			http.Error(w, "Error deleting category", http.StatusInternalServerError)
		}
		log.Printf("Error deleting category: %v", err)
		return
	}

	// Silinen kategori bilgisini JSON olarak döndürüyoruz
	respondWithJSON(w, 200, category)
}

func createColor(w http.ResponseWriter, r *http.Request) {
	var color Color
	err := json.NewDecoder(r.Body).Decode(&color)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// SQL sorgusu ile yeni renk ekliyoruz
	sql := `INSERT INTO "Colors" ("color_id", "name") 
			VALUES ($1, $2) RETURNING "color_id"`

	// Yeni UUID oluşturuyoruz
	colorID := uuid.New()
	err = db.QueryRow(context.Background(), sql, colorID, color.Name).
		Scan(&colorID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	color.ColorID = colorID
	respondWithJSON(w, 200, color)
}

func getColors(w http.ResponseWriter, r *http.Request) {
	sql := `SELECT "color_id", "name" FROM "Colors"`

	rows, err := db.Query(context.Background(), sql)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var colors []Color
	for rows.Next() {
		var color Color
		err := rows.Scan(&color.ColorID, &color.Name)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		colors = append(colors, color)
	}

	if err := rows.Err(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respondWithJSON(w, 200, colors)
}

func getColorByID(w http.ResponseWriter, r *http.Request) {
	colorID := chi.URLParam(r, "id") // URL parametresinden ID'yi alıyoruz

	sql := `SELECT "color_id", "name" 
			FROM "Colors" WHERE "color_id" = $1`

	var color Color
	err := db.QueryRow(context.Background(), sql, colorID).Scan(&color.ColorID, &color.Name)
	if err != nil {
		if err.Error() == "no rows in result set" {
			http.Error(w, "Color not found", http.StatusNotFound)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	respondWithJSON(w, 200, color)
}

func updateColor(w http.ResponseWriter, r *http.Request) {
	// URL parametresinden color_id'yi alıyoruz
	colorID := chi.URLParam(r, "id")

	var color Color
	// Request'ten gelen JSON verisini Color struct'ına parse ediyoruz
	err := json.NewDecoder(r.Body).Decode(&color)
	if err != nil {
		http.Error(w, "Invalid JSON body", http.StatusBadRequest)
		return
	}

	// SQL sorgusunu hazırlıyoruz (Rengi güncellemek için)
	sql := `UPDATE "Colors" 
			SET "name" = $1
			WHERE "color_id" = $2 
			RETURNING "color_id", "name"`

	// Veritabanı sorgusunu çalıştırıyoruz
	err = db.QueryRow(context.Background(), sql, color.Name, colorID).
		Scan(&color.ColorID, &color.Name)

	if err != nil {
		http.Error(w, "Error updating color", http.StatusInternalServerError)
		log.Printf("Error updating color: %v", err)
		return
	}

	// Güncellenmiş renk bilgisini JSON olarak döndürüyoruz
	respondWithJSON(w, 200, color)
}

func deleteColor(w http.ResponseWriter, r *http.Request) {
	// URL parametresinden color_id'yi alıyoruz
	colorID := chi.URLParam(r, "id")

	var color Color
	// SQL sorgusu ile rengi siliyoruz
	sql := `DELETE FROM "Colors" WHERE "color_id" = $1 
			RETURNING "color_id", "name"`

	// Veritabanı sorgusunu çalıştırıyoruz
	err := db.QueryRow(context.Background(), sql, colorID).
		Scan(&color.ColorID, &color.Name)

	if err != nil {
		// Hata durumunda 404 ya da 500 hatası dönüyoruz
		if err.Error() == "no rows in result set" {
			// Eğer renk bulunamazsa
			http.Error(w, "Color not found", http.StatusNotFound)
		} else {
			http.Error(w, "Error deleting color", http.StatusInternalServerError)
		}
		log.Printf("Error deleting color: %v", err)
		return
	}

	// Silinen renk bilgisini JSON olarak döndürüyoruz
	respondWithJSON(w, 200, color)
}

func createCartItem(w http.ResponseWriter, r *http.Request) {
	var cartItem CartItem
	err := json.NewDecoder(r.Body).Decode(&cartItem)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// SQL sorgusu ile yeni sepet öğesini ekliyoruz
	sql := `INSERT INTO "CartItems" ("cart_id", "price_id", "quantity", "status") 
			VALUES ($1, $2, $3, $4) RETURNING "cart_id"`

	// Yeni UUID'yi oluşturuyoruz
	cartID := uuid.New()
	err = db.QueryRow(context.Background(), sql,
		cartID, cartItem.PriceID, cartItem.Quantity, cartItem.Status).
		Scan(&cartID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	cartItem.CartID = cartID
	respondWithJSON(w, 200, cartItem)
}

func getCartItems(w http.ResponseWriter, r *http.Request) {
	sql := `SELECT "cart_id", "price_id", "quantity", "status" FROM "CartItems"`

	rows, err := db.Query(context.Background(), sql)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var cartItems []CartItem
	for rows.Next() {
		var cartItem CartItem
		err := rows.Scan(&cartItem.CartID, &cartItem.PriceID, &cartItem.Quantity, &cartItem.Status)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		cartItems = append(cartItems, cartItem)
	}

	if err := rows.Err(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respondWithJSON(w, 200, cartItems)
}

func getCartItemByID(w http.ResponseWriter, r *http.Request) {
	cartID := chi.URLParam(r, "id") // URL parametresinden CartID'yi alıyoruz

	sql := `SELECT "cart_id", "price_id", "quantity", "status" 
			FROM "CartItems" WHERE "cart_id" = $1`

	var cartItem CartItem
	err := db.QueryRow(context.Background(), sql, cartID).Scan(&cartItem.CartID, &cartItem.PriceID, &cartItem.Quantity, &cartItem.Status)
	if err != nil {
		if err.Error() == "no rows in result set" {
			http.Error(w, "Cart item not found", http.StatusNotFound)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	respondWithJSON(w, 200, cartItem)
}

func updateCartItem(w http.ResponseWriter, r *http.Request) {
	// URL parametresinden cart_id'yi alıyoruz
	cartID := chi.URLParam(r, "id")

	var cartItem CartItem
	// Request'ten gelen JSON verisini CartItem struct'ına parse ediyoruz
	err := json.NewDecoder(r.Body).Decode(&cartItem)
	if err != nil {
		http.Error(w, "Invalid JSON body", http.StatusBadRequest)
		return
	}

	// SQL sorgusunu hazırlıyoruz (Sepet öğesini güncellemek için)
	sql := `UPDATE "CartItems" 
			SET "price_id" = $1, "quantity" = $2, "status" = $3
			WHERE "cart_id" = $4 
			RETURNING "cart_id", "price_id", "quantity", "status"`

	// Veritabanı sorgusunu çalıştırıyoruz
	err = db.QueryRow(context.Background(), sql,
		cartItem.PriceID, cartItem.Quantity, cartItem.Status, cartID).
		Scan(&cartItem.CartID, &cartItem.PriceID, &cartItem.Quantity, &cartItem.Status)

	if err != nil {
		http.Error(w, "Error updating cart item", http.StatusInternalServerError)
		log.Printf("Error updating cart item: %v", err)
		return
	}

	// Güncellenmiş sepet öğesi bilgisini JSON olarak döndürüyoruz
	respondWithJSON(w, 200, cartItem)
}

func deleteCartItem(w http.ResponseWriter, r *http.Request) {
	// URL parametresinden cart_id'yi alıyoruz
	cartID := chi.URLParam(r, "id")

	var cartItem CartItem
	// SQL sorgusu ile sepet öğesini siliyoruz
	sql := `DELETE FROM "CartItems" WHERE "cart_id" = $1 
			RETURNING "cart_id", "price_id", "quantity", "status"`

	// Veritabanı sorgusunu çalıştırıyoruz
	err := db.QueryRow(context.Background(), sql, cartID).
		Scan(&cartItem.CartID, &cartItem.PriceID, &cartItem.Quantity, &cartItem.Status)

	if err != nil {
		// Hata durumunda 404 ya da 500 hatası dönüyoruz
		if err.Error() == "no rows in result set" {
			// Eğer sepet öğesi bulunamazsa
			http.Error(w, "Cart item not found", http.StatusNotFound)
		} else {
			http.Error(w, "Error deleting cart item", http.StatusInternalServerError)
		}
		log.Printf("Error deleting cart item: %v", err)
		return
	}

	// Silinen sepet öğesi bilgisini JSON olarak döndürüyoruz
	respondWithJSON(w, 200, cartItem)
}

func createOrderItem(w http.ResponseWriter, r *http.Request) {
	var orderItem OrderItem
	err := json.NewDecoder(r.Body).Decode(&orderItem)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// SQL sorgusu ile yeni sipariş öğesini ekliyoruz
	sql := `INSERT INTO "OrderItems" ("order_id", "price_id", "quantity", "price_at_purchase") 
			VALUES ($1, $2, $3, $4) RETURNING "order_id"`

	// Yeni UUID'yi oluşturuyoruz
	orderID := uuid.New()
	err = db.QueryRow(context.Background(), sql,
		orderID, orderItem.PriceID, orderItem.Quantity, orderItem.PriceAtPurchase).
		Scan(&orderID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	orderItem.OrderID = orderID
	respondWithJSON(w, 200, orderItem)
}

func getOrderItems(w http.ResponseWriter, r *http.Request) {
	sql := `SELECT "order_id", "price_id", "quantity", "price_at_purchase" FROM "OrderItems"`

	rows, err := db.Query(context.Background(), sql)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var orderItems []OrderItem
	for rows.Next() {
		var orderItem OrderItem
		err := rows.Scan(&orderItem.OrderID, &orderItem.PriceID, &orderItem.Quantity, &orderItem.PriceAtPurchase)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		orderItems = append(orderItems, orderItem)
	}

	if err := rows.Err(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respondWithJSON(w, 200, orderItems)
}

func getOrderItemByID(w http.ResponseWriter, r *http.Request) {
	orderID := chi.URLParam(r, "id") // URL parametresinden OrderID'yi alıyoruz

	sql := `SELECT "order_id", "price_id", "quantity", "price_at_purchase" 
			FROM "OrderItems" WHERE "order_id" = $1`

	var orderItem OrderItem
	err := db.QueryRow(context.Background(), sql, orderID).Scan(&orderItem.OrderID, &orderItem.PriceID, &orderItem.Quantity, &orderItem.PriceAtPurchase)
	if err != nil {
		if err.Error() == "no rows in result set" {
			http.Error(w, "Order item not found", http.StatusNotFound)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	respondWithJSON(w, 200, orderItem)
}

func updateOrderItem(w http.ResponseWriter, r *http.Request) {
	// URL parametresinden order_id'yi alıyoruz
	orderID := chi.URLParam(r, "id")

	var orderItem OrderItem
	// Request'ten gelen JSON verisini OrderItem struct'ına parse ediyoruz
	err := json.NewDecoder(r.Body).Decode(&orderItem)
	if err != nil {
		http.Error(w, "Invalid JSON body", http.StatusBadRequest)
		return
	}

	// SQL sorgusunu hazırlıyoruz (Sipariş öğesini güncellemek için)
	sql := `UPDATE "OrderItems" 
			SET "price_id" = $1, "quantity" = $2, "price_at_purchase" = $3
			WHERE "order_id" = $4 
			RETURNING "order_id", "price_id", "quantity", "price_at_purchase"`

	// Veritabanı sorgusunu çalıştırıyoruz
	err = db.QueryRow(context.Background(), sql,
		orderItem.PriceID, orderItem.Quantity, orderItem.PriceAtPurchase, orderID).
		Scan(&orderItem.OrderID, &orderItem.PriceID, &orderItem.Quantity, &orderItem.PriceAtPurchase)

	if err != nil {
		http.Error(w, "Error updating order item", http.StatusInternalServerError)
		log.Printf("Error updating order item: %v", err)
		return
	}

	// Güncellenmiş sipariş öğesi bilgisini JSON olarak döndürüyoruz
	respondWithJSON(w, 200, orderItem)
}

func deleteOrderItem(w http.ResponseWriter, r *http.Request) {
	// URL parametresinden order_id'yi alıyoruz
	orderID := chi.URLParam(r, "id")

	var orderItem OrderItem
	// SQL sorgusu ile sipariş öğesini siliyoruz
	sql := `DELETE FROM "OrderItems" WHERE "order_id" = $1 
			RETURNING "order_id", "price_id", "quantity", "price_at_purchase"`

	// Veritabanı sorgusunu çalıştırıyoruz
	err := db.QueryRow(context.Background(), sql, orderID).
		Scan(&orderItem.OrderID, &orderItem.PriceID, &orderItem.Quantity, &orderItem.PriceAtPurchase)

	if err != nil {
		// Hata durumunda 404 ya da 500 hatası dönüyoruz
		if err.Error() == "no rows in result set" {
			// Eğer sipariş öğesi bulunamazsa
			http.Error(w, "Order item not found", http.StatusNotFound)
		} else {
			http.Error(w, "Error deleting order item", http.StatusInternalServerError)
		}
		log.Printf("Error deleting order item: %v", err)
		return
	}

	// Silinen sipariş öğesi bilgisini JSON olarak döndürüyoruz
	respondWithJSON(w, 200, orderItem)
}

func createMarketComment(w http.ResponseWriter, r *http.Request) {
	var marketComment MarketComment
	err := json.NewDecoder(r.Body).Decode(&marketComment)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// SQL sorgusu ile yeni yorum ekliyoruz
	sql := `INSERT INTO "MarketComments" ("mcomment_id", "customer_id", "market_id", "star", "comment", "date") 
			VALUES ($1, $2, $3, $4, $5, $6) RETURNING "mcomment_id"`

	// Yeni UUID'yi oluşturuyoruz
	mcommentID := uuid.New()
	err = db.QueryRow(context.Background(), sql,
		mcommentID, marketComment.CustomerID, marketComment.MarketID,
		marketComment.Star, marketComment.Comment, marketComment.Date).
		Scan(&mcommentID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	marketComment.MCommentID = mcommentID
	respondWithJSON(w, 200, marketComment)
}

func getMarketComments(w http.ResponseWriter, r *http.Request) {
	sql := `SELECT "mcomment_id", "customer_id", "market_id", "star", "comment", "date" FROM "MarketComments"`

	rows, err := db.Query(context.Background(), sql)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var marketComments []MarketComment
	for rows.Next() {
		var marketComment MarketComment
		err := rows.Scan(&marketComment.MCommentID, &marketComment.CustomerID, &marketComment.MarketID,
			&marketComment.Star, &marketComment.Comment, &marketComment.Date)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		marketComments = append(marketComments, marketComment)
	}

	if err := rows.Err(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respondWithJSON(w, 200, marketComments)
}

func getMarketCommentByID(w http.ResponseWriter, r *http.Request) {
	mcommentID := chi.URLParam(r, "id") // URL parametresinden MCommentID'yi alıyoruz

	sql := `SELECT "mcomment_id", "customer_id", "market_id", "star", "comment", "date" 
			FROM "MarketComments" WHERE "mcomment_id" = $1`

	var marketComment MarketComment
	err := db.QueryRow(context.Background(), sql, mcommentID).Scan(&marketComment.MCommentID, &marketComment.CustomerID,
		&marketComment.MarketID, &marketComment.Star, &marketComment.Comment, &marketComment.Date)
	if err != nil {
		if err.Error() == "no rows in result set" {
			http.Error(w, "Market comment not found", http.StatusNotFound)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	respondWithJSON(w, 200, marketComment)
}

func updateMarketComment(w http.ResponseWriter, r *http.Request) {
	// URL parametresinden mcomment_id'yi alıyoruz
	mcommentID := chi.URLParam(r, "id")

	var marketComment MarketComment
	// Request'ten gelen JSON verisini MarketComment struct'ına parse ediyoruz
	err := json.NewDecoder(r.Body).Decode(&marketComment)
	if err != nil {
		http.Error(w, "Invalid JSON body", http.StatusBadRequest)
		return
	}

	// SQL sorgusunu hazırlıyoruz (Yorumu güncellemek için)
	sql := `UPDATE "MarketComments" 
			SET "customer_id" = $1, "market_id" = $2, "star" = $3, "comment" = $4, "date" = $5
			WHERE "mcomment_id" = $6 
			RETURNING "mcomment_id", "customer_id", "market_id", "star", "comment", "date"`

	// Veritabanı sorgusunu çalıştırıyoruz
	err = db.QueryRow(context.Background(), sql,
		marketComment.CustomerID, marketComment.MarketID, marketComment.Star,
		marketComment.Comment, marketComment.Date, mcommentID).
		Scan(&marketComment.MCommentID, &marketComment.CustomerID, &marketComment.MarketID,
			&marketComment.Star, &marketComment.Comment, &marketComment.Date)

	if err != nil {
		http.Error(w, "Error updating market comment", http.StatusInternalServerError)
		log.Printf("Error updating market comment: %v", err)
		return
	}

	// Güncellenmiş yorum bilgisini JSON olarak döndürüyoruz
	respondWithJSON(w, 200, marketComment)
}

func deleteMarketComment(w http.ResponseWriter, r *http.Request) {
	// URL parametresinden mcomment_id'yi alıyoruz
	mcommentID := chi.URLParam(r, "id")

	var marketComment MarketComment
	// SQL sorgusu ile yorumu siliyoruz
	sql := `DELETE FROM "MarketComments" WHERE "mcomment_id" = $1 
			RETURNING "mcomment_id", "customer_id", "market_id", "star", "comment", "date"`

	// Veritabanı sorgusunu çalıştırıyoruz
	err := db.QueryRow(context.Background(), sql, mcommentID).
		Scan(&marketComment.MCommentID, &marketComment.CustomerID, &marketComment.MarketID,
			&marketComment.Star, &marketComment.Comment, &marketComment.Date)

	if err != nil {
		// Hata durumunda 404 ya da 500 hatası dönüyoruz
		if err.Error() == "no rows in result set" {
			// Eğer yorum bulunamazsa
			http.Error(w, "Market comment not found", http.StatusNotFound)
		} else {
			http.Error(w, "Error deleting market comment", http.StatusInternalServerError)
		}
		log.Printf("Error deleting market comment: %v", err)
		return
	}

	// Silinen yorum bilgisini JSON olarak döndürüyoruz
	respondWithJSON(w, 200, marketComment)
}

func createProductComment(w http.ResponseWriter, r *http.Request) {
	var productComment ProductComment
	err := json.NewDecoder(r.Body).Decode(&productComment)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// SQL sorgusu ile yeni yorum ekliyoruz
	sql := `INSERT INTO "ProductComments" ("pcomment_id", "customer_id", "product_id", "star", "comment", "date") 
			VALUES ($1, $2, $3, $4, $5, $6) RETURNING "pcomment_id"`

	// Yeni UUID'yi oluşturuyoruz
	pcommentID := uuid.New()
	err = db.QueryRow(context.Background(), sql,
		pcommentID, productComment.CustomerID, productComment.ProductID,
		productComment.Star, productComment.Comment, productComment.Date).
		Scan(&pcommentID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	productComment.PCommentID = pcommentID
	respondWithJSON(w, 200, productComment)
}

func getProductComments(w http.ResponseWriter, r *http.Request) {
	sql := `SELECT "pcomment_id", "customer_id", "product_id", "star", "comment", "date" FROM "ProductComments"`

	rows, err := db.Query(context.Background(), sql)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var productComments []ProductComment
	for rows.Next() {
		var productComment ProductComment
		err := rows.Scan(&productComment.PCommentID, &productComment.CustomerID, &productComment.ProductID,
			&productComment.Star, &productComment.Comment, &productComment.Date)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		productComments = append(productComments, productComment)
	}

	if err := rows.Err(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respondWithJSON(w, 200, productComments)
}

func getProductCommentByID(w http.ResponseWriter, r *http.Request) {
	pcommentID := chi.URLParam(r, "id") // URL parametresinden PCommentID'yi alıyoruz

	sql := `SELECT "pcomment_id", "customer_id", "product_id", "star", "comment", "date" 
			FROM "ProductComments" WHERE "pcomment_id" = $1`

	var productComment ProductComment
	err := db.QueryRow(context.Background(), sql, pcommentID).Scan(&productComment.PCommentID, &productComment.CustomerID,
		&productComment.ProductID, &productComment.Star, &productComment.Comment, &productComment.Date)
	if err != nil {
		if err.Error() == "no rows in result set" {
			http.Error(w, "Product comment not found", http.StatusNotFound)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	respondWithJSON(w, 200, productComment)
}

func updateProductComment(w http.ResponseWriter, r *http.Request) {
	// URL parametresinden pcomment_id'yi alıyoruz
	pcommentID := chi.URLParam(r, "id")

	var productComment ProductComment
	// Request'ten gelen JSON verisini ProductComment struct'ına parse ediyoruz
	err := json.NewDecoder(r.Body).Decode(&productComment)
	if err != nil {
		http.Error(w, "Invalid JSON body", http.StatusBadRequest)
		return
	}

	// SQL sorgusunu hazırlıyoruz (Yorumu güncellemek için)
	sql := `UPDATE "ProductComments" 
			SET "customer_id" = $1, "product_id" = $2, "star" = $3, "comment" = $4, "date" = $5
			WHERE "pcomment_id" = $6 
			RETURNING "pcomment_id", "customer_id", "product_id", "star", "comment", "date"`

	// Veritabanı sorgusunu çalıştırıyoruz
	err = db.QueryRow(context.Background(), sql,
		productComment.CustomerID, productComment.ProductID, productComment.Star,
		productComment.Comment, productComment.Date, pcommentID).
		Scan(&productComment.PCommentID, &productComment.CustomerID, &productComment.ProductID,
			&productComment.Star, &productComment.Comment, &productComment.Date)

	if err != nil {
		http.Error(w, "Error updating product comment", http.StatusInternalServerError)
		log.Printf("Error updating product comment: %v", err)
		return
	}

	// Güncellenmiş yorum bilgisini JSON olarak döndürüyoruz
	respondWithJSON(w, 200, productComment)
}

func deleteProductComment(w http.ResponseWriter, r *http.Request) {
	// URL parametresinden pcomment_id'yi alıyoruz
	pcommentID := chi.URLParam(r, "id")

	var productComment ProductComment
	// SQL sorgusu ile yorumu siliyoruz
	sql := `DELETE FROM "ProductComments" WHERE "pcomment_id" = $1 
			RETURNING "pcomment_id", "customer_id", "product_id", "star", "comment", "date"`

	// Veritabanı sorgusunu çalıştırıyoruz
	err := db.QueryRow(context.Background(), sql, pcommentID).
		Scan(&productComment.PCommentID, &productComment.CustomerID, &productComment.ProductID,
			&productComment.Star, &productComment.Comment, &productComment.Date)

	if err != nil {
		// Hata durumunda 404 ya da 500 hatası dönüyoruz
		if err.Error() == "no rows in result set" {
			// Eğer yorum bulunamazsa
			http.Error(w, "Product comment not found", http.StatusNotFound)
		} else {
			http.Error(w, "Error deleting product comment", http.StatusInternalServerError)
		}
		log.Printf("Error deleting product comment: %v", err)
		return
	}

	// Silinen yorum bilgisini JSON olarak döndürüyoruz
	respondWithJSON(w, 200, productComment)
}
