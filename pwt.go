package pwt

import (
	"bytes"
	"crypto"
	"errors"
	gogoproto "github.com/gogo/protobuf/proto"
	"strings"
	"time"
)

var (
	ErrTokenInvalid     = errors.New("token is invalid")
	ErrSignatureInvalid = errors.New("signature is invalid")
	ErrMethodInvalid    = errors.New("signing method is invalid")
	ErrKeyInvalid       = errors.New("key is invalid")
)

type SigningMethod interface {
	Verify(signature, body []byte, key interface{}) (bool, error)
	Sign(body []byte, key interface{}) ([]byte, error)
}

func GetSigningMethod(algo Algo) (SigningMethod, error) {
	switch algo {
	case ALGO_HS256:
		return &SigningMethodHMAC{hash: crypto.SHA256}, nil
	case ALGO_HS384:
		return &SigningMethodHMAC{hash: crypto.SHA384}, nil
	case ALGO_HS512:
		return &SigningMethodHMAC{hash: crypto.SHA512}, nil
	//case ALGO_RS256:
	//case ALGO_RS384:
	//case ALGO_RS512:
	//case ALGO_ES256:
	//case ALGO_ES384:
	//case ALGO_ES512:
	//case ALGO_PS256:
	//case ALGO_PS384:
	default:
		return nil, ErrMethodInvalid
	}
}

type KeyFunc = func(method SigningMethod) (interface{}, error)

func Sign(token *Token, kf KeyFunc, codec Codec) (string, error) {
	method, err := GetSigningMethod(token.Algo)
	if err != nil {
		return "", err
	}

	key, err := kf(method)
	if err != nil {
		return "", err
	}

	body, err := gogoproto.Marshal(token)
	if err != nil {
		return "", err
	}

	sign, err := method.Sign(body, key)
	if err != nil {
		return "", err
	}

	signature, err := codec.Encode(sign)
	if err != nil {
		return "", err
	}

	payload, err := codec.Encode(body)
	if err != nil {
		return "", err
	}

	var res bytes.Buffer
	res.WriteString(signature)
	res.WriteByte('.')
	res.WriteString(payload)
	return res.String(), nil
}

func Verify(tok string, kf KeyFunc, codec Codec) (*Token, error) {
	parts := strings.Split(tok, ".")
	if len(parts) != 2 {
		return nil, ErrSignatureInvalid
	}

	sign, err := codec.Decode(parts[0])
	if err != nil {
		return nil, err
	}
	body, err := codec.Decode(parts[1])
	if err != nil {
		return nil, err
	}

	token := new(Token)
	err = gogoproto.Unmarshal(body, token)
	if err != nil {
		return nil, ErrSignatureInvalid
	}
	now := time.Now()
	if token.Claims.ExpiresAt.Before(now) || token.Claims.NotBefore.After(now) {
		return nil, ErrTokenInvalid
	}

	method, err := GetSigningMethod(token.Algo)
	if err != nil {
		return nil, err
	}

	key, err := kf(method)
	if err != nil {
		return nil, err
	}

	valid, err := method.Verify(sign, body, key)
	if err != nil {
		return nil, err
	}
	if !valid {
		return nil, ErrTokenInvalid
	}
	return token, nil
}
