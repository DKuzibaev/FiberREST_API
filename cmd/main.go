package main

import "fiberrest_api/internal/server"

func main() {
	app := server.NewApp("tasks")
	server.SetupRoutes("tasks", app)
	app.Listen(":8080")
}
