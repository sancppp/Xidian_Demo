package mysql

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"sync"
	"tianzhenxiongProject/model"
)

var MysqlDB *Mysql
var once sync.Once

type Mysql struct {
	conn *gorm.DB
}

//封装mysql，对外提供 MysqlDB.GetConn()接口
func Default() {
	temp := &Mysql{}
	dsn := "root:root123456@tcp(127.0.0.1:3306)/TzxDemo?charset=utf8mb4&parseTime=True&loc=Local"
	conn, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database,err:" + err.Error())
	}
	sqlDB, err := conn.DB()
	if err != nil {
		panic(err)
	}
	_ = conn.AutoMigrate(&model.Student{})
	sqlDB.SetMaxOpenConns(100) //设置数据库连接池最大连接数
	sqlDB.SetMaxIdleConns(20)  //连接池最大允许的空闲连接数，如果没有sql任务需要执行的连接数大于20，超过的连接会被连接池关闭。
	temp.conn = conn
	MysqlDB = temp
}

//封装mysql，对外提供 MysqlDB.GetConn()接口
func (mysql *Mysql) GetConn() *gorm.DB {
	return mysql.conn
}
