package main

import (
	"log"

	"github.com/aidosgal/janabyte/janabyte-core/cmd/janabyte"
	"github.com/aidosgal/janabyte/janabyte-core/config"
	"github.com/aidosgal/janabyte/janabyte-core/database"
	"github.com/go-sql-driver/mysql"
)

func main() {
    db, err := database.NewMySQLStorage(mysql.Config{
        User:                   config.Envs.DBUser,
        Passwd:                 config.Envs.DBPassword,
        Addr:                   config.Envs.DBAddress,
        DBName:                 config.Envs.DBName,
        Net:                    "tcp",
        AllowNativePasswords:   true,
        ParseTime:              true,
    })

    if err != nil {
        log.Fatal(err)
    }

    server := janabyte.NewApiServer(":8080", db)

    if err := server.Run(); err != nil {
        log.Fatal(err)
    }
}
