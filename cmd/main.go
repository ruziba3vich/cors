package main

import (
	"log"
	"os"

	"github.com/ruziba3vich/cors/internal/http"
)

func main() {
	logger := log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	logger.Fatal(http.Run(logger))
}
