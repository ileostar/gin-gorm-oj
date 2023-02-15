package models

import "gorm.io/gorm"

type Problem struct {
	gorm.Model
	Identity   string `gorm:"column:identity;type:varchar(36);" json:"identity"`        //问题表的唯一标识
	CategoryId string `gorm:"column:category_id;type:varchar(255);" json:"category_id"` //分类Id，以逗号分类
	Title      string `gorm:"column:title;type:varchar(255);" json:"title"`             //文章标题
	Content    string `gorm:"column:content;type:text;" json:"Content"`                 //文章正文
	MaxRuntime int    `gorm:"column:max_runtime;type:int"json:"max_runtime"`            //最大运行时间
	MaxMem     int    `gorm:"column:max_mem;type:int"json:"max_mem"`                    //最大运行内存

}

func (table *Problem) Tablename() string {
	return "problem"
}
