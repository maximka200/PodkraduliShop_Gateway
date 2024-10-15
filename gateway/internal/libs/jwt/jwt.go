package jwt

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func CheckJWT(ctx context.Context, log *slog.Logger, jwtToken string) (bool, error) {
	const op = "libs.jwt"

	sign := []byte(os.Getenv("SECRET_KEY"))
	if sign == nil {
		return false, fmt.Errorf("SECRET_KEY is empty")
	}

	token, err := jwt.Parse(jwtToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("%s: %v", op, token.Header["alg"])
		}
		return sign, nil
	})
	if err != nil {
		return false, fmt.Errorf("%s: %w", op, err)
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		exp, ok := claims["exp"].(float64) // exp returned that float64
		if !ok {
			return false, fmt.Errorf("%s: cannot assert exp to float64", op)
		}

		expTime := time.Unix(int64(exp), 0)
		if time.Now().After(expTime) {
			return false, fmt.Errorf("%s: token has expired", op)
		}

		id := int(claims["uid"].(float64))
		if !ok {
			return false, fmt.Errorf("%s: cannot assert uid ", op)
		}

		log.Info(fmt.Sprintf("%s: %s, %d", op, expTime, id))

	} else {
		return false, fmt.Errorf("%s: token is not valid", op)
	}

	return true, nil
}
