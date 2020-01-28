package app

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"regexp"
	"time"
)

func jwtGenerateToken(uuid string) (string, error) {
	now := time.Now()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		`iss`: uuid,
		`exp`: now.AddDate(0, 1, 0).Unix(),
		`nbf`: now.Unix(),
		`iat`: now.Unix(),
	})
	tokenString, err := token.SignedString([]byte(env.JwtSecret))
	if err != nil {
		return "", log.TError(grpc.Errorf(codes.Internal, "jwt failed"), err, uuid)
	}
	return tokenString, nil
}

// returns uuid, isValidated
func jwtValidateToken(bearerTokenString string) (string, bool) {
	if bearerTokenString == "" {
		log.Error("empty token")
		return "", false
	}

	re := regexp.MustCompile(`(?:b|B)earer +(.*)`)
	authToken := re.FindStringSubmatch(bearerTokenString)
	if authToken == nil {
		log.Error("regex bearer <token> failed")
		return "", false
	}

	token, err := jwt.Parse(authToken[1], func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf(`Unexpected signing method: %v`, t.Header[`alg`])
		}
		return []byte(env.JwtSecret), nil
	})

	if err != nil {
		log.Error(err)
		return "", false
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return "", false
	}

	if claims[`exp`].(float64) < float64(time.Now().Unix()) {
		log.Error(fmt.Errorf("Token expired"))
		return "", false
	}
	uuid := claims[`iss`].(string)
	return uuid, true
}

func jwtRefreshToken(token string) (*string, error) {
	return nil, errors.New("unimplemented")
}

const (
	msgUnauthorized = `Unauthorized`
)
