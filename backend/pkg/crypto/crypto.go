package crypto

import (
	"fmt"
	"log"
	"strings"
	"time"
	config2 "users_api/internal/pkg/config"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

// HashAndSalt takes a password as input and returns the hashed and salted version of the password.
// It uses bcrypt algorithm to generate the hash with minimum cost.
// If an error occurs during the hashing process, it will be logged and an empty string will be returned.
func HashAndSalt(pwd []byte) string {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		log.Println(err)
	}
	return string(hash)
}

// ComparePasswords compares a hashed password with a plain password and returns true if they match, false otherwise.
// It takes a hashedPwd string and a plainPwd []byte as input.
// The function uses bcrypt.CompareHashAndPassword to compare the hashed password with the plain password.
// If there is an error during the comparison, it returns false.
// Otherwise, it returns true.
func ComparePasswords(hashedPwd string, plainPwd []byte) bool {
	byteHash := []byte(hashedPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, plainPwd)
	if err != nil {
		return false
	}
	return true
}

// CreateToken generates a JWT token for the given username.
// It sets the "authorized" claim to true, the "username" claim to the provided username,
// and the "exp" claim to the current time plus one year.
// The token is signed using the secret key from the server configuration.
// If an error occurs during token creation, it returns a "token creation error" string and the error.
// Otherwise, it returns the generated token and nil error.
func CreateToken(username string) (string, error) {
	config := config2.GetConfig()

	var err error
	//Creating Access Token
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["username"] = username
	atClaims["exp"] = time.Now().Add(time.Hour * 24 * 365).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS512, atClaims)
	token, err := at.SignedString([]byte(config.Server.Secret)) // SECRET
	if err != nil {
		return "token creation error", err
	}
	return token, nil
}

// ValidateToken validates the given token string.
// It splits the token string by "Bearer " and checks if there are exactly two parts.
// Then it retrieves the token string from the second part.
// It retrieves the secret key from the server configuration.
// It parses the token using the secret key and checks if it is valid.
// Returns true if the token is valid, otherwise false.
func ValidateToken(tokenString string) bool {
	tokenStrings := strings.Split(tokenString, "Bearer ")
	if len(tokenStrings) != 2 {
		return false
	}
	tokenString = tokenStrings[1]
	config := config2.GetConfig()
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("there was an error")
		}
		return []byte(config.Server.Secret), nil
	})
	if err != nil {
		return false
	}
	return token.Valid
}
