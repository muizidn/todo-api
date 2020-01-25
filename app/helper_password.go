package app

import (
	"crypto/sha1"
	"encoding/hex"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

func checkPassword(request string, record string) error {
	h := sha1.New()
	h.Write([]byte(request))
	sha1Hash := hex.EncodeToString(h.Sum(nil))
	if sha1Hash != record {
		return grpc.Errorf(codes.Unauthenticated, "not authenticated")
	}
	return nil
}
