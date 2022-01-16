package Services

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"os"
	"time"
)

type JWTService interface {
	GenerateToken(name string, admin bool) string
	ValidateToken(tokenString string) (*jwt.Token, error)
}

type jwtCustomerClaims struct {
	Name  string `json:"name"`
	Admin bool   `json:"admin"`
	jwt.StandardClaims
}

type jwtServices struct {
	secretKey string
	issure    string
}

func (j jwtServices) GenerateToken(name string, admin bool) string {
	claims := &jwtCustomerClaims{
		name,
		admin,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 48).Unix(),
			Issuer:    j.issure,
			IssuedAt:  time.Now().Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	//encoded string
	t, err := token.SignedString([]byte(j.secretKey))
	if err != nil {
		panic(err)
	}
	return t
}

func (j jwtServices) ValidateToken(tokenString string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, isvalid := token.Method.(*jwt.SigningMethodHMAC); !isvalid {
			return nil, fmt.Errorf("Invalid token", token.Header["alg"])
		}
		return []byte(j.secretKey), nil
	})
}

func JWTAuthService() JWTService {
	return &jwtServices{
		secretKey: getSecretKey(),
		issure:    "Mehdi",
	}
}

func getSecretKey() string {
	secret := os.Getenv("SECRET")
	if secret == "" {
		secret = "16101374"
	}
	return secret
}
