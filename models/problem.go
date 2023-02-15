package models

import "gorm.io/gorm"

type Problem struct {
	gorm.Model
	Identity   string `gorm:"column:identity;type:varchar(36);" json:"identity"`        //问题表的唯一标识
	CategoryId string `gorm:"column:category_id;type:varchar(255);" json:"category_id"` //分类Id，以逗号分类
	Title      string `gorm:"column:title;type:varchar(255);" json:"title"`             //文章标题
	Content    string `gorm:"column:content;type:text;" json:"Content"`                 //文章正文
}

func (table Problem) Tablename() string {

}