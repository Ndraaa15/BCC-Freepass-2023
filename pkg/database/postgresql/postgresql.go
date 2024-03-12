package mongodb

import (
	"bcc-freepass-2023/pkg/config"
	"context"

	"github.com/jackc/pgx/v5"
)

func InitPostgreSQL() error {
	ctx := context.Background()

	conn, err := pgx.Connect(ctx, config.LoadConfigPostgresql())
	if err != nil {
		return err
	}
	defer conn.Close(ctx)

	return nil
}
