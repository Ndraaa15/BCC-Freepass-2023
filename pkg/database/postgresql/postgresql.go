package mongodb

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func InitPostgreSQL() error {
	ctx := context.Background()

	conn, err := pgx.Connect(ctx, "user=pqgotest dbname=pqgotest sslmode=verify-full")
	if err != nil {
		return err
	}
	defer conn.Close(ctx)

	return nil
}
