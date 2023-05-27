package utility

type EncryptorOption func(*Encryptor)

func WithSecretKey(secretKey []byte) EncryptorOption {
	return func(e *Encryptor) {
		e.secretKey = secretKey
	}
}
