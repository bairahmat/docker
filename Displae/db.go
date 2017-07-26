package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func db_start() {
	db, err := sql.Open("postgres", "postgresql://root@localhost:26257/db_displae?sslmode=disable")
	if err != nil {
		log.Fatal("error connecting to the database: ", err)
	}
	if _, err := db.Exec("CREATE TABLE IF NOT EXISTS shipper (id serial,name string unique,logo string not null default '', primary key(id))"); err != nil {
		log.Fatal(err)
	}
	if _, err := db.Exec("create table IF NOT EXISTS country (id serial,code string unique,name string unique,primary key(id))"); err != nil {
		log.Fatal(err)
	}
	if _, err := db.Exec("CREATE TABLE IF NOT EXISTS bank (id serial,name string unique,bank_code string unique ,swift_code string unique,country string not null references country(id),index (country),primary key(id))"); err != nil {
		log.Fatal(err)
	}

}

func main() {
	db_start()
}
