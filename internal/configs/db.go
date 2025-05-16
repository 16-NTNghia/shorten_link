package configs

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func ConnectDB() *sql.DB {
	LoadEnv()

	db_port := GetEnv("DB_PORT")
	db_username := GetEnv("DB_USERNAME")
	db_password := GetEnv("DB_PASSWORD")
	db_name := GetEnv("DB_NAME")
	service := GetEnv("SERVICE")

	dataSourceName := fmt.Sprintf("%s:%s@tcp(localhost:%s)/%s?parseTime=true", db_username, db_password, db_port, db_name)

	db, err := sql.Open(service, dataSourceName)
	if err != nil {
		log.Panic("Không thể khởi tạo kết nối:", err)
	}

	// Kiểm tra kết nối thật sự
	if err := db.Ping(); err != nil {
		log.Panic("Không thể kết nối database:", err)
	}

	log.Println("Kết nối database thành công!")

	return db
}
