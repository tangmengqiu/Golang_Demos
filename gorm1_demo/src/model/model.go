package model

import (
	"fmt"

	"github.com/glog"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql" // mysql
	"github.com/spf13/viper"
)

//Product .
type Product struct {
	gorm.Model	//add three columes: createAt、updateAt、deleteAt
	Code  string
	Price uint
}

//InitDB init db
func InitDB() {

	dbName := viper.GetString("db.name")
	user := viper.GetString("db.username")
	password := viper.GetString("db.password")
	host := viper.GetString("db.addr")

	//connet to mysql ,you must ensure there is a mysql db named "db.name" in your local machine
	db, err := gorm.Open("mysql", user+":"+password+"@tcp("+host+")/"+dbName+"?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		glog.Errorf("open db filed:%v", err)
		panic(err)
	}
	fmt.Println("test create a table :Product and insert a record,and then check it:")
	// create table Product
	db.AutoMigrate(&Product{})

	//insert data
	db.Create(&Product{Code: "0001", Price: 10})

	//select * from product where id =1;
	var pd Product
	db.First(&pd, 1)

	fmt.Printf("product in db: %s whose id =1 \nit's Code is:%s\nit's Price is:%v",dbName, pd.Code,pd.Price)

}
