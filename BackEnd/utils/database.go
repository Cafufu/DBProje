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

// this function is used to check if any tuple exist with the given userId in the carbon_footprint table in database

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
func CheckBill(conn *pgx.Conn, myBill Bill) bool {
	var exist bool
	query := "SELECT EXISTS ( SELECT 1 FROM bills WHERE user_id = $1 AND type_id = $2 AND bill_name = $3 AND year = $4 AND month = $5);"

	err := conn.QueryRow(context.Background(), query, myBill.UserId, myBill.TypeName, myBill.BillName, myBill.Year, myBill.Month).Scan(&exist)

	if err != nil {
		log.Fatal(err)
	}

	return exist
}
func UpdateBill(conn *pgx.Conn, myBill Bill) int {
	query := "UPDATE bills SET amount = $1 WHERE user_id = $1 AND type_id = $2 AND bill_name = $3 AND year = $4 AND month = $5;"
	_, err := conn.Exec(context.Background(), query, myBill.UserId, myBill.TypeName, myBill.BillName, myBill.Year, myBill.Month)
	if err != nil {
		log.Fatal(err)
		return 0
	}
	return 1

}
func CheckCarbonFootprint(conn *pgx.Conn, userId int) bool {
	var exist bool
	query := "Select carbon_exist($1)"
	ctx := context.Background()
	err := conn.QueryRow(ctx, query, userId).Scan(&exist)
	if err != nil {
		log.Fatal(err)
	}

	return exist
}
func InsertCarbonFootprint(conn *pgx.Conn, userId int) int {
	value := "0" //eğer kullanıcı ilk defa fatura ekliyor ise karbon ayak izi 0 olarak başlatılır daha sonra update edilir.
	str := "INSERT INTO carbon_footprint(user_id,value)   VALUES ($1,$2);"
	_, err := conn.Exec(context.Background(), str, userId, value)
	if err != nil {
		log.Fatal(err)
		return 0
	}
	return 1

}
func UpdateCarbonFootPrint(conn *pgx.Conn, userId int) int {
	exist := CheckCarbonFootprint(conn, userId)
	if !exist { // eger ilk defa fatura ekleniyorsa once tabloya o userı ekle.
		InsertCarbonFootprint(conn, userId)
	}

	var value string        // gecici degisken olarak kullanıldı
	var totalAmount float64 // gecici degisken olarak kullanıldı
	totalAmount = 0
	typeId := 1
	rows, err := conn.Query(context.Background(), "SELECT amount FROM bills WHERE user_id=$1 AND type_id=$2 ", userId, typeId) //butun elektrik faturalarının toplam miktarı
	if err != nil {
		log.Printf("Elektrik faturaları sorgusu başarısız: %v", err)
		return 0
	}
	defer rows.Close()
	for rows.Next() {
		if err := rows.Scan(&value); err != nil {
			log.Fatal(err)
		}
		floatValue, err := strconv.ParseFloat(value, 64) // 64, float64 için
		if err != nil {
			fmt.Println("Error while converting string to float:", err)
		}
		totalAmount += floatValue
	} // bu işlemlerin sonunda total electric faturası totalAmount değişkeninde float tipinde tutuluyor.
	totalWastedElectric := totalAmount / 4               //Elektrik: 1 kWh = 4 TL
	totalCarbonForElectric := 0.35 * totalWastedElectric // Elektrik: 0,35 kg CO₂/kWh --> emisyon faktörü

	totalAmount = 0
	typeId = 2
	rows1, err1 := conn.Query(context.Background(), "SELECT amount FROM bills WHERE user_id=$1 AND type_id=$2 ", userId, typeId) //butun su faturalarının toplam miktarı
	if err1 != nil {
		log.Printf("su faturaları sorgusu başarısız: %v", err)
		return 0
	}
	defer rows1.Close()
	for rows1.Next() {
		if err := rows1.Scan(&value); err != nil {
			log.Fatal(err)
		}
		floatValue, err := strconv.ParseFloat(value, 64) // 64, float64 için
		if err != nil {
			fmt.Println("Error while converting string to float:", err)
		}
		totalAmount += floatValue
	} // bu işlemlerin sonunda total su faturası totalAmount değişkeninde float tipinde tutuluyor.
	totalWastedWater := totalAmount / 25          //Su: 1 m³ = 25 TL
	totalCarbonForWater := 0.4 * totalWastedWater //Su: 0,4 kg CO₂/m³ --> emisyon faktörü

	totalAmount = 0
	typeId = 3
	rows2, err2 := conn.Query(context.Background(), "SELECT amount FROM bills WHERE user_id=$1 AND type_id=$2 ", userId, typeId) //butun dogalgaz faturalarının toplam miktarı
	if err2 != nil {
		log.Printf("dogalgaz faturaları sorgusu başarısız: %v", err)
		return 0
	}
	defer rows2.Close()
	for rows2.Next() {
		if err := rows2.Scan(&value); err != nil {
			log.Fatal(err)
		}
		floatValue, err := strconv.ParseFloat(value, 64) // 64, float64 için

		if err != nil {
			fmt.Println("Error while converting string to float:", err)
		}
		totalAmount += floatValue
	} // bu işlemlerin sonunda total dogalgaz faturası totalAmount değişkeninde float tipinde tutuluyor.
	totalWastedGas := totalAmount / 15         //Doğalgaz: 1 m³ = 15 TL
	totalCarbonForGas := 1.92 * totalWastedGas //Doğalgaz: 1,92 kg CO₂/m³ --> emisyon faktörü

	totalCarbon := totalCarbonForElectric + totalCarbonForGas + totalCarbonForWater
	//totalCarbon degiskeninde butun faturalardan elde edilen toplam karbon ayakizi değeri float olarak saklı bunu stringe cevirmelyiiz.
	carbon := strconv.FormatFloat(totalCarbon, 'f', 2, 64) // 2: ondalık basamak sayısı
	query := "UPDATE carbon_footprint SET value = $1 WHERE user_id = $2;"
	_, err3 := conn.Exec(context.Background(), query, carbon, userId)
	if err3 != nil {
		log.Fatal(err)
		return 0
	}
	return 1
}
func ShowCarbonFootPrint(conn *pgx.Conn, userId int) string {
	var value string
	exist := CheckCarbonFootprint(conn, userId)
	if exist {
		rows, _ := conn.Query(context.Background(), "SELECT value FROM carbon_footprint WHERE user_id=$1 ", userId)
		defer rows.Close()
		if rows.Next() {
			if err := rows.Scan(&value); err != nil {
				log.Fatal(err)
			}
		} else {
			return "-1"
		}

	} else {
		return "-1" // kullanıcının henüz bir faturası yok. karbon ayak izi hesaplanmamış.
	}
	return value
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
func ShowBills(conn *pgx.Conn, bill BillInfo) []Bill {
	var bills []Bill
	rows, err := conn.Query(context.Background(), "SELECT bill_name,year,month,amount FROM bills WHERE user_id=$1 AND type_id=$2", bill.UserId, bill.TypeId)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		var b Bill
		err := rows.Scan(&b.BillName, &b.Year, &b.Month, &b.Amount)
		b.UserId = bill.UserId
		if err != nil {
			log.Fatal(err)
		}
		b.TypeName = "random"
		bills = append(bills, b)
	}

	return bills
}
