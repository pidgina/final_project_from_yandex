package service

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"

	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	jwt.RegisteredClaims
	PasswordHash string `json:"passwordhash"`
}
type TokenJWT struct {
	Token string `json:"token"`
}

var Secret string = "YandexPracticum"
var InvalidPassword = errors.New("Неверный пароль")
var ErrServer = errors.New("Внутреняя ошибка")

func LoginAuther(passer PasswordType) (TokenJWT, error) {
	var tokenJWT TokenJWT

	if passer.Password != AutherENV() {
		return TokenJWT{}, fmt.Errorf("%w: Неверный пароль", InvalidPassword)
	} else {
		hashedPassword := sha256.Sum256([]byte(passer.Password))
		hashStringPassword := hex.EncodeToString(hashedPassword[:])

		var claims Claims
		claims.PasswordHash = hashStringPassword

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		tokenString, err := token.SignedString([]byte(Secret))
		if err != nil {
			return TokenJWT{}, fmt.Errorf("%w: Ошибка перевода токена в строку", ErrServer)
		}

		tokenJWT.Token = tokenString
	}

	return tokenJWT, nil
}
