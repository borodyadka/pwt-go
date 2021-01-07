package pwt

import (
	"encoding/base64"
	"encoding/hex"
)

type Codec interface {
	Encode(input []byte) (string, error)
	Decode(input string) ([]byte, error)
}

type Base16Codec struct {}

func (Base16Codec) Encode(input []byte) (string, error) {
	return hex.EncodeToString(input), nil
}

func (Base16Codec) Decode(input string) ([]byte, error) {
	return hex.DecodeString(input)
}

type Base64Codec struct {}

func (Base64Codec) Encode(input []byte) (string, error) {
	return base64.StdEncoding.EncodeToString(input), nil
}

func (Base64Codec) Decode(input string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(input)
}

var (
	Base16Coder = &Base16Codec{}
	Base64Coder = &Base64Codec{}
)
