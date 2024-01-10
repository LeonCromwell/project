package auth

import (
	"crypto/rand"
	"math/big"
)

func CreateVerificationCode() (string, error) {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	codeLength := 6

	randomCode := make([]byte, codeLength)
	for i := range randomCode {
		idx, err := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		if err != nil {
			return "", err
		}
		randomCode[i] = charset[idx.Int64()]
	}

	return string(randomCode), nil
}

func VerifyCode(code, expectedCode string) bool {
	return code == expectedCode
}
