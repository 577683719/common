package test_util

import (
	"fmt"
	"github.com/577683719/common/service_api/repository/dao"
	"strings"
)

var (
	host     = "124.71.69.65"
	port     = "3306"
	database = "ecs_schema"
	username = "root"
	password = "Bmx@123"
	charset  = "utf8mb4"
)

func InitDB() {
	dsn := strings.Join([]string{username, ":", password, "@tcp(", host, ":", port, ")/", database, "?charset=" + charset + "&parseTime=true"}, "")
	err := dao.Database(dsn)
	if err != nil {
		fmt.Println(err)
	}
}

var (
	host_prod     = "124.220.9.214"
	port_prod     = "30306"
	database_prod = "ecs_schema"
	username_prod = "root"
	password_prod = "Bmx@123"
	charset_prod  = "utf8mb4"
)

func InitDBProd() {
	dsn := strings.Join([]string{username_prod, ":", password_prod, "@tcp(", host_prod, ":", port_prod, ")/", database_prod, "?charset=" + charset_prod + "&parseTime=true"}, "")
	err := dao.Database(dsn)
	if err != nil {
		fmt.Println(err)
	}
}
