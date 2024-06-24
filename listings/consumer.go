type Consumer struct {
	Ready     chan bool
	repo      repository.Repository
	storage   sync.Map
	batchSize int
	log       logger.Logger
}

func NewConsumer(ready chan bool, repo repository.Repository, log logger.Logger) *Consumer {
	return &Consumer{
		Ready:     ready,
		repo:      repo,
		storage:   sync.Map{},
		batchSize: 50,
		log:       log,
	}
}

func (consumer *Consumer) Setup(sarama.ConsumerGroupSession) error {
	close(consumer.Ready)
	return nil
}

func (*Consumer) Cleanup(sarama.ConsumerGroupSession) error {
	return nil
}

