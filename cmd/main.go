package main

import (
	"log"

	"github.com/aidosgal/janabyte/janabyte-core/cmd/janabyte"
)

func main() {
    server := janabyte.NewApiServer(":8080", nil)

    if err := server.Run(); err != nil {
        log.Fatal(err)
    }
}
