package auth

import (
	"dena-hackathon21/entity"
	"fmt"
	jwt "github.com/dgrijalva/jwt-go"
	"os"
	"strconv"
	"time"
)

type JWTHandler struct {
	SigninKey string
}

func NewJWTHandler() (*JWTHandler, error) {
	return &JWTHandler{}, nil
}

func (j JWTHandler) GenerateJWTToken(userID uint64) (string, error) {

	claims := entity.Claims{
		Sub: strconv.Itoa(int(userID)),
		Iat: time.Now().Unix(),
		Exp: time.Now().Add(time.Hour * 24).Unix(),
	}

	jwtEntity := entity.JWT{
		SigninMethod: jwt.SigningMethodHS256,
		Claims:       claims,
	}

	// 電子署名
	tokenString, _ := jwtEntity.ToTokenString(os.Getenv("SIGNINGKEY"))
	return tokenString, nil
}

func (j JWTHandler) GetClaimsFromToken(tokenStr string) (*entity.Claims, error) {
	claims, err := entity.ParseJwtToken(tokenStr, jwt.SigningMethodHS256, os.Getenv("SIGNINGKEY"))
	return claims, err
}

func (j JWTHandler) Valid(tokenStr string) (bool, error) {
	claims, err := j.GetClaimsFromToken(tokenStr)
	if err != nil {
		return false, err
	}
	if claims.Exp < time.Now().Unix() {
		return false, fmt.Errorf("timeout")
	}
	return true, nil
}

func (j JWTHandler) GetUserIDFromToken(tokenStr string) (uint64, error) {
	claims, err := j.GetClaimsFromToken(tokenStr)
	if err != nil {
		return 0, err
	}

	i, err := strconv.Atoi(claims.Sub)
	if err != nil {
		return 0, err
	}
	return uint64(i), nil
}
