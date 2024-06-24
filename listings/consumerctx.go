func (s *ServiceImpl) StartConsumer() {
	for {

		groups := make([]string, len(s.groups))
		i := 0

		for key := range s.groups {
			groups[i] = key
			i += 1
		}

		s.log.Infow("starting consumer...", "topics", groups)

		if err := s.client.Consume(s.ctx,
												groups, s.consumer); err != nil {
			if errors.Is(err,
					sarama.ErrClosedConsumerGroup) {
				return
			}
			s.log.Fatalf("Error from consumer: %v", err)
		}

		if s.ctx.Err() != nil {
			s.log.Errorw("got consumer error",
				"err", s.ctx.Err())
			return
		}

		s.consumer.Ready = make(chan bool)
	}
}
