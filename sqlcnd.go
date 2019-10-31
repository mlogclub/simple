package simple

import "github.com/jinzhu/gorm"

type SqlCnd struct {
	Params []ParamPair
	Orders []OrderByCol
	Limit  int
	Offset int
}

func NewSqlCnd(query string, args ...interface{}) *SqlCnd {
	q := &SqlCnd{}
	q.Params = append(q.Params, ParamPair{query, args})
	return q
}

func (s *SqlCnd) Asc(column string) *SqlCnd {
	s.Orders = append(s.Orders, OrderByCol{Column: column, Asc: true})
	return s
}

func (s *SqlCnd) Desc(column string) *SqlCnd {
	s.Orders = append(s.Orders, OrderByCol{Column: column, Asc: false})
	return s
}

func (s *SqlCnd) Size(size int) *SqlCnd {
	s.Limit = size
	return s
}

func (s *SqlCnd) Page(page, size int) *SqlCnd {
	p := Paging{Page: page, Limit: size}
	s.Limit = p.Limit
	s.Offset = p.Offset()
	return s
}

func (s *SqlCnd) Query(db *gorm.DB) *gorm.DB {
	ret := db
	if len(s.Params) > 0 {
		for _, param := range s.Params {
			ret = ret.Where(param.Query, param.Args...)
		}
	}
	if len(s.Orders) > 0 {
		for _, order := range s.Orders {
			if order.Asc {
				ret = ret.Order(order.Column + " ASC")
			} else {
				ret = ret.Order(order.Column + " DESC")
			}
		}
	}
	if s.Limit > 0 {
		ret = ret.Limit(s.Limit)
	}
	if s.Offset > 0 {
		ret = ret.Limit(s.Offset)
	}
	return ret
}
