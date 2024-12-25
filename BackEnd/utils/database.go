package utils

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
)

func DbConnect() *pgx.Conn {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	connStr := os.Getenv("DATABASE_URL")
	conn, err := pgx.Connect(context.Background(), connStr)
	if err != nil {
		panic(err)
	}
	return conn
}
func CheckUser(conn *pgx.Conn, userName string) int {

	rows, err := conn.Query(context.Background(), "SELECT id,name FROM users WHERE user_name=$1", userName)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	if rows.Next() == false {
		return 1
	} else {
		return 0
	}

	// for rows.Next() {
	// 	var id int32
	// 	var name string
	// 	if err := rows.Scan(&id, &name); err != nil {
	// 		panic(err)
	// 	}
	// 	fmt.Printf("%d | %s\n", id, name)
	// }

}

//	func insert(conn *pgx.Conn, userInfo User) {
//		name := userInfo.name
//		id := userInfo.id
//		_, err := conn.Exec(context.Background(), "INSERT INTO Users(id,name)   VALUES ($1,$2);", id, name)
//		if err != nil {
//			panic(err)
//		}
//	}
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
