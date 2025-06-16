package logic

import (
	"context"
	"errors"

	"github.com/upinmcSE/GoLoad/internal/configs"
	"golang.org/x/crypto/bcrypt"
)

type Hash interface {
	Hash(ctx context.Context, data string) (string, error)
	IsHashEqual(ctx context.Context, data string, hashed string) (bool, error)
}

type hash struct {
	accountConfig configs.Account
}

// Hash implements Hash.
func (h *hash) Hash(ctx context.Context, data string) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(data), h.accountConfig.HashCost)
	if err != nil {
		// logger.With(zap.Error(err)).Error("failed to hash data")
		return "", err // status.Error(codes.Internal, "failed to hash data")
	}

	return string(hashed), nil
}

// IsHashEqual implements Hash.
func (h *hash) IsHashEqual(ctx context.Context, data string, hashed string) (bool, error) {
	if err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(data)); err != nil {
		// logger.With(zap.Error(err)).Error("failed to compare hash and password")
		if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
			return false, nil // status.Error(codes.Unauthenticated, "password does not match")
		}
		return false, err // status.Error(codes.Internal, "failed to compare hash and password")
	}
	return true, nil
}

func NewHash(accountConfig configs.Account) Hash {
	return &hash{
		accountConfig: accountConfig,
	}
}
