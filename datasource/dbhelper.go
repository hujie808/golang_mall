package datasource

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql" //初始化mysql
	"github.com/go-xorm/xorm"
	"log"
	"sync"
	"web_iris/golang_mall/conf"
)

var dbLock sync.Mutex

var masterInstance *xorm.Engine

//调用数据库是的单例
func InstanceDbMaster() *xorm.Engine {

	if masterInstance != nil {
		return masterInstance //如果存在就直接返回
	}
	dbLock.Lock() //创建时锁定
	defer dbLock.Unlock()
	if masterInstance != nil {
		return masterInstance //并发进入时,再次判断
	}
	return NewDbMaster() //如果不存在就创建
}

func NewDbMaster() *xorm.Engine {
	sourcename := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8",
		conf.DbMaster.User,
		conf.DbMaster.Psw,
		conf.DbMaster.Host,
		conf.DbMaster.Port,
		conf.DbMaster.Database)

	instance, err := xorm.NewEngine(conf.DriverName, sourcename)
	if err != nil {
		log.Fatal("dbhelper.InstanceDbMaster NewEngine error ", err)
		return nil
	}
	instance.ShowSQL(true)
	//instance.ShowSQL(false)
	masterInstance = instance
	return masterInstance
}
