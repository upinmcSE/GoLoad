package database

import (
	"context"

	"github.com/doug-martin/goqu/v9"
)

type AccountPassword struct {
	OfAccountID uint64 `db:"of_account_id" goqu:"skipupdate"`
	Hash        string `db:"hash"`
}

type AccountPasswordDataAccessor interface {
	CreateAccountPassword(ctx context.Context, accountPassword AccountPassword) error
	GetAccountPassword(ctx context.Context, ofAccountID uint64) (AccountPassword, error)
	UpdateAccountPassword(ctx context.Context, accountPassword AccountPassword) error
	WithDatabase(database Database) AccountPasswordDataAccessor
}

type accountPasswordDataAccessor struct {
	database Database
	//logger   *zap.Logger
}

// CreateAccountPassword implements AccountPasswordDataAccessor.
func (a *accountPasswordDataAccessor) CreateAccountPassword(ctx context.Context, accountPassword AccountPassword) error {
	panic("unimplemented")
}

// GetAccountPassword implements AccountPasswordDataAccessor.
func (a *accountPasswordDataAccessor) GetAccountPassword(ctx context.Context, ofAccountID uint64) (AccountPassword, error) {
	panic("unimplemented")
}

// UpdateAccountPassword implements AccountPasswordDataAccessor.
func (a *accountPasswordDataAccessor) UpdateAccountPassword(ctx context.Context, accountPassword AccountPassword) error {
	panic("unimplemented")
}

// WithDatabase implements AccountPasswordDataAccessor.
func (a *accountPasswordDataAccessor) WithDatabase(database Database) AccountPasswordDataAccessor {
	return &accountPasswordDataAccessor{
		database: database,
		//logger:   a.logger,
	}
}

func NewAccountPasswordDataAccessor(database *goqu.Database) AccountPasswordDataAccessor {
	return &accountPasswordDataAccessor{
		database: database,
		//logger:   logger,
	}
}
