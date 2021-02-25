package crypto


import (
/*
    "crypto/aes"
    "crypto/cipher"
    "crypto/rand"
    "errors"
    "fmt"
    "io"
    "log"
*/
	"github.com/go-ozzo/ozzo-validation/v4"
    "strings"
)

func encrypt(plaintext []byte, key string) {


	err := validation.Validate(key,
		validation.Required,       // not empty
		validation.Length(5, 100), // 

	/*
    c, err := aes.NewCipher(key)
    if err != nil {
        return nil, err
    }

    gcm, err := cipher.NewGCM(c)
    if err != nil {
        return nil, err
    }

    nonce := make([]byte, gcm.NonceSize())
    if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
        return nil, err
    }

    return gcm.Seal(nonce, nonce, plaintext, nil), nil
    */
}