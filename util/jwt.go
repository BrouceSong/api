package util

import (
	jwt "github.com/dgrijalva/jwt-go"
	"time"
)

type JWTClaims struct {
    jwt.StandardClaims
    UserID int `json:"user_id"`
    Username string `json:"username"`
}

var (
	Secret = "go-api-secret"
)

func GetToken(claims *JWTClaims) (string, error) {
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    signedToken, err := token.SignedString([]byte(Secret))
    if err != nil {
        return "", err
    }
    return signedToken, nil
}

//验证jwt token
func VerifyToken(strToken string) (*JWTClaims, error) {
    token, err := jwt.ParseWithClaims(strToken, &JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
        return []byte(Secret), nil
    })
    if err != nil {
        return nil, err
    }

    claims, ok := token.Claims.(*JWTClaims)
    if !ok {
        return nil, err
    }
    if err := token.Claims.Valid(); err != nil {
        return nil, err
    }
    return claims, nil
}

//刷新token
func RefreshToken(strToken string) (string, error) {
    claims, err := VerifyToken(strToken)
    if err != nil {
        return "", err
    }
    claims.ExpiresAt = time.Now().Unix() + (claims.ExpiresAt - claims.IssuedAt)
    signedToken, err := GetToken(claims)
    if err != nil {
        return "", err
	}
	return signedToken, nil
}
