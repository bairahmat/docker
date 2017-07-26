package main

import (
	"database/sql"
	"fmt"
	"log"
	"math/rand"
	"time"

	_ "github.com/lib/pq"
)

func start() {
	fmt.Println("Silahkan menikmati layanan kami, Anda sedangan memikirkan apa ?")
}

// func connect() *sql.DB {
// 	var db, err = sql.Open("postgres", "postgresql://root@localhost:26257/db_displae?sslmode=disable") //default cockroach setting
// 	err = db.Ping()
// 	if err != nil {
// 		fmt.Println("database tidak bisa dihubungi")
// 		os.Exit(0)
//
// 	}
// 	return db
// }

func main() {
	var phone_number, password string
	db, err := sql.Open("postgres", "postgresql://root@localhost:26257/db_displae?sslmode=disable")
	if err != nil {
		log.Fatal("error connecting to the database: ", err)
	}
	// Create tabel user
	if _, err := db.Exec(
		"CREATE TABLE IF NOT EXISTS pengguna (id SERIAL PRIMARY KEY, full_name VARCHAR(40), phone_number VARCHAR(13), id_profile VARCHAR NULL, password VARCHAR, invitation VARCHAR, otp_conf VARCHAR)"); err != nil {
		log.Fatal(err)
	}
	// if _, err := db.Exec(
	// 	"INSERT INTO pengguna (full_name, phone_number, id_profile, password, invitation, otp_conf) VALUES ('Jihar Al Gifari','082325600996','j1','123','12','3' )"); err != nil {
	// 	log.Fatal(err)
	// }
	//Menampilkan data pengguna
	rows, err := db.Query("SELECT phone_number, password FROM pengguna")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		if err := rows.Scan(&phone_number, &password); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%s %s \n", phone_number, password)
	}
	fmt.Printf("%s %s \n", phone_number, password)
	phone := "0823256009946"
	pass := "123456"
	waktu := rand.NewSource(time.Now().UnixNano())
	jihar := rand.New(waktu)
	fmt.Println(jihar.Intn(100))
	if phone == phone_number {
		if pass == password {
			start()
		} else {
			fmt.Println("Password anda salah")
		}
	} else {
		fmt.Println("Username anda belum terdaftar")
	}
}
