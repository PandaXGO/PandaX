package token

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"strings"
	"time"
)

type Claims struct {
	UserId   int64
	TenantId int64
	UserName string
	RoleId   int64
	RoleKey  string
	DeptId   int64
	PostId   int64
	jwt.StandardClaims
}

type JWT struct {
	SignedKeyID  string
	SignedKey    []byte
	SignedMethod jwt.SigningMethod
}

var (
	TokenExpired     = errors.New("token is expired")
	TokenNotValidYet = errors.New("token not active yet")
	TokenMalformed   = errors.New("that's not even a token")
	TokenInvalid     = errors.New("couldn't handle this token")
)

func NewJWT(kid string, key []byte, method jwt.SigningMethod) *JWT {
	return &JWT{
		SignedKeyID:  kid,
		SignedKey:    key,
		SignedMethod: method,
	}
}

// CreateToken 创建一个token
func (j *JWT) CreateToken(claims Claims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &claims)
	var key interface{}
	if j.isEs() {
		v, err := jwt.ParseECPrivateKeyFromPEM(j.SignedKey)
		if err != nil {
			return "", err
		}
		key = v
	} else if j.isRsOrPS() {
		v, err := jwt.ParseRSAPrivateKeyFromPEM(j.SignedKey)
		if err != nil {
			return "", err
		}
		key = v
	} else if j.isHs() {
		key = j.SignedKey
	} else {
		return "", errors.New("unsupported sign method")
	}
	return token.SignedString(key)
}

// ParseToken 解析 token
func (j *JWT) ParseToken(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (i interface{}, e error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("parse error")
		}
		return j.SignedKey, nil
	})
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, TokenMalformed
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				// Token is expired
				return nil, TokenExpired
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, TokenNotValidYet
			} else {
				return nil, TokenInvalid
			}
		}
	}
	if token != nil {
		if claims, ok := token.Claims.(*Claims); ok && token.Valid {
			return claims, nil
		}
		return nil, TokenInvalid

	} else {
		return nil, TokenInvalid

	}

}

// 更新token
func (j *JWT) RefreshToken(tokenString string) (string, error) {
	jwt.TimeFunc = func() time.Time {
		return time.Unix(0, 0)
	}
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return j.SignedKey, nil
	})
	if err != nil {
		return "", err
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		jwt.TimeFunc = time.Now
		claims.StandardClaims.ExpiresAt = time.Now().Unix() + 60*60*24*7
		return j.CreateToken(*claims)
	}
	return "", TokenInvalid
}

func (a *JWT) isEs() bool {
	return strings.HasPrefix(a.SignedMethod.Alg(), "ES")
}

func (a *JWT) isRsOrPS() bool {
	isRs := strings.HasPrefix(a.SignedMethod.Alg(), "RS")
	isPs := strings.HasPrefix(a.SignedMethod.Alg(), "PS")
	return isRs || isPs
}

func (a *JWT) isHs() bool {
	return strings.HasPrefix(a.SignedMethod.Alg(), "HS")
}
