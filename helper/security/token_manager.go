package security

import (
	"bytes"
	"context"
	"encoding/gob"
	"ikomers-be/config"
	"ikomers-be/model"
	"ikomers-be/model/helper"

	jwt "github.com/square/go-jose"
)

var (
	accessTokenKey  = []byte(config.GetConfig().AccessTokenKey)
	refreshTokenKey = []byte(config.GetConfig().RefreshTokenKey)
)

type jwtManager struct{}

func NewTokenManager() helper.TokenManager {
	return &jwtManager{}
}

func (m *jwtManager) GenerateAccessToken(ctx context.Context, user model.User) (string, error) {
	signer, err := jwt.NewSigner(jwt.SigningKey{Algorithm: jwt.HS256, Key: accessTokenKey}, nil)
	if err != nil {
		return "", err
	}

	token, err := m.Encode(signer, user)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (m *jwtManager) GenerateRefreshToken(ctx context.Context, user model.User) (string, error) {
	signer, err := jwt.NewSigner(jwt.SigningKey{Algorithm: jwt.HS256, Key: refreshTokenKey}, nil)
	if err != nil {
		return "", err
	}

	token, err := m.Encode(signer, user)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (m *jwtManager) Encode(signer jwt.Signer, o interface{}) (string, error) {
	var buf bytes.Buffer
	err := gob.NewEncoder(&buf).Encode(o)
	if err != nil {
		return "", err
	}

	obj, err := signer.Sign(buf.Bytes())
	if err != nil {
		return "", err
	}

	str, err := obj.CompactSerialize()
	if err != nil {
		return "", err
	}

	return str, nil
}
