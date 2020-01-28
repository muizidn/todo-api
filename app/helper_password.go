package app

import (
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

func passwordGenerateHash(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return "", log.TError(grpc.Errorf(codes.Internal, "Password hashing error"), err)
	}
	return string(bytes), nil
}

func passwordValidate(password string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		log.Error(err, hash, `:`, password)
		return false
	}
	return true
}
