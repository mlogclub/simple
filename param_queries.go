package simple

import (
	"github.com/jinzhu/gorm"
	"github.com/kataras/iris"

	"github.com/m-log/simple/strcase"
)

type ParamQueries struct {
	Ctx         iris.Context
	Queries     []queryPair  // 条件
	OrderByCols []OrderByCol // 排序
	Paging      *Paging      // 分页
}

func NewParamQueries(ctx iris.Context) *ParamQueries {
	return &ParamQueries{
		Ctx: ctx,
	}
}

type queryPair struct {
	Query string        // 查询
	Args  []interface{} // 参数
}

func (o *ParamQueries) value(column string) string {
	name := strcase.ToLowerCamel(column)
	return o.Ctx.FormValue(name)
}

func (o *ParamQueries) Eq(column string, args ...interface{}) *ParamQueries {
	o.Where(column+" = ?", args)
	return o
}

func (o *ParamQueries) EqAuto(column string) *ParamQueries {
	value := o.value(column)
	if len(value) > 0 {
		return o.Eq(column, value)
	}
	return o
}

func (o *ParamQueries) NotEq(column string, args ...interface{}) *ParamQueries {
	o.Where(column+" <> ?", args)
	return o
}

func (o *ParamQueries) NotEqAuto(column string) *ParamQueries {
	value := o.value(column)
	if len(value) > 0 {
		return o.NotEq(column, value)
	}
	return o
}

func (o *ParamQueries) Gt(column string, args ...interface{}) *ParamQueries {
	o.Where(column+" > ?", args)
	return o
}

func (o *ParamQueries) GtAuto(column string) *ParamQueries {
	value := o.value(column)
	if len(value) > 0 {
		return o.Gt(column, value)
	}
	return o
}

func (o *ParamQueries) Gte(column string, args ...interface{}) *ParamQueries {
	o.Where(column+" >= ?", args)
	return o
}

func (o *ParamQueries) GteAuto(column string) *ParamQueries {
	value := o.value(column)
	if len(value) > 0 {
		return o.Gte(column, value)
	}
	return o
}

func (o *ParamQueries) Lt(column string, args ...interface{}) *ParamQueries {
	o.Where(column+" < ?", args)
	return o
}

func (o *ParamQueries) LtAuto(column string) *ParamQueries {
	value := o.value(column)
	if len(value) > 0 {
		return o.Lt(column, value)
	}
	return o
}

func (o *ParamQueries) Lte(column string, args ...interface{}) *ParamQueries {
	o.Where(column+" <= ?", args)
	return o
}

func (o *ParamQueries) LteAuto(column string) *ParamQueries {
	value := o.value(column)
	if len(value) > 0 {
		return o.Lte(column, value)
	}
	return o
}

func (o *ParamQueries) Like(column string, str string) *ParamQueries {
	o.Where(column+" like ?", "%"+str+"%")
	return o
}

func (o *ParamQueries) LikeAuto(column string) *ParamQueries {
	value := o.value(column)
	if len(value) > 0 {
		return o.Like(column, value)
	}
	return o
}

func (o *ParamQueries) Where(query string, args ...interface{}) *ParamQueries {
	o.Queries = append(o.Queries, queryPair{query, args})
	return o
}

func (o *ParamQueries) Asc(column string) *ParamQueries {
	return o.OrderBy(column, true)
}

func (o *ParamQueries) Desc(column string) *ParamQueries {
	return o.OrderBy(column, false)
}

func (o *ParamQueries) OrderBy(column string, asc bool) *ParamQueries {
	o.OrderByCols = append(o.OrderByCols, OrderByCol{Column: column, Asc: asc})
	return o
}

func (o *ParamQueries) Limit(limit int) *ParamQueries {
	o.Page(1, limit)
	return o
}

func (o *ParamQueries) Page(page, limit int) *ParamQueries {
	if o.Paging == nil {
		o.Paging = &Paging{Page: page, Limit: limit}
	} else {
		o.Paging.Page = page
		o.Paging.Limit = limit
	}
	return o
}

func (o *ParamQueries) PageAuto() *ParamQueries {
	paging := GetPaging(o.Ctx)
	return o.Page(paging.Page, paging.Limit)
}

func (o *ParamQueries) StartQuery(db *gorm.DB) *gorm.DB {
	retDb := db

	// where
	if len(o.Queries) > 0 {
		for _, query := range o.Queries {
			retDb = retDb.Where(query.Query, query.Args...)
		}
	}

	// order by
	if len(o.OrderByCols) > 0 {
		for _, sort := range o.OrderByCols {
			if sort.Asc {
				retDb = retDb.Order(sort.Column + " asc")
			} else {
				retDb = retDb.Order(sort.Column + " desc")
			}
		}
	}

	// limit
	if o.Paging != nil && o.Paging.Limit > 0 {
		retDb = retDb.Limit(o.Paging.Limit)
	}

	// offset
	if o.Paging != nil && o.Paging.Offset() > 0 {
		retDb = retDb.Offset(o.Paging.Offset())
	}

	return retDb
}

func (o *ParamQueries) StartCount(db *gorm.DB) *gorm.DB {
	retDb := db

	// where
	if len(o.Queries) > 0 {
		for _, query := range o.Queries {
			retDb = retDb.Where(query.Query, query.Args...)
		}
	}

	return retDb
}
