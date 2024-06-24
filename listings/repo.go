package clickhouse

import (
	"context"
	"fmt"
	"net"
	"time"

	"github.com/ClickHouse/clickhouse-go/v2"
	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"
	"github.com/t1d333/log_collector/internal/repository"
)

type Repository struct {
	conn driver.Conn
}

func NewRepository(conn driver.Conn) repository.Repository {
	if err := conn.Exec(context.Background(), `SET allow_experimental_object_type = 1;`); err != nil {
	}

	return &Repository{
		conn: conn,
	}
}

func (r *Repository) BulkInsert(ctx context.Context,
				table string, logs []string) error {
	if err := r.conn.Exec(ctx, fmt.Sprintf(`
		CREATE TABLE IF NOT EXISTS %s (
			timestamp DateTime64 DEFAULT current_timestamp(),
			payload JSON
		) ENGINE = MergeTree() ORDER BY timestamp;
		`, table)); err != nil {
		return err
	}

	batch, err := r.conn.PrepareBatch(ctx,
	fmt.Sprintf("INSERT INTO %s", table))
	if err != nil {
		return err
	}

	for _, log := range logs {
		if err := batch.Append(time.Now(), log); err != nil {
			return err
		}
	}

	return batch.Send()
}

func (r *Repository) Insert(ctx context.Context, table string, log string) error {
	if err := r.conn.Exec(ctx, fmt.Sprintf(`
		CREATE TABLE IF NOT EXISTS %s (
			timestamp DateTime64 DEFAULT current_timestamp(),
			payload JSON
		) ENGINE = MergeTree() ORDER BY timestamp;
		`, table)); err != nil {
		return err
	}

	if err := r.conn.Exec(ctx, fmt.Sprintf("INSERT INTO %s FORMAT JSONEachRow {\"payload\": %s}", table, log)); err != nil {
		return fmt.Errorf("Repository.Insert(db: %s, log: %s): %w", table, log, err)
	}

	return nil
}
