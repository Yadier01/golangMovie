package main

import (
	"database/sql"
	"fmt"
	"log"

	internal "github.com/Yadier01/golangMovie/internal/routes"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Printf("Error reading .env file: %s \n", err)
	}
	connStr := viper.GetString("connStr")

	conn, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("failed to open db %s: %s", connStr, err)
	}

	defer conn.Close()

	srv := internal.NewServer(conn)
	srv.New()
}
