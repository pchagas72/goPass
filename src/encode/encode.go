package encode

import (
	"encoding/base64"

	"github.com/pchagas72/goPass/src/base"
)

func EncodeBase64(password string) string {
	var encoded string = base64.RawStdEncoding.EncodeToString([]byte(password))
	return encoded
}

func DecodeBase64(encoded_password string) string {
	decoded, error := base64.RawStdEncoding.DecodeString(encoded_password)
	base.Check(error)
	return string(decoded)
}
