package test

import (
	"fmt"
	"getcahrzp.cn/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"testing"
)

func TestGormTest(t *testing.T) {
	dsn := "root:root@tcp(127.0.0.1:3306)/gin_gorm_oj?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // 使用单数表名
		},
	})
	if err != nil {
		t.Fatal(err)
	}
	data := make([]*models.Problem, 0)
	err = db.Find(&data).Error
	db.Find(&data)
	if err != nil {
		t.Fatal(err)
	}
	for _, v := range data {
		fmt.Printf("Problem ==> %v \n", v)
	}
}
