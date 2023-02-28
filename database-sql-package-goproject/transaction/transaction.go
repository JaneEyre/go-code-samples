package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/go-sql-driver/mysql"
)

var db *sql.DB

type Album struct {
	ID       int64
	Title    string
	Artist   string
	Price    float32
	Quantity int64
}

type Album_trx struct {
	TRX_ID    int64
	TRX_CHECK int64
}

func main() {
	// Capture connection properties.
	cfg := mysql.Config{
		User:   os.Getenv("DBUSER"),
		Passwd: os.Getenv("DBPASS"),
		Net:    "tcp",
		Addr:   "127.0.0.1:3306",
		DBName: "recordings",
	}
	// Get a database handle.
	var err error
	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Println("Connected!")

	// Start the transaction
	ctx := context.Background()
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	// First query
	_, err = tx.ExecContext(ctx, "INSERT INTO album (title, artist, price, quantity) VALUES ('Master of Puppets', 'Metallica', '49', '1')")
	if err != nil {
		tx.Rollback()
		return
	}

	// Second query
	_, err = tx.ExecContext(ctx, "INSERT INTO album_trx (trx_check) VALUES (1)")
	if err != nil {
		tx.Rollback()
		fmt.Println("Transaction declined")
		return
	}

	// If no errors, commit the transaction
	err = tx.Commit()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Transaction accepted!")
}
