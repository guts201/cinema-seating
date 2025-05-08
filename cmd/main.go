package main

import (
	config "cinema/pkg/carbon"

	_ "github.com/go-sql-driver/mysql"

	"cinema/internal/server"
)

func main() {
	flags := config.ParseFlags()
	server.Run(flags)
}
