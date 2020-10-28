package db_mysql

import (
	"database/sql"
	"fmt"
	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
)

var Db *sql.DB

func Connect() {
	config := beego.AppConfig
	dbDriver := config.String("db_driver")
	dbUser := config.String("db_user")
	db_Ip := config.String("db_ip")
	db_Name := config.String("db_name")
	db_Password := config.String("db_password")

	connUrl := dbUser + ":" + db_Password +  "@tcp(" + db_Ip + ")/" + db_Name + "?charset=utf8"
	fmt.Println(connUrl)
	db, err := sql.Open(dbDriver, connUrl)


	if err != nil {
		panic("数据库连接错误，请检查配置")
	}
	Db = db
	fmt.Println(Db)
}
