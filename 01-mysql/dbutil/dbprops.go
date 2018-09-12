package dbutil

import "fmt"

// Info ...
type Info struct {
	ID   int
	Name string
}

// DbDriver ...
const DbDriver = "mysql"
// User ...
const User = "root"
// Password ...
const Password = "123456"
// DbName ...
const DbName = "go_db1"
// TableName ...
const TableName = "person"

const MysqlServerIP = "192.168.41.166" //
// DataSourceName ...
// dataSourceName := "root:tyler@tcp(127.0.0.1:3306)/go_db1?charset=utf8"
var DataSourceName = fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8",
	User, Password,MysqlServerIP, DbName)
