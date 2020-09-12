package simple

import (
	"database/sql"
	"fmt"
	"testing"
)

type TestUser struct {
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

type TestArticle struct {
	GormModel
	Title      string `gorm:"not null"`
	Content    string `gorm:"not null"`
	CreateTime int64  `json:"createTime" form:"createTime"`
	UpdateTime int64  `json:"updateTime" form:"updateTime"`
}

func TestQueryParams(t *testing.T) {
	models := []interface{}{&TestUser{}, &TestArticle{}}

	if err := OpenDB("root:123456@tcp(localhost:3306)/bbsgo_db?charset=utf8mb4&parseTime=True&loc=Local",
		nil, 5, 20, models...); err != nil {
		panic(err)
	}

	var users []TestUser
	NewSqlCnd().Cols("id", "status", "nickname").In("id", []int64{1, 2, 3}).Find(db, &users)

	for _, user := range users {
		fmt.Println(user.Id, user.Nickname)
	}
}
