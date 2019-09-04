package simple

import (
	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
)

type GormModel struct {
	Id int64 `gorm:"PRIMARY_KEY;AUTO_INCREMENT" json:"id"`
}

type DBConfiguration struct {
	Dialect        string
	Url            string
	MaxIdle        int
	MaxActive      int
	EnableLogModel bool
	Models         []interface{}
}

var db *gorm.DB

// 打开数据库
func OpenDB(conf *DBConfiguration) (*gorm.DB, error) {
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return "t_" + defaultTableName
	}

	var err error
	db, err := gorm.Open(conf.Dialect, conf.Url)
	if err != nil {
		return nil, err
	}

	db.SingularTable(true)

	maxIdle := 10
	if conf.MaxIdle > 0 {
		maxIdle = conf.MaxIdle
	}
	db.DB().SetMaxIdleConns(maxIdle)

	maxActive := 50
	if conf.MaxActive > 0 {
		maxActive = conf.MaxActive
	}
	db.DB().SetMaxOpenConns(maxActive)

	db.LogMode(conf.EnableLogModel)

	if err != nil {
		log.Errorf("opens database failed: %s", err.Error())
	}
	if len(conf.Models) > 0 {
		if err = db.AutoMigrate(conf.Models...).Error; nil != err {
			log.Errorf("auto migrate tables failed: %s", err.Error())
		}
	}
	return db
}

// 关闭连接
func CloseDB() {
	if db == nil {
		return
	}
	if err := db.Close(); nil != err {
		log.Errorf("Disconnect from database failed: %s", err.Error())
	}
}

// 获取数据库链接
func GetDB() *gorm.DB {
	return db
}

// 事务环绕
func Tx(db *gorm.DB, txFunc func(tx *gorm.DB) error) (err error) {
	tx := db.Begin()
	if tx.Error != nil {
		return
	}

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			panic(r) // re-throw panic after Rollback
		} else if err != nil {
			tx.Rollback()
		} else {
			err = tx.Commit().Error
		}
	}()

	err = txFunc(tx)
	return err
}

// 排序信息
type OrderByCol struct {
	Column string // 排序字段
	Asc    bool   // 是否正序
}

// 分页返回数据
type PageResult struct {
	Page    *Paging     `json:"page"`    // 分页信息
	Results interface{} `json:"results"` // 数据
}

// Cursor分页返回数据
type CursorResult struct {
	Results interface{} `json:"results"` // 数据
	Cursor  string      `json:"cursor"`  // 下一页
}

// 分页请求数据
type Paging struct {
	Page  int `json:"page"`  // 页码
	Limit int `json:"limit"` // 每页条数
	Total int `json:"total"` // 总数据条数
}

func (p *Paging) Offset() int {
	offset := 0
	if p.Page > 0 {
		offset = (p.Page - 1) * p.Limit
	}
	return offset
}

func (p *Paging) TotalPage() int {
	if p.Total == 0 || p.Limit == 0 {
		return 0
	}
	totalPage := p.Total / p.Limit
	if p.Total%p.Limit > 0 {
		totalPage = totalPage + 1
	}
	return totalPage
}
