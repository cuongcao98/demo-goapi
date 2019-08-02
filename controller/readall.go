package controller

import (
	"project/database"
	"github.com/gin-gonic/gin"
	"time"
)

func ReadAll(c * gin.Context)  {
	db := database.Conn()
	rows, err := db.Query("SELECT * FROM thongtins")
	if err != nil{
		c.JSON(500, gin.H{
			"messages" : "Story not found",
		});
	}

	var thongtin []Thongtin

	for rows.Next(){
		var id,sex,national_id int
		var username, password, email, phone string
		var birth time.Time
		var height float64

		err = rows.Scan(&id, &username, &password, &email, &birth, &sex, &phone, &national_id, &height)
		if err != nil {
			panic(err.Error())
		}
		
		// thongtin.Id = id
		// thongtin.Username = username
		// thongtin.Password = password
		// thongtin.Email = email
		// thongtin.Birth = birth
		// thongtin.Sex = sex
		// thongtin.Phone = phone
		// thongtin.National_id = national_id
		// thongtin.Height = height
		thongtin = append(thongtin, Thongtin{Id: id, Username: username, Password: password, Email: email, Birth: birth, Sex: sex, Phone: phone, National_id: national_id, Height: height})
	}

	c.JSON(200, thongtin)
	defer db.Close() // Hoãn lại việc close database connect cho đến khi hàm Read() thực hiệc xong
}
