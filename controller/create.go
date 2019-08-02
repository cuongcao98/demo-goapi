package controller

import (
	"project/database"
	"github.com/gin-gonic/gin"
	"time"
	"encoding/hex"
	"crypto/md5"
)
func Create(c * gin.Context){
	db := database.Conn()

	type Adduser struct {
		Username string `form:"username" json:"username" binding:"required"`
		Password string	`form:"password" json:"password" binding:"required"`
		Email string `form:"email" json:"email" binding:"required"`
		Birth time.Time `form:"birth" json:"birth" binding:"required"`
		Sex int `form:"sex" json:"sex" binding:"required"`
		Phone string `form:"phone" json:"phone" binding:"required"`
		National_id int `form:"national_id" json:"national_id" binding:"required"`
		Height float64 `form:"height" json:"height" binding:"required"`
	}

	var json Adduser

	if err := c.ShouldBindJSON(&json); err == nil {
		insPost, err := db.Prepare("INSERT INTO thongtins VALUES(null,?,?,?,?,?,?,?,?)",)
		if err != nil {
			c.JSON(500, gin.H{
				"messages" : err,
			})
		}
	

		insPost.Exec(json.Username, GetMD5Hash(json.Password), json.Email, json.Birth, json.Sex, json.Phone, json.National_id, json.Height) 
		c.JSON(200, gin.H{
			"messages": "inserted",
		})

	} else {
		c.JSON(500, gin.H{"error": err.Error()})
	}

	defer db.Close()
}

func GetMD5Hash(text string) string {
    hasher := md5.New()
    hasher.Write([]byte(text))
    return hex.EncodeToString(hasher.Sum(nil))
}