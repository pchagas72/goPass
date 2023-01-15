package main

import (
	"github.com/pchagas72/goPass/src/base"
	"github.com/pchagas72/goPass/src/encode"
	"github.com/pchagas72/goPass/src/encrypt"
)

func main() {
	var password string = base.GeneratePassword(16, 1, 1, 1)
	println(password)
	var key string = base.CreateKey("senha super segura")
	println(key)
	println("key OK")
	println(len(key))
	cipher := encrypt.Encrypt(password, key)
	println("encrypt OK")
	cipher_encoded := encode.EncodeBase64(string(cipher))
	println("encode OK")
	println(cipher_encoded)
	println("Entering decoding tests...")
	cipher_decoded := encode.DecodeBase64(cipher_encoded)
	println("decode OK")
	cipher_decrypted := encrypt.Decrypt(cipher_decoded, key)
	println("Decrypt OK")
	println(cipher_decrypted)

}
