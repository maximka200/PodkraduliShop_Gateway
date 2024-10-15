package main

import (
	"context"
	"fmt"
	"gateway/internal/libs/jwt"
	"log/slog"
	"os"
)

func main() {
	err := os.Setenv("SECRET_KEY", "test-secret")
	if err != nil {
		fmt.Print(err.Error())
	}
	log := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	token := "eyJhbGciOiJIUzI1NiIsInR5cCIsdssdsdsa16IkpXVCJ9.eyJhcHBfaWQiOjEsImVtYWlsIjoic2FkYXNkQGFkc3NzYXNzYWQucnUiLCJleHAiOjE3MjkwMDU5MzAsInVpZCI6Nn0.h-KuLtFgZEh5E6xYo4yjvyyTHoStQgLsKsuDvwOuYY0"

	ok, err := jwt.CheckJWT(context.Background(), log, token)
	if err != nil {
		fmt.Print(err.Error())
	}
	fmt.Printf("%t", ok)
}
