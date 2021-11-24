package sign

import (
	"crypto/rsa"
	"gebi/app/Http/Serializer"
	"github.com/golang-jwt/jwt"
	"io/ioutil"
	"time"
)

type JwtService struct {
	*jwt.StandardClaims
	TokenType string
	Value     interface{}
}

func (customJwt JwtService) SignedData(data interface{}, private_key_path string) (string, error) {
	t := jwt.New(jwt.GetSigningMethod("RS256"))

	t.Claims = &JwtService{
		&jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 4).Unix(),
		},
		"level1",
		data,
	}

	return t.SignedString(getPrivateSignKey(private_key_path))
}

func (customJwt JwtService) ParseData(tokenString string) interface{} {
	// Parse the token
	token, err := jwt.ParseWithClaims(tokenString, &JwtService{}, func(token *jwt.Token) (interface{}, error) {
		return getPublicSignKey("public_key.pub"), nil
	})

	if err != nil {
		Serializer.Err(500, "解密token失败", err)
	}

	claims := token.Claims.(*JwtService)
	return claims.Value
}

func getPublicSignKey(public_key_path string) *rsa.PublicKey {
	keyData, err := ioutil.ReadFile(public_key_path)
	if err != nil {
		Serializer.Err(500, "读取公钥失败", err)
	}

	signKey, err := jwt.ParseRSAPublicKeyFromPEM(keyData)
	if err != nil {
		Serializer.Err(500, "读取解析公钥失败", err)
	}

	return signKey
}

func getPrivateSignKey(private_key_path string) *rsa.PrivateKey {
	keyData, err := ioutil.ReadFile(private_key_path)
	if err != nil {
		Serializer.Err(500, "读取秘钥失败", err)
	}

	signKey, err := jwt.ParseRSAPrivateKeyFromPEM(keyData)
	if err != nil {
		Serializer.Err(500, "读取解析秘钥失败", err)
	}

	return signKey
}
