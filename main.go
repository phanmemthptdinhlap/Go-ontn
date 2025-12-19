package main

import (
	"fmt"
	gosql "go-ontn/go-sql"
)

func Khoitaodulieu(dbpath string) (*gosql.Data, error) {
	data, err := gosql.NewData(dbpath)
	if err != nil {
		return nil, err
	}
	return data, nil
}
func themcauhinh(data *gosql.Data, config gosql.Config) int {
	return data.AddConfig(config)
}

var db *gosql.Data

func main() {
	var err error
	db, err = Khoitaodulieu("ontn.db")
	if err != nil {
		fmt.Println("Lỗi khi khởi tạo cơ sở dữ liệu:", err)
		return
	}
	defer db.Close()

	config := gosql.Config{
		API:      "AIzaSyCsrZ7qUlb-taDEop8trr7tgPsWkCjzAi4",
		VaiTro:   "#VAI TRÒ: Bạn là một chuyên gia sư phạm giàu kinh nghiệm, bạn giảng dạy môn tin học, bạn có nhiều kinh nghiệm với chương trình giao dục phổ thông 2018.",
		HuongDan: "Hướng dẫn:",
		NhiemVu:  "Nhiệm vụ:",
		NoiDung:  "Nội dung:",
	}

	id := themcauhinh(db, config)
	fmt.Printf("Đã thêm cấu hình mới với ID: %d\n", id)
}
