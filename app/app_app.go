package app

func Start(port int) {
	setupLogger()
	setupConfig()
	// db := setupDatabase()
	// go setupGrpcServer(port, db)
	// go setupSocketIOServer()
	setupSocketIOServer()
}
