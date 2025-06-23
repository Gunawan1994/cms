package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strings"
	"sync"

	_ "github.com/lib/pq" // PostgreSQL driver
)

var (
	dbSql *sql.DB
	once  sync.Once
)

func NewConn() *sql.DB {
	var err error
	conn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("POSTGRES_ADDR"),
		os.Getenv("POSTGRES_PORT"),
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_DB"))

	once.Do(func() {
		dbSql, err = sql.Open("postgres", conn)
		if err != nil {
			panic(err)
		}
	})

	restore(dbSql)

	return dbSql
}

func restore(db *sql.DB) {
	sqlDump := `-- Enable the UUID extension in PostgreSQL
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- SQL dump for creating the "user" table in PostgreSQL
CREATE TABLE IF NOT EXISTS "user" (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    first_name VARCHAR(50) NOT NULL,
    last_name VARCHAR(50) NOT NULL,
    phone_number VARCHAR(15) NOT NULL,
    address VARCHAR(255) NOT NULL,
    pin VARCHAR(10) NOT NULL,
    created_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- SQL dump for creating the "topup" table in PostgreSQL
CREATE TABLE IF NOT EXISTS "topup" (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    amount NUMERIC(10, 2) NOT NULL,
    balance_before NUMERIC(10, 2) NOT NULL,
    balance_after NUMERIC(10, 2) NOT NULL,
    created_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- SQL dump for creating the "payment" table in PostgreSQL
CREATE TABLE IF NOT EXISTS "payment" (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    amount NUMERIC(10, 2) NOT NULL,
    remarks VARCHAR(255),
    balance_before NUMERIC(10, 2) NOT NULL,
    balance_after NUMERIC(10, 2) NOT NULL,
    created_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Example inserts into the "payment" table
INSERT INTO "payment" (id, amount, remarks, balance_before, balance_after) VALUES
(uuid_generate_v4(), 50.00, 'Payment for services', 150.00, 100.00);

-- Example inserts into the "topup" table
INSERT INTO "topup" (id, amount, balance_before, balance_after) VALUES
(uuid_generate_v4(), 100.00, 50.00, 150.00);

-- Example inserts into the "topup" table
INSERT INTO "topup" (id, amount, balance_before, balance_after) VALUES
(uuid_generate_v4(), 100.00, 50.00, 150.00);
`

	data := strings.Split(sqlDump, ";")
	for _, stmt := range data {
		stmt = strings.TrimSpace(stmt)
		if stmt == "" {
			continue
		}
		_, err := db.Exec(stmt)
		if err != nil {
			log.Fatal(err)
		}
	}

	fmt.Println("Schema SQL executed successfully")
}
