package token

import (
	"fmt"
	"goparking/configs"
	"strings"

	"github.com/aead/chacha20poly1305"
	"github.com/o1egl/paseto"
)

const (
	SymmetricKey = "12345678901234567890123456789012" // 30 days
)

// PasetoMaker is a PASETO token maker
type PasetoMaker struct {
	paseto       *paseto.V2
	symmetricKey []byte
}

// NewPasetoMaker creates a new PasetoMaker
func NewPasetoMaker() (*PasetoMaker, error) {
	if len(SymmetricKey) != chacha20poly1305.KeySize {
		return nil, fmt.Errorf("invalid key size: must be exactly %d characters", chacha20poly1305.KeySize)
	}

	maker := &PasetoMaker{
		paseto:       paseto.NewV2(),
		symmetricKey: []byte(SymmetricKey),
	}

	return maker, nil
}

func (maker *PasetoMaker) GenerateAccessToken(payload *AuthPayload) string {
	config := configs.LoadConfig()
	newPayload := NewAuthPayload(payload.ID, payload.Email, payload.Role, config.Access_Token_Duration, AccessTokenType)

	token, err := maker.paseto.Encrypt(maker.symmetricKey, newPayload, nil)
	if err != nil {
		return ""
	}

	return token
}

func (maker *PasetoMaker) GenerateRefreshToken(payload *AuthPayload) string {
	config := configs.LoadConfig()
	newPayload := NewAuthPayload(payload.ID, payload.Email, payload.Role, config.Access_Token_Duration, RefreshTokenType)

	token, err := maker.paseto.Encrypt(maker.symmetricKey, newPayload, nil)
	if err != nil {
		return ""
	}

	return token
}

func (maker *PasetoMaker) ValidateToken(pasetoToken string) (*AuthPayload, error) {
	payload := &AuthPayload{}

	tokenParts := strings.Split(pasetoToken, " ")
	if len(tokenParts) == 2 && tokenParts[0] == "Bearer" {
		pasetoToken = tokenParts[1]
	}

	err := maker.paseto.Decrypt(pasetoToken, maker.symmetricKey, payload, nil)
	if err != nil {
		return nil, ErrInvalidToken
	}

	err = payload.Valid()
	if err != nil {
		return nil, err
	}

	return payload, nil
}
