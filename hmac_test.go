package pwt

import (
	"crypto"
	"testing"
)

var (
	key = []byte("test-key")
	hmac256 = &SigningMethodHMAC{hash: crypto.SHA256}
)

func TestSigningMethodHMAC_Sign(t *testing.T) {
	data := mustMarshalToken(t, stubToken())
	sign, err := hmac256.Sign(data, key)
	handleError(t, err)
	t.Fatal(toBase64(sign))
}

func TestSigningMethodHMAC_Verify(t *testing.T) {

}
