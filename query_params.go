package simple

import (
	"github.com/jinzhu/gorm"
	"github.com/kataras/iris"

	"github.com/mlogclub/simple/strcase"
)

type QueryParams struct {
	Ctx         iris.Context
	Queries     []queryPair  // 条件
	OrderByCols []OrderByCol // 排序
	Paging      *Paging      // 分页
}

func NewQueryParams(ctx iris.Context) *QueryParams {
	return &QueryParams{
		Ctx: ctx,
	}
}

type queryPair struct {
	Query string        // 查询
	Args  []interface{} // 参数
}

func (o *QueryParams) value(column string) string {
	name := strcase.ToLowerCamel(column)
	return o.Ctx.FormValue(name)
}

func (o *QueryParams) Eq(column string, args ...interface{}) *QueryParams {
	o.Where(column+" = ?", args)
	return o
}

func (o *QueryParams) EqAuto(column string) *QueryParams {
	value := o.value(column)
	if len(value) > 0 {
		return o.Eq(column, value)
	}
	return o
}

func (o *QueryParams) NotEq(column string, args ...interface{}) *QueryParams {
	o.Where(column+" <> ?", args)
	return o
}

func (o *QueryParams) NotEqAuto(column string) *QueryParams {
	value := o.value(column)
	if len(value) > 0 {
		return o.NotEq(column, value)
	}
	return o
}

func (o *QueryParams) Gt(column string, args ...interface{}) *QueryParams {
	o.Where(column+" > ?", args)
	return o
}

func (o *QueryParams) GtAuto(column string) *QueryParams {
	value := o.value(column)
	if len(value) > 0 {
		return o.Gt(column, value)
	}
	return o
}

func (o *QueryParams) Gte(column string, args ...interface{}) *QueryParams {
	o.Where(column+" >= ?", args)
	return o
}

func (o *QueryParams) GteAuto(column string) *QueryParams {
	value := o.value(column)
	if len(value) > 0 {
		return o.Gte(column, value)
	}
	return o
}

func (o *QueryParams) Lt(column string, args ...interface{}) *QueryParams {
	o.Where(column+" < ?", args)
	return o
}

func (o *QueryParams) LtAuto(column string) *QueryParams {
	value := o.value(column)
	if len(value) > 0 {
		return o.Lt(column, value)
	}
	return o
}

func (o *QueryParams) Lte(column string, args ...interface{}) *QueryParams {
	o.Where(column+" <= ?", args)
	return o
}

func (o *QueryParams) LteAuto(column string) *QueryParams {
	value := o.value(column)
	if len(value) > 0 {
		return o.Lte(column, value)
	}
	return o
}

func (o *QueryParams) Like(column string, str string) *QueryParams {
	o.Where(column+" like ?", "%"+str+"%")
	return o
}

func (o *QueryParams) LikeAuto(column string) *QueryParams {
	value := o.value(column)
	if len(value) > 0 {
		return o.Like(column, value)
	}
	return o
}

func (o *QueryParams) Where(query string, args ...interface{}) *QueryParams {
	o.Queries = append(o.Queries, queryPair{query, args})
	return o
}

func (o *QueryParams) Asc(column string) *QueryParams {
	return o.OrderBy(column, true)
}

func (o *QueryParams) Desc(column string) *QueryParams {
	return o.OrderBy(column, false)
}

func (o *QueryParams) OrderBy(column string, asc bool) *QueryParams {
	o.OrderByCols = append(o.OrderByCols, OrderByCol{Column: column, Asc: asc})
	return o
}

func (o *QueryParams) Limit(limit int) *QueryParams {
	o.Page(1, limit)
	return o
}

func (o *QueryParams) Page(page, limit int) *QueryParams {
	if o.Paging == nil {
		o.Paging = &Paging{Page: page, Limit: limit}
	} else {
		o.Paging.Page = page
		o.Paging.Limit = limit
	}
	return o
}

func (o *QueryParams) PageAuto() *QueryParams {
	paging := GetPaging(o.Ctx)
	return o.Page(paging.Page, paging.Limit)
}

func (o *QueryParams) StartQuery(db *gorm.DB) *gorm.DB {
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

func (o *QueryParams) StartCount(db *gorm.DB) *gorm.DB {
	retDb := db

	// where
	if len(o.Queries) > 0 {
		for _, query := range o.Queries {
			retDb = retDb.Where(query.Query, query.Args...)
		}
	}

	return retDb
}
