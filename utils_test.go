package pwt

import (
	"encoding/base64"
	gogoproto "github.com/gogo/protobuf/proto"
	"testing"
	"time"
)

func stubDate() *time.Time {
	ts, _ := time.Parse(time.RFC3339, "2021-01-07T18:41:00Z")
	return &ts
}

func handleError(t *testing.T, err error) {
	if err != nil {
		t.Fatal(err)
	}
}

func stubToken() *Token {
	return &Token{
		Algo:   ALGO_HS384,
		Claims: Claims{
			IssuedAt:  stubDate(),
			ExpiresAt: stubDate(),
			NotBefore: stubDate(),
			PWTID:     "test",
			Audience:  "test",
			Issuer:    "test",
			Subject:   "test",
		},
		Extra:  nil,
	}
}

func mustMarshalToken(t *testing.T, token *Token) []byte {
	body, err := gogoproto.Marshal(token)
	handleError(t, err)
	return body
}

func toBase64(data []byte) string {
	return base64.StdEncoding.EncodeToString(data)
}
