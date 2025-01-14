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
