package utils

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"

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

func CheckLogin(conn *pgx.Conn, login LoginInput) int {

	var id int

	rows, err := conn.Query(context.Background(), "SELECT id FROM users WHERE user_name=$1 AND password=$2", login.UserName, login.Password)

	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	if rows.Next() == false {
		return -1
	} else {
		if err := rows.Scan(&id); err != nil {
			log.Fatal(err)
		}
		return id
	}

}
func Insert(conn *pgx.Conn, customer Customer) {
	str := "INSERT INTO users(name,surname,password,user_name,phone_number)   VALUES ($1,$2,$3,$4,$5);"
	_, err := conn.Exec(context.Background(), str, customer.Name, customer.Surname, customer.Password, customer.UserName, customer.PhoneNumber)
	if err != nil {
		log.Fatal(err)
	}
}
func query(conn *pgx.Conn) {
	rows, err := conn.Query(context.Background(), "SELECT * FROM users WHERE id=10")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		var id int32
		var name string
		if err := rows.Scan(&id, &name); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%d | %s\n", id, name)
	}
}

func InsertBill(conn *pgx.Conn, bill Bill) int {
	typeId, _ := strconv.Atoi(bill.TypeName)

	str := "INSERT INTO bills(user_id,type_id,bill_name,year,month,amount)   VALUES ($1,$2,$3,$4,$5,$6);"
	_, err := conn.Exec(context.Background(), str, bill.UserId, typeId, bill.BillName, bill.Year, bill.Month, bill.Amount)
	if err != nil {
		log.Fatal(err)
		return 0
	}
	return 1
}
