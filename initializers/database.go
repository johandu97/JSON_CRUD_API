package initializers

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDB() {

	// Thông tin kết nối MySQL
	// dsn := "root:my-secret-pw@tcp(127.0.0.1:3306)/crud?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := "root:my-secret-pw@tcp(127.0.0.1:3306)/crud?charset=utf8mb4&parseTime=True&loc=Local"

	// Mở kết nối tới DB
	var err error

	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		// Nếu có lỗi
		fmt.Println("Kết nối DB thất bại:", err)
		return
	}

	// Nếu không có lỗi
	fmt.Println("Kết nối DB thành công:", DB)
}
