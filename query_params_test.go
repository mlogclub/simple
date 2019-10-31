package simple

import (
	"database/sql"
	"testing"
)

type User struct {
	GormModel
	Username    sql.NullString `gorm:"size:32;unique;" json:"username" form:"username"`
	Email       sql.NullString `gorm:"size:128;unique;" json:"email" form:"email"`
	Nickname    string         `gorm:"size:16;" json:"nickname" form:"nickname"`
	Avatar      string         `gorm:"type:text" json:"avatar" form:"avatar"`
	Password    string         `gorm:"size:512" json:"password" form:"password"`
	Status      int            `gorm:"index:idx_status;not null" json:"status" form:"status"`
	Roles       string         `gorm:"type:text" json:"roles" form:"roles"`
	Type        int            `gorm:"not null" json:"type" form:"type"`
	Description string         `gorm:"type:text" json:"description" form:"description"`
	CreateTime  int64          `json:"createTime" form:"createTime"`
	UpdateTime  int64          `json:"updateTime" form:"updateTime"`
}

func TestQueryParams(t *testing.T) {
	if err := OpenDB("root:123456@tcp(localhost:3306)/mlog_db2?charset=utf8mb4&parseTime=True&loc=Local", 5, 20, true); err != nil {
		panic(err)
	}
	var users []User
	NewQueryParams(nil).Eq("username").Page().Desc("id").Query(DB()).Find(&users)
	NewSqlCnd().Where("username = ? or email = ?", "username", "email").Where("password = ?", 123).Query(DB()).Find(&users)
}
