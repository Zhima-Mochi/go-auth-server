package sessionmanager

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"io"

	"github.com/Zhima-Mochi/go-user-service/internal/sessionManager/session"
	"github.com/google/uuid"
)

type sessionManager struct {
	Name      string
	MaxAge    int
	Secure    bool
	HttpOnly  bool
	SecretKey []byte
}

// generateSessionID generates a new session id.
func generateSessionID() string {
	return uuid.New().String()
}

// encrypt encrypts data with key.
func encrypt(data, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
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

// encodeSessionData encodes session data with secret key.
func encodeSessionData(session *session.Session, secretKey []byte) (string, error) {
	// serialize session data to JSON
	data, err := json.Marshal(session)
	if err != nil {
		return "", err
	}

	// encrypt session data with secret key
	encryptedData, err := encrypt(data, secretKey)
	if err != nil {
		return "", err
	}

	// encode encrypted data as base64 string
	encodedData := base64.StdEncoding.EncodeToString(encryptedData)

	return encodedData, nil
}
