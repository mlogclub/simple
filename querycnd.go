package simple

import "github.com/jinzhu/gorm"

type QueryCnd struct {
	Query  string
	Args   []interface{}
	Orders []string
	Limit  int
	Offset int
}

func NewQueryCnd(query string, args ...interface{}) *QueryCnd {
	q := &QueryCnd{}
	q.Query = query
	q.Args = args
	return q
}

func (q *QueryCnd) Order(order string) *QueryCnd {
	q.Orders = append(q.Orders, order)
	return q
}

func (q *QueryCnd) Size(size int) *QueryCnd {
	q.Limit = size
	return q
}

func (q *QueryCnd) Page(page, size int) *QueryCnd {
	p := Paging{Page: page, Limit: size}
	q.Limit = p.Limit
	q.Offset = p.Offset()
	return q
}

func (q *QueryCnd) DoQuery(db *gorm.DB) *gorm.DB {
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
