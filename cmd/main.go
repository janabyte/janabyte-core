package main

import (
	"github.com/aidosgal/janabyte/janabyte-core/cmd/janabyte"
	"github.com/aidosgal/janabyte/janabyte-core/internal/logger"
	"log"
	"os"

	"github.com/aidosgal/janabyte/janabyte-core/config"
	"github.com/aidosgal/janabyte/janabyte-core/database"
	"github.com/fatih/color"
)

func main() {
	color.NoColor = false
	sloger := logger.SetupLogger()
	cfg := config.InitConfig()

	storage, err := database.New(cfg)
	if err != nil {
		sloger.Error("Error initializing database", logger.Err(err))
		os.Exit(1)
	}
	sloger.Info("Database created")

	server := janabyte.NewApiServer(":8090", storage.DB)

	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
