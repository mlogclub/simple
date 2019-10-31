package simple

import (
	"github.com/kataras/iris"

	"github.com/mlogclub/simple/strcase"
)

type QueryParams struct {
	Ctx iris.Context
	SqlCnd
}

func NewQueryParams(ctx iris.Context) *QueryParams {
	return &QueryParams{
		Ctx: ctx,
	}
}

func (q *QueryParams) getValueByColumn(column string) string {
	if q.Ctx == nil {
		return "xxx"
	}
	fieldName := strcase.ToLowerCamel(column)
	return q.Ctx.FormValue(fieldName)
}

func (q *QueryParams) EqByReq(column string) *QueryParams {
	value := q.getValueByColumn(column)
	if len(value) > 0 {
		q.Eq(column, value)
	}
	return q
}

func (q *QueryParams) NotEqByReq(column string) *QueryParams {
	value := q.getValueByColumn(column)
	if len(value) > 0 {
		q.NotEq(column, value)
	}
	return q
}

func (q *QueryParams) GtByReq(column string) *QueryParams {
	value := q.getValueByColumn(column)
	if len(value) > 0 {
		q.Gt(column, value)
	}
	return q
}

func (q *QueryParams) GteByReq(column string) *QueryParams {
	value := q.getValueByColumn(column)
	if len(value) > 0 {
		q.Gte(column, value)
	}
	return q
}

func (q *QueryParams) LtByReq(column string) *QueryParams {
	value := q.getValueByColumn(column)
	if len(value) > 0 {
		q.Lt(column, value)
	}
	return q
}

func (q *QueryParams) LteByReq(column string) *QueryParams {
	value := q.getValueByColumn(column)
	if len(value) > 0 {
		q.Lte(column, value)
	}
	return q
}

func (q *QueryParams) LikeByReq(column string) *QueryParams {
	value := q.getValueByColumn(column)
	if len(value) > 0 {
		q.Like(column, value)
	}
	return q
}

func (q *QueryParams) PageByReq() *QueryParams {
	if q.Ctx == nil {
		return q
	}
	paging := GetPaging(q.Ctx)
	q.Page(paging.Page, paging.Limit)
	return q
}
