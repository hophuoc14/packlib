package main

func (app *Application) routes() {
	app.server.GET("/health", app.handler.HealthCheck)
}