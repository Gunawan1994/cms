package main

// import (
// 	"database/sql"
// 	"net/http"
// 	"os"

// 	"cms/api/handlers"
// 	"cms/pkg/db"

// 	"github.com/labstack/echo/v4"
// 	_ "github.com/lib/pq"
// 	log "github.com/sirupsen/logrus"
// )

// func initLogger() {
// 	log.SetFormatter(&log.JSONFormatter{
// 		TimestampFormat: "2006-01-02 15:04:05.999",
// 	})
// }

// func configAndStartServer() {
// 	initLogger()
// 	dbClient := dbClient()

// 	htmlEcho := setWebRouter(dbClient)
// 	start(htmlEcho)
// }

// func setWebRouter(dbClient *sql.DB) *echo.Echo {
// 	e := echo.New()

// 	root := e.Group("")
// 	root.GET("/ping", func(c echo.Context) error {
// 		return c.String(http.StatusOK, "PONG!")
// 	})

// 	// Register Routes
// 	handlers.NewRoutes(dbClient).RegisterServices(e)

// 	return e
// }

// func start(htmlEcho *echo.Echo) {
// 	// Start Run HTML Echo
// 	if err := htmlEcho.Start(os.Getenv("LISTEN_PORT")); err != nil {
// 		log.WithField("error", err).Error("Unable to start the server")
// 		os.Exit(1)
// 	}
// }

// func dbClient() *sql.DB {
// 	return db.NewConn()
// }

// func main() {
// 	initLogger()
// 	configAndStartServer()
// }
