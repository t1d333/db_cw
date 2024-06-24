	for {
		select {
		case message, ok := <-claim.Messages():
			if !ok {
				consumer.log.Info(
					"message channel was closed")
				return nil
			}

			val, ok := consumer
									.storage
									.Load(message.Topic)

			logs := []string{}

			if !ok {
				consumer.storage.Store(
					message.Topic,
					[]string{string(message.Value)})
			} else {
				logs = append(
						val.([]string),
						string(message.Value))
				consumer.storage.Store(message.Topic, logs)
			}

			if len(logs) == consumer.batchSize {
				err := consumer.repo.BulkInsert(
					context.Background(),
					message.Topic,
					logs,
				)
				if err != nil {
					return err
				}
				consumer.storage.Store(
					message.Topic,
					[]string{})
			}
			session.MarkMessage(message, "")
		case <-session.Context().Done():
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
				}
				return true
			})
			return nil
		}
	}
