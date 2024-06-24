func main() {
	log := logger.NewLogger()
	conn, err := repository.CreateClickhouseConnection(
				"clickhouse", 
				"default", "",
				"default", 9000)
	if err != nil {
		log.Fatal("failed to create clickhouse connection", err)
	}
	r := gin.Default()
	v1 := r.Group("/api/v1")
	wg := &sync.WaitGroup{}
	repo := repository.NewRepository(conn)
	ctx, cancel := context.WithCancel(context.Background())
	consumer := consumer.NewConsumer(make(chan bool), repo, log)
	service := service.NewService(
			ctx,
			cancel, wg,
			log, consumer,
			[]string{"kafka:9092"}
		)

	delivery := http.NewDelivery(v1, log, service)
	delivery.RegisterHandlers()
	r.GET("/swagger/*any",
		ginSwagger.WrapHandler(
			swaggerFiles.Handler))

	wg.Add(1)
	go func() {
		defer wg.Done()
		log.Info("starting gin on port 8080")
		if err := r.Run(":8080"); err != nil {
			log.Fatal("failed to run gin app", err)
		}
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		service.StartConsumer()
	}()
	<-consumer.Ready
	log.Info("Sarama consumer up and running!...")
	wg.Wait()
}
