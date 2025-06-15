package database

type AccountPassword struct {
	OfAccountID uint64 `db:"of_account_id" goqu:"skipupdate"`
	Hash        string `db:"hash"`
}