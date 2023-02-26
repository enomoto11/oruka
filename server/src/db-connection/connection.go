package dbconnection

import (
	"context"
	"fmt"
	"log"
	"oruka/ent"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func Connect() *ent.Client {
	path := fmt.Sprintf("%s:%s@tcp(db:3306)/%s", os.Getenv("MYSQL_USER"), os.Getenv("MYSQL_PASSWORD"), os.Getenv("MYSQL_DATABASE"))

	client, err := ent.Open("mysql", path)

	if err != nil {
		log.Fatalf("failed connect to mysql: %v", err)
	}

	defer client.Close()

	ctx := context.Background()

	if err := client.Schema.Create(ctx); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	return client
}
