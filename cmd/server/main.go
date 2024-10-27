package main

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/Yadier01/golangMovie/internal/server"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

func main() {
	//viper config
	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Printf("Error reading .env file: %s \n", err)
	}
	connStr := viper.GetString("connStr")

	conn, err := sql.Open("postgres", connStr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to open db %s: %s", connStr, err)
		os.Exit(1)
	}

	defer conn.Close()

	srv := server.NewServer(conn)
	srv.New()
}
