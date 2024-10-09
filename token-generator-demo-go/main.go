package main

import (
	"crypto/rsa"
	"encoding/json"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
	"time"
)

var privateKey *rsa.PrivateKey

func initKeys() error {
	keyPath := os.Getenv("PRIVATE_KEY_PATH")
	if keyPath == "" {
		// use example key if no key path is provided
		log.Println("No private key path provided, using example key")
		keyPath = "private.pem"
	}

	keyData, err := os.ReadFile(keyPath)
	if err != nil {
		return fmt.Errorf("Error reading private key file: %v", err)
	}

	privateKey, err = jwt.ParseRSAPrivateKeyFromPEM(keyData)
	if err != nil {
		return fmt.Errorf("Error parsing private key: %v", err)
	}
	return nil
}

func getTokenHandler(w http.ResponseWriter, r *http.Request) {
	sub := r.URL.Query().Get("sub")
	if sub == "" {
		http.Error(w, "Missing 'sub' parameter", http.StatusBadRequest)
		return
	}

	now := time.Now().Unix()

	// Create the Claims
	claims := jwt.MapClaims{
		"sub": sub,                                  // Subject (e.g., user ID)
		"iat": now,                                  // Issued at
		"nbf": now,                                  // Not before
		"exp": time.Now().Add(time.Hour * 6).Unix(), // Expires in 6 hours
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	signedToken, err := token.SignedString(privateKey)
	if err != nil {
		http.Error(w, "Error signing token", http.StatusInternalServerError)
		return
	}

	// Set the Content-Type header to application/json
	w.Header().Set("Content-Type", "application/json")

	// Create a response map
	response := map[string]string{
		"token": signedToken,
	}

	// Encode the response as JSON
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, "Error encoding JSON", http.StatusInternalServerError)
		return
	}
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found or error reading .env file")
	}

	err = initKeys()
	if err != nil {
		log.Fatalf("Failed to initialize keys: %v", err)
	}

	http.HandleFunc("/get-token/", getTokenHandler)
	log.Println("Starting server on :3000")
	log.Fatal(http.ListenAndServe(":3000", nil))
}
