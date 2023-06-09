package test

import (
	"gin_gorm_oj/define"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"testing"
)

func TestGormTest(t *testing.T) {
	dsn := "root:root@tcp(127.0.0.1:3306)/gin_gorm_oj?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		t.Fatal(err)
	}

	db.AutoMigrate(define.ProblemBasic{})

	//problem := models.ProblemBasic{
	//	Identity:   "111",
	//	CategoryId: "1",
	//	Title:      "www",
	//	Content:    "asfjsdnjf",
	//	MaxMem:     23,
	//	MaxRunTime: 234,
	//}
	//db.Save(&problem)
	//db.AutoMigrate(models.Problem{})
	//data := make([]*models.ProblemBasic, 0)
	//err = db.Find(&data).Error
	//if err != nil {
	//	t.Fatal(err)
	//}
	//for _, v := range data {
	//	fmt.Printf("problem ===> %v\n", v)
	//}
}
