package encryptor

type EncryptorOption func(*encryptor)

func WithSecretKey(secretKey []byte) EncryptorOption {
	return func(e *encryptor) {
		e.secretKey = secretKey
	}
}
