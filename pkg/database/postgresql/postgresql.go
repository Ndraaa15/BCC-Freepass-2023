package postgresql

import (
	"bcc-freepass-2023/pkg/config"
	"context"

	"github.com/jackc/pgx/v5"
)

func InitPostgreSQL() (*pgx.Conn, error) {
	ctx := context.Background()

	dbConn, err := pgx.Connect(ctx, config.LoadConfigPostgresql())
	if err != nil {
		return nil, err
	}
	defer dbConn.Close(ctx)

	return dbConn, nil
}
