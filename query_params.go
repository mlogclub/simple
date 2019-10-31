package simple

import (
	"github.com/jinzhu/gorm"
	"github.com/kataras/iris"

	"github.com/mlogclub/simple/strcase"
)

type QueryParams struct {
	Ctx iris.Context
	Cnd SqlCnd
}

func NewQueryParams(ctx iris.Context) *QueryParams {
	return &QueryParams{
		Ctx: ctx,
		Cnd: SqlCnd{},
	}
}

func (q *QueryParams) getValueByColumn(column string) string {
	if q.Ctx == nil {
		return "xxx"
	}
	fieldName := strcase.ToLowerCamel(column)
	return q.Ctx.FormValue(fieldName)
}

func (q *QueryParams) Eq(column string) *QueryParams {
	value := q.getValueByColumn(column)
	if len(value) > 0 {
		q.Cnd.Eq(column, value)
	}
	return q
}

func (q *QueryParams) NotEq(column string) *QueryParams {
	value := q.getValueByColumn(column)
	if len(value) > 0 {
		q.Cnd.NotEq(column, value)
	}
	return q
}

func (q *QueryParams) Gt(column string) *QueryParams {
	value := q.getValueByColumn(column)
	if len(value) > 0 {
		q.Cnd.Gt(column, value)
	}
	return q
}

func (q *QueryParams) Gte(column string) *QueryParams {
	value := q.getValueByColumn(column)
	if len(value) > 0 {
		q.Cnd.Gte(column, value)
	}
	return q
}

func (q *QueryParams) Lt(column string) *QueryParams {
	value := q.getValueByColumn(column)
	if len(value) > 0 {
		q.Cnd.Lt(column, value)
	}
	return q
}

func (q *QueryParams) Lte(column string) *QueryParams {
	value := q.getValueByColumn(column)
	if len(value) > 0 {
		q.Cnd.Lte(column, value)
	}
	return q
}

func (q *QueryParams) Like(column string) *QueryParams {
	value := q.getValueByColumn(column)
	if len(value) > 0 {
		q.Cnd.Like(column, value)
	}
	return q
}

func (q *QueryParams) Asc(column string) *QueryParams {
	q.Cnd.Asc(column)
	return q
}

func (q *QueryParams) Desc(column string) *QueryParams {
	q.Cnd.Desc(column)
	return q
}

func (q *QueryParams) Page() *QueryParams {
	if q.Ctx == nil {
		return q
	}
	paging := GetPaging(q.Ctx)
	q.Cnd.Page(paging.Page, paging.Limit)
	return q
}

func (q *QueryParams) Find(db *gorm.DB, out interface{}) error {
	return q.Cnd.Find(db, out)
}

func (q *QueryParams) FindOne(db *gorm.DB, out interface{}) error {
	return q.Cnd.FindOne(db, out)
}

func (q *QueryParams) Count(db *gorm.DB, model interface{}) (int64, error) {
	return q.Cnd.Count(db, model)
}
