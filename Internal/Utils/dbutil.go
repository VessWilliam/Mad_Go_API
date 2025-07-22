package utils

import (
	"context"

	"github.com/jmoiron/sqlx"
)

func NameGetContext(ctx context.Context, db *sqlx.DB, query string, dest any, arg any) error {
	param, err := db.PrepareNamedContext(ctx, query)
	if err != nil {
		return err
	}

	defer param.Close()
	return param.GetContext(ctx, dest, arg)
}
