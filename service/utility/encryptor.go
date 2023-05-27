package utility

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"io"

	"github.com/Zhima-Mochi/go-authentication-service/external"
)

var _ external.Encryptor = (*Encryptor)(nil)

type Encryptor struct {
	secretKey []byte
}

// NewEncryptor creates a new Encryptor.
func NewEncryptor(options ...EncryptorOption) *Encryptor {
	e := &Encryptor{
		secretKey: generateSecretKey(),
	}

	for _, option := range options {
		option(e)
	}

	return e
}

// generateSecretKey generates a new secret key.
func generateSecretKey() []byte {
	key := make([]byte, 32)
	if _, err := io.ReadFull(rand.Reader, key); err != nil {
		panic(err)
	}
	return key
}

// Encrypt encrypts data with secret key.
func (e *Encryptor) Encrypt(data []byte) ([]byte, error) {
	block, err := aes.NewCipher(e.secretKey)
	if err != nil {
		return nil, err
	}

	ciphertext := make([]byte, aes.BlockSize+len(data))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return nil, err
	}

	stream := cipher.NewCTR(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], data)

	return ciphertext, nil
}

// Decrypt decrypts data with secret key.
func (e *Encryptor) Decrypt(data []byte) ([]byte, error) {
	block, err := aes.NewCipher(e.secretKey)
	if err != nil {
		return nil, err
	}

	iv := data[:aes.BlockSize]
	data = data[aes.BlockSize:]

	stream := cipher.NewCTR(block, iv)
	stream.XORKeyStream(data, data)

	return data, nil
}
