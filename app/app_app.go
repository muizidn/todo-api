package app

func Start(port int) {
	setupConfig()
	setupLogger()
	db := setupDatabase()
	setupGrpcServer(port, db)
}
