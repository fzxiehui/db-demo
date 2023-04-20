package model

import (
	"database/sql/driver"
	"fmt"
	"strings"
	"time"

	"gorm.io/gorm"
)

/*
 * 定义了两个业务模型
 * 1. Passport 用户账号
 * 2. User 用户信息
 */

type Username string

type Password string

func (p *Password) Scan(src interface{}) error {
	*p = Password(fmt.Sprintf("@%v@", src))
	return nil
}

func (p *Password) Value() (driver.Value, error) {
	*p = Password(strings.Trim(string(*p), "@"))
	return p, nil
}

type User struct {
	gorm.Model        // ID uint CreatAt time.Time UpdateAt time.Time DeleteAt gorm.DeleteAt If it is repeated with the definition will be ignored
	ID         uint   `gorm:"primary_key"`
	Name       string `gorm:"column:name"`
	Age        int    `gorm:"column:age;type:varchar(64)"`
	Role       string `gorm:"column:role;type:varchar(64)"`
	Friends    []User `gorm:"-"` // only local used gorm ignore
}

type Passport struct {
	ID        int
	Username  Username // will be field.String
	Password  Password // will be field.Field because type Password set Scan and Value
	LoginTime time.Time
}
