package controller

import (
	"time"
)

type Thongtin struct{
	Id int 
	Username string
	Password string	
	Email string 
	Birth time.Time 
	Sex int 
	Phone string 
	National_id int 
	Height float64 
}