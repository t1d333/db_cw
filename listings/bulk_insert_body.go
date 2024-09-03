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
