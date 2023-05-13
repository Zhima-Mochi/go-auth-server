package sessionmanager

import (
	"context"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"io"
	"time"

	"github.com/Zhima-Mochi/go-user-service/external/cache"
	"github.com/Zhima-Mochi/go-user-service/service/sessionManager/session"
	"github.com/google/uuid"
)

var (
	timeNow = time.Now

	defaultSessionManager = &sessionManager{
		Name:     "session_id",
		MaxAge:   3600,
		Secure:   false,
		HttpOnly: true,
	}
)

type sessionManager struct {
	Name      string
	MaxAge    int
	Secure    bool
	HttpOnly  bool
	SecretKey []byte
	cache     cache.Cache
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

func NewSessionManager(cache cache.Cache, options ...SessionManagerOption) *sessionManager {
	sm := defaultSessionManager

	sm.cache = cache

	for _, option := range options {
		option(sm)
	}

	return sm
}

// NewSession creates a new session.
func (sm *sessionManager) NewSession(ctx context.Context, data map[string]interface{}) (*session.Session, error) {
	// generate session id
	id := generateSessionID()

	// calculate session expiration time
	expiresAt := timeNow().Add(time.Duration(sm.MaxAge) * time.Second)

	// create new session
	s := session.NewSession(id, expiresAt, data)

	err := sm.cache.Set(ctx, id, s, time.Duration(sm.MaxAge)*time.Second)
	if err != nil {
		return nil, err
	}

	return s, nil
}

// GetSession gets session by id.
func (sm *sessionManager) GetSession(ctx context.Context, id string) (*session.Session, error) {
	// get session data from cache
	s, err := sm.cache.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	return s.(*session.Session), nil
}

// UpdateSession updates session by id.
func (sm *sessionManager) UpdateSession(ctx context.Context, id string, s *session.Session) error {
	err := sm.cache.Set(ctx, id, s, time.Duration(sm.MaxAge)*time.Second)
	if err != nil {
		return err
	}

	return nil
}

// DeleteSession deletes session by id.
func (sm *sessionManager) DeleteSession(ctx context.Context, id string) error {
	err := sm.cache.Delete(ctx, id)
	if err != nil {
		return err
	}

	return nil
}
