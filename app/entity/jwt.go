package entity

import (
	"fmt"
	jwt "github.com/dgrijalva/jwt-go"
)

type JWT struct {
	SigninMethod jwt.SigningMethod
	Claims       Claims
}

type Claims struct {
	Sub string
	Iat int64
	Exp int64
}

func NewJWT(claims Claims, signinMethod jwt.SigningMethod) (*JWT, error) {
	return &JWT{
		SigninMethod: signinMethod,
		Claims:       claims,
	}, nil
}

func (j JWT) ToTokenString(signingKey string) (string, error) {
	token := jwt.New(j.SigninMethod)

	claims := token.Claims.(jwt.MapClaims)
	claims["sub"] = j.Claims.Sub
	claims["iat"] = j.Claims.Iat
	claims["exp"] = j.Claims.Exp

	token = jwt.NewWithClaims(j.SigninMethod, claims)
	// // 電子署名
	tokenString, _ := token.SignedString([]byte(signingKey))
	return tokenString, nil
}

func ParseJwtToken(signedString string, signingMethod jwt.SigningMethod, secret string) (*Claims, error) {
	token, err := jwt.Parse(signedString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(jwt.SigningMethod); !ok {
			return "", fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secret), nil
	})

	if err != nil {
		return nil, err
	}

	if token == nil {
		return nil, fmt.Errorf("not found token in %s:", signedString)
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	fmt.Printf("%+v", claims)
	if !ok {
		return nil, fmt.Errorf("not found claims in %s", signedString)
	}
	sub, ok := claims["sub"].(string)
	if !ok {
		return nil, fmt.Errorf("not found %s in %s", "sub", signedString)
	}
	iat, ok := claims["iat"].(float64)
	if !ok {
		return nil, fmt.Errorf("not found %s in %s", "iat", signedString)
	}
	exp, ok := claims["exp"].(float64)
	if !ok {
		return nil, fmt.Errorf("not found %s in %s", "exp", signedString)
	}

	return &Claims{
		Sub: sub,
		Iat: int64(iat),
		Exp: int64(exp),
	}, nil
}
