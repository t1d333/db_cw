type Repository struct {
	conn driver.Conn
	log  logger.Logger
}
func NewRepository(conn driver.Conn, log logger.Logger) repository.Repository {
	if err := conn.Exec(context.Background(), `SET allow_experimental_object_type = 1;`); err != nil {
	}
	return &Repository{
		conn: conn,
		log:  log,
	}
}
