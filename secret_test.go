package secret

import (
	"log"
	"testing"
)

func TestSecret(t *testing.T) {
	key := "1234567890123456"
	text := "some text"
	hex := Encrypt(key, text)
	text2 := Decrypt(key, hex)
	log.Println("input:", text)
	log.Println("hex:", hex)
	log.Println("decrypt:", text2)

}
