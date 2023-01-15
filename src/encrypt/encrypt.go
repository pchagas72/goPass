package encrypt

import (
	"crypto/aes"

	"github.com/pchagas72/goPass/src/base"
)

func Encrypt(password string, key string) []byte {
	c, err := aes.NewCipher([]byte(key))
	base.Check(err)
	var encrypted_bytes []byte = make([]byte, len(password))
	c.Encrypt(encrypted_bytes, []byte(password))
	return encrypted_bytes
}

func Decrypt(cipher_text string, key string) string {
	c, err := aes.NewCipher([]byte(key))
	base.Check(err)
	var decrypted_bytes []byte = make([]byte, len(cipher_text))
	c.Decrypt(decrypted_bytes, []byte(cipher_text))
	return string(decrypted_bytes)
}
