package authorization

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/hreluz/go-api-crud/model"
)

func GenerateToken(data *model.Login) (string, error) {
	claim := model.Claim{
		Email: data.Email,
		StandardClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 2)),
			Issuer:    "Hector",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claim.StandardClaims)
	signedToken, err := token.SignedString(signKey)

	if err != nil {
		return "", err
	}

	return signedToken, nil
}
