package models

import (
	"database/sql/driver"
	"fmt"
	"gin_gorm_oj/define"
	"time"
)

type MyTime time.Time

func (t MyTime) MarshalJSON() ([]byte, error) {
	formatted := fmt.Sprintf("\"%v\"", time.Time(t).Format(define.DateLayout))
	return []byte(formatted), nil
}

func (t MyTime) Value() (driver.Value, error) {
	tTime := time.Time(t)
	return tTime.Format(define.DateLayout), nil
}
