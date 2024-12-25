package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
)

type User struct {
	name string
	id   int
}

func database() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	connStr := os.Getenv("DATABASE_URL")
	conn, err := pgx.Connect(context.Background(), connStr)
	if err != nil {
		panic(err)
	}
	defer conn.Close(context.Background())
	createTable(conn)
	// userInfo := User{
	// 	name: "Furkan",
	// 	id:   10,
	// }
	// insert(conn, userInfo)
	query(conn)
}
func createTable(conn *pgx.Conn) {
	_, err := conn.Exec(context.Background(), "CREATE TABLE IF NOT EXISTS Users(id SERIAL PRIMARY KEY, name TEXT);")
	if err != nil {
		panic(err)
	}
}
func insert(conn *pgx.Conn, userInfo User) {
	name := userInfo.name
	id := userInfo.id
	_, err := conn.Exec(context.Background(), "INSERT INTO Users(id,name)   VALUES ($1,$2);", id, name)
	if err != nil {
		panic(err)
	}
}
func query(conn *pgx.Conn) {
	rows, err := conn.Query(context.Background(), "SELECT * FROM users WHERE id=10")
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		var id int32
		var name string
		if err := rows.Scan(&id, &name); err != nil {
			panic(err)
		}
		fmt.Printf("%d | %s\n", id, name)
	}
}