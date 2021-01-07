package pwt

import (
	"crypto"
	"crypto/hmac"
	_ "crypto/sha256"
	_ "crypto/sha512"
	"errors"
)

type SigningMethodHMAC struct {
	hash crypto.Hash
}

var (
	ErrHashUnavailable = errors.New("hash is not available")
)

func (m *SigningMethodHMAC) Sign(data []byte, key interface{}) ([]byte, error) {
	if !m.hash.Available() {
		return nil, ErrHashUnavailable
	}
	keybuff, ok := key.([]byte)
	if !ok {
		return nil, ErrKeyInvalid
	}

	hasher := hmac.New(m.hash.New, keybuff)
	hasher.Write(data)
	return hasher.Sum(nil), nil
}

func (m *SigningMethodHMAC) Verify(signature, body []byte, key interface{}) (bool, error) {
	if !m.hash.Available() {
		return false, ErrHashUnavailable
	}
	keybuff, ok := key.([]byte)
	if !ok {
		return false, ErrKeyInvalid
	}

	hasher := hmac.New(m.hash.New, keybuff)
	hasher.Write(body)
	return hmac.Equal(signature, hasher.Sum(nil)), nil
}
