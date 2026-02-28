package auth

import (
	"errors"
	"time"

	"github.com/Mickdevv/moonless/backend/api/utils"
	"github.com/alexedwards/argon2id"
	"github.com/golang-jwt/jwt/v5"
)

func makeJWT(serverCfg *utils.ServerCfg, claims CustomJwtClaims) (string, error) {
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	s, err := t.SignedString([]byte(serverCfg.JWT_SECRET))
	if err != nil {
		return "", err
	}
	return s, nil
}

func validateJWT(serverCfg *utils.ServerCfg, token string) (CustomJwtClaims, error) {
	t, err := jwt.ParseWithClaims(token, &CustomJwtClaims{}, func(t *jwt.Token) (any, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); ok {
			return nil, errors.New("Incorrect signing method")
		}
		return []byte(serverCfg.JWT_SECRET), nil
	}, jwt.WithLeeway(5*time.Second))
	if err != nil {
		return CustomJwtClaims{}, err
	}

	if !t.Valid {
		return CustomJwtClaims{}, errors.New("Invalid token")
	}

	if c, ok := t.Claims.(*CustomJwtClaims); ok {
		return *c, nil
	}

	return CustomJwtClaims{}, errors.New("JWT validation: Unknown error occurred")
}

func makePassword(password string) (string, error) {

	hash, err := argon2id.CreateHash(password, argon2id.DefaultParams)
	if err != nil {
		return "", err
	}
	return hash, nil
}

func checkPassword(password string, hash string) (bool, error) {
	ok, err := argon2id.ComparePasswordAndHash(password, hash)
	if err != nil {
		return false, err
	}

	return ok, nil
}
