package models

import (
	"gorm.io/gorm"
)

type ProblemBasic struct {
	gorm.Model
	Identity          string             `gorm:"column:identity;type:varchar(36);" json:"identity"` //问题表的唯一标识
	ProblemCategories []*ProblemCategory `gorm:"foreignKey:problem_id;references:id"`               //关联问题分类表
	Title             string             `gorm:"column:title;type:varchar(255);" json:"title"`      //文章标题
	Content           string             `gorm:"column:content;type:text;" json:"Content"`          //文章正文
	MaxRuntime        int                `gorm:"column:max_runtime;type:int"json:"max_runtime"`     //最大运行时间
	MaxMem            int                `gorm:"column:max_mem;type:int"json:"max_mem"`             //最大运行内存

}

func (table *ProblemBasic) Tablename() string {
	return "problem_basic"
}

func GetProblemList(keyword, categoryIdentity string) *gorm.DB {
	tx := Db.Model(new(ProblemBasic)).Preload("ProblemCategories").Preload("ProblemCategories.CategoryBasic").
		Where("title like ? OR content like ?", "%"+keyword+"%", "%"+keyword+"%")
	if categoryIdentity != "" {
		tx.Joins("RIGHT JOIN problem_category pc on pc.problem_id = problem_basic.id").
			Where("pc.category_id = (SELECT cb.id FROM category_basic cb.identity = ?)", categoryIdentity)
	}
	return tx
}
