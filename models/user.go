package models

import (
	"BCP/db_mysql"
	"BCP/utils"
	"fmt"
)

type User struct {
	Id            int      `form:"id"`
	Phone         string  `form:"phone"`
	Password      string   `form:"password"`
}



func (u User) AddUser() (int64,error){
	fmt.Println(u.Password,u.Phone)
	//脱敏
	u.Password = utils.MD5HashString(u.Password)

	rs, err := db_mysql.Db.Exec("insert into user(phone,password)"+
		"value(?,?)", u.Phone, u.Password)

	if err != nil {
		return 0, err
	}
	id, err := rs.RowsAffected()
	if err != nil {
		return 0, err
	}
	return  id, nil
}


func (u User) QueryUser()  (*User,error){
	//把脱敏的密码的md5值重新赋值为密码进行存储
	u.Password = utils.MD5HashString(u.Password)

	row := db_mysql.Db.QueryRow("select phone from user where phone = ? and password = ?",
		u.Phone, u.Password)

	err := row.Scan(&u.Phone)
	if err != nil {
		return nil,err
	}
	return &u,nil
}

func (u User) QueryUserByPhone() (*User, error){
	fmt.Println( u.Phone)
	row := db_mysql.Db.QueryRow("select id from user where phone =?",u.Phone)

	err := row.Scan(&u.Id)

	if err != nil {
			return nil, err
	}
	return &u, nil
}