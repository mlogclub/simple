package simple

import "github.com/jinzhu/gorm"

type SqlCnd struct {
	Query  string
	Args   []interface{}
	Orders []string
	Limit  int
	Offset int
}

func NewSqlCnd(query string, args ...interface{}) *SqlCnd {
	q := &SqlCnd{}
	q.Query = query
	q.Args = args
	return q
}

func (q *SqlCnd) Order(order string) *SqlCnd {
	q.Orders = append(q.Orders, order)
	return q
}

func (q *SqlCnd) Size(size int) *SqlCnd {
	q.Limit = size
	return q
}

func (q *SqlCnd) Page(page, size int) *SqlCnd {
	p := Paging{Page: page, Limit: size}
	q.Limit = p.Limit
	q.Offset = p.Offset()
	return q
}

func (q *SqlCnd) Exec(db *gorm.DB) *gorm.DB {
	ret := db.Where(q.Query, q.Args...)
	if q.Limit > 0 {
		ret = ret.Limit(q.Limit)
	}
	if q.Offset > 0 {
		ret = ret.Limit(q.Offset)
	}
	if len(q.Orders) > 0 {
		for _, order := range q.Orders {
			ret = ret.Order(order)
		}
	}
	return ret
}
