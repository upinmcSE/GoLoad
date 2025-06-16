package database

import (
	"context"

	"github.com/doug-martin/goqu/v9"
)

var (
	TabNameAccounts = goqu.T("accounts")

	//ErrAccountNotFound = status.Error(codes.NotFound, "account not found")
)

const (
	ColNameAccountsID          = "id"
	ColNameAccountsAccountName = "account_name"
)

type Account struct {
	ID          uint64 `db:"id" goqu:"skipinsert,skipupdate"`
	AccountName string `db:"account_name"`
}

type AccountDataAccessor interface {
	CreateAccount(ctx context.Context, account Account) (uint64, error)
	GetAccountByID(ctx context.Context, id uint64) (Account, error)
	GetAccountByAccountName(ctx context.Context, accountName string) (Account, error)
	WithDatabase(database Database) AccountDataAccessor
}

type accountDataAccessor struct {
	database Database
}

// WithDatabase implements AccountDataAccessor.
func (a *accountDataAccessor) WithDatabase(database Database) AccountDataAccessor {
	return &accountDataAccessor{
		database: database,
	}
}

// CreateAccount implements AccountDataAccessor.
func (a *accountDataAccessor) CreateAccount(ctx context.Context, account Account) (uint64, error) {
	result, err := a.database.
		Insert(TabNameAccounts).
		Rows(goqu.Record{
			ColNameAccountsAccountName: account.AccountName,
		}).
		Executor().
		ExecContext(ctx)
	if err != nil {
		//logger.With(zap.Error(err)).Error("failed to create account")
		return 0, nil //status.Error(codes.Internal, "failed to create account")
	}

	lastInsertedID, err := result.LastInsertId()
	if err != nil {
		//logger.With(zap.Error(err)).Error("failed to get last inserted id")
		return 0, nil //status.Error(codes.Internal, "failed to get last inserted id")
	}

	return uint64(lastInsertedID), nil

}

// GetAccountByAccountName implements AccountDataAccessor.
func (a *accountDataAccessor) GetAccountByAccountName(ctx context.Context, accountName string) (Account, error) {
	panic("unimplemented")
}

// GetAccountByID implements AccountDataAccessor.
func (a *accountDataAccessor) GetAccountByID(ctx context.Context, id uint64) (Account, error) {
	panic("unimplemented")
}

func NewAccountDataAccessor(database *goqu.Database) AccountDataAccessor {
	return &accountDataAccessor{
		database: database,
	}
}
