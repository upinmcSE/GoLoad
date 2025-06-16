package logic

import (
	"context"
	"database/sql"
	"errors"

	"github.com/doug-martin/goqu/v9"
	"github.com/upinmcSE/GoLoad/internal/dataaccess/database"
)

type CreateAccountParams struct {
	AccountName string
	Password string
}

type CreateAccountOutput struct {
	ID          uint64
	AccountName string
}

type CreateSessionParams struct {
	AccountName string
	Password    string
}

type CreateSessionOutput struct {
	//Account *go_load.Account
	Token   string
}


type Account interface {
	CreateAccount(ctx context.Context, params CreateAccountParams) (CreateAccountOutput, error)
	CreateSession(ctx context.Context, params CreateSessionParams) (CreateSessionOutput, error)
}

type account struct {
	goquDatabase *goqu.Database
	accountDataAccessor database.AccountDataAccessor
	accountPasswordDataAccessor database.AccountPasswordDataAccessor
	hashLogic Hash
}

func (a account) isAccountAccountNameTaken(ctx context.Context, accountName string) (bool, error) {
	//logger := utils.LoggerWithContext(ctx, a.logger).With(zap.String("account_name", accountName))

	// accountNameTaken, err := a.takenAccountNameCache.Has(ctx, accountName)
	// if err != nil {
	// 	logger.With(zap.Error(err)).Warn("failed to get account name from taken set in cache, will fall back to database")
	// } else if accountNameTaken {
	// 	return true, nil
	// }

	_, err := a.accountDataAccessor.GetAccountByAccountName(ctx, accountName)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return false, nil
		}

		return false, err
	}

	

	// err = a.takenAccountNameCache.Add(ctx, accountName)
	// if err != nil {
	// 	logger.With(zap.Error(err)).Warn("failed to set account name into taken set in cache")
	// }

	return true, nil
}

func (a *account) CreateAccount(ctx context.Context, params CreateAccountParams) (CreateAccountOutput, error) {
	var accountId uint64

	txErr := a.goquDatabase.WithTx(func(td *goqu.TxDatabase) error {
		usernameTaken, err := a.isAccountAccountNameTaken(ctx, params.AccountName)
		if err != nil {
			return  err // status.Error(codes.Internal, "failed to check if account name is taken")
		}

		if usernameTaken {
			return  errors.New("account name is already taken") // status.Error(codes.AlreadyExists, "account name is already taken")
		}

		accountId, err := a.accountDataAccessor.WithDatabase(td).CreateAccount(ctx, database.Account{
			AccountName: params.AccountName,
		}); 

		if err != nil {
			return err // status.Error(codes.Internal, "failed to create account")
		}

		hashedPassword, err := a.hashLogic.Hash(ctx, params.Password)
		if err != nil {
			return err // status.Error(codes.Internal, "failed to hash password")
		}

		if err := a.accountPasswordDataAccessor.WithDatabase(td).CreateAccountPassword(ctx, database.AccountPassword{
			OfAccountID: accountId,
			Hash: hashedPassword,
		}); err != nil {
			return err // status.Error(codes.Internal, "failed to create account password")
		}
		return nil
	})

	if txErr != nil {
		return CreateAccountOutput{}, txErr // status.Error(codes.Internal, "failed to create account transaction")
	}

	return CreateAccountOutput{
			ID:       accountId,
			AccountName: params.AccountName,
		}, nil
}


func (a *account) CreateSession(ctx context.Context, params CreateSessionParams) (CreateSessionOutput, error) {
	panic("unimplemented")
}

func NewAccount(
	goquDatabase *goqu.Database,
	accountDataAccessor database.AccountDataAccessor,
	accountPasswordDataAccessor database.AccountPasswordDataAccessor,
	hashLogic Hash,
) Account {
	return &account{
		goquDatabase: goquDatabase,
		accountDataAccessor: accountDataAccessor,
		accountPasswordDataAccessor: accountPasswordDataAccessor,
		hashLogic: hashLogic,
	}
}
