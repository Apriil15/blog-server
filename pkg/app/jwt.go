package app

import (
	"time"

	"github.com/Apriil15/blog-server/global"
	"github.com/Apriil15/blog-server/pkg/util"
	"github.com/dgrijalva/jwt-go"
)

type Claims struct {
	AppKey    string `json:"app_key"`
	AppSecret string `json:"app_secret"`
	jwt.StandardClaims
}

// Get JWT secret written in config.yaml
func GetJWTSecret() []byte {
	return []byte(global.JWTSetting.Secret)
}

// Generate JWT token
func GenerateToken(appKey, appSecret string) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(global.JWTSetting.Expire)
	claims := Claims{
		AppKey:    util.EncodeMD5(appKey),
		AppSecret: util.EncodeMD5(appSecret),
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    global.JWTSetting.Issuer,
		},
	}

	// use claims and get a instance of token
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// use secret, and get the complete, signed token
	token, err := tokenClaims.SignedString(GetJWTSecret())
	return token, err
}

// Parse JWT token
func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return GetJWTSecret(), nil
	})

	if tokenClaims != nil {
		claims, ok := tokenClaims.Claims.(*Claims)
		if ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err
}
