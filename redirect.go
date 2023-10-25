package main

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
)

func main() {
	apiKey := "API_KEY" // Specify your API key

	phtoken := "your_jwt_token_here" // Replace with the JWT token you want to decode

	jwtResponse := 0
	var jwtPhone string

	token, err := jwt.Parse(phtoken, func(token *jwt.Token) (interface{}, error) {
		// Check the signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Invalid signing method")
		}
		// Return the key used for signing
		return []byte(apiKey), nil
	})

	if err == nil && token.Valid {
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			jwtResponse = 0 // Invalid JWT
		} else {
			jwtResponse = 1
			countryCode, _ := claims["country_code"].(string)
			phoneNo, _ := claims["phone_no"].(string)
			jwtPhone = countryCode + phoneNo
		}
	} else {
		jwtResponse = 0 // Invalid JWT
	}

	fmt.Printf("JWT Response: %d\n", jwtResponse)
	fmt.Printf("JWT Phone: %s\n", jwtPhone)
}
