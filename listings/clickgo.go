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

	batch, err := r.conn.PrepareBatch(ctx, fmt.Sprintf("INSERT INTO %s", table))
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
