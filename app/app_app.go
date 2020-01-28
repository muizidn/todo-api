package app

func Start(port int) {
	setupLogger()
	setupConfig()
	db := setupDatabase()
	setupGrpcServer(port, db)
}
