package util

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"math/rand"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func EncryptPassword(password, salt string) string {
	mp := md5.New()
	mp.Write([]byte(password))

	ms := md5.New()
	ms.Write([]byte(hex.EncodeToString(mp.Sum(nil)) + salt))

	return hex.EncodeToString(ms.Sum(nil))
}

func BuildJwtToken(key string, Claims jwt.MapClaims) (string, error) {
	token := jwt.New(jwt.GetSigningMethod("HS256"))
	token.Claims = Claims
	tokenString, err := token.SignedString([]byte(key))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func ParseJwtToken(key, token string) (jwt.MapClaims, error) {
	t, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(key), nil

	})

	if err != nil || !t.Valid {
		return nil, err
	}

	if mc, ok := t.Claims.(jwt.MapClaims); ok {
		return mc, nil
	}

	return nil, fmt.Errorf("interface.(jwt.MapClaims) error")
}

func GetRandomNumber(length int) string {
	bytes := []byte("0123456789")

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	result := make([]byte, 0, length)
	for i := 0; i < length; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}

	return string(result)
}

func GetRandomString(length int) string {
	bytes := []byte("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	var result []byte
	for i := 0; i < length; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}

	return string(result)
}

func Md5(src string) string {
	h := md5.New()
	h.Write([]byte(src))
	d := h.Sum(nil)

	return string(hex.EncodeToString(d))
}
