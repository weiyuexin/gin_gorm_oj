package models

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	Name     string `gorm:"column:name;type:varchar(100)" json:"name"`        //分类名称
	Identity string `gorm:"column:identity;type:varchar(36)" json:"identity"` //分类的唯一标识
	ParentId uint   `gorm:"column:parent_id;type:int(11);" json:"parent_id"`  //父级ID
	//ProblemId  uint   `gorm:"column:problem_id;type:int(11);" json:"problem_id"`   // 问题的ID
	//CategoryId uint   `gorm:"column:category_id;type:int(11);" json:"category_id"` // 分类的ID
}

func (table *Category) TableName() string {
	return "category"
}
