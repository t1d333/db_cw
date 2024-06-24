go func() {
	for {
		select {
		case <-ticker.C:
			consumer.storage.Range(
				func(key, value any) bool {
				logs := value.([]string)
				if len(logs) == 0 {
					return true
				}
				err := consumer.repo.BulkInsert(
					context.Background(),
					key.(string),
					logs,
				)
				if err != nil {
					consumer.logger.Errorw(
						"failing to insert logs",
						"error", err)
				} else {
					consumer.storage.Store(key, []string{})
				}
				return true
			})
		}
	}
}()
