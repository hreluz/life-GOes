package authorization

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/hreluz/echo-framework/model"
)

func GenerateToken(data *model.Login) (string, error) {
	claim := model.Claim{
		Email: data.Email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 2)),
			Issuer:    "Hector",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claim.RegisteredClaims)
	signedToken, err := token.SignedString(signKey)

	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func ValidateToken(t string) (model.Claim, error) {
	token, err := jwt.ParseWithClaims(t, &model.Claim{}, verifyFunction)

	if err != nil {
		return model.Claim{}, err
	}

	if !token.Valid {
		return model.Claim{}, errors.New("token not valid")
	}

	claim, ok := token.Claims.(*model.Claim)

	if !ok {
		return model.Claim{}, errors.New("claim could not be found")
	}

	return *claim, nil
}

func verifyFunction(t *jwt.Token) (interface{}, error) {
	return verifyKey, nil
}
