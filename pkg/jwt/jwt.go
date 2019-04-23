package jwt

import (
	"errors"
	"strconv"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

const tokenTTL = time.Hour * 1

type Data struct {
	User uint
	Role string
	TTL  *time.Time
}

func (d *Data) Expired() bool {
	if d.TTL == nil {
		return true
	}
	now := time.Now().UTC()
	return now.After(*d.TTL)
}

func Build(data Data) (*string, error) {
	ttl := time.Now().UTC().Add(tokenTTL)
	if data.TTL != nil {
		ttl = *data.TTL
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user": data.User,
		"role": data.Role,
		"ttl":  strconv.FormatInt(ttl.Unix(), 10),
	})
	strToken, err := token.SignedString(config.secretKey)
	if err != nil {
		return nil, err
	}
	return &strToken, nil
}

func Parse(strToken string) (*Data, error) {
	token, err := jwt.Parse(strToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return config.secretKey, nil
	})
	if err != nil {
		return nil, errors.New("invalid token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if err != nil || !ok || !token.Valid {
		return nil, errors.New("invalid token")
	}

	uttl, err := strconv.ParseInt(claims["ttl"].(string), 10, 64)
	if err != nil {
		return nil, errors.New("invalid ttl format")
	}
	ttl := time.Unix(uttl, 0)
	return &Data{
		User: uint(claims["user"].(float64)),
		Role: claims["role"].(string),
		TTL:  &ttl,
	}, nil
}
