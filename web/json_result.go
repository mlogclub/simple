package web

import (
	"github.com/mlogclub/simple/common/structs"
	"github.com/mlogclub/simple/sqls"
)

type JsonResult struct {
	ErrorCode int    `json:"errorCode"`
	Message   string `json:"message"`
	Data      any    `json:"data"`
	Success   bool   `json:"success"`
}

func Json(code int, message string, data any, success bool) *JsonResult {
	return &JsonResult{
		ErrorCode: code,
		Message:   message,
		Data:      data,
		Success:   success,
	}
}

func JsonData(data any) *JsonResult {
	return &JsonResult{
		ErrorCode: 0,
		Data:      data,
		Success:   true,
	}
}

func JsonItemList(data []any) *JsonResult {
	if data == nil {
		data = []any{}
	}
	return &JsonResult{
		ErrorCode: 0,
		Data:      data,
		Success:   true,
	}
}

func JsonPageData(results any, page *sqls.Paging) *JsonResult {
	if results == nil {
		results = []any{}
	}
	return JsonData(&PageResult{
		Results: results,
		Page:    page,
	})
}

func JsonCursorData(results any, cursor string, hasMore bool) *JsonResult {
	if results == nil {
		results = []any{}
	}
	return JsonData(&CursorResult{
		Results: results,
		Cursor:  cursor,
		HasMore: hasMore,
	})
}

func JsonSuccess() *JsonResult {
	return &JsonResult{
		ErrorCode: 0,
		Data:      nil,
		Success:   true,
	}
}

func JsonError(err error) *JsonResult {
	if err == nil {
		return JsonSuccess()
	}
	if e, ok := err.(*CodeError); ok {
		return &JsonResult{
			ErrorCode: e.Code,
			Message:   e.Message,
			Data:      e.Data,
			Success:   false,
		}
	}
	return &JsonResult{
		ErrorCode: 0,
		Message:   err.Error(),
		Data:      nil,
		Success:   false,
	}
}

func JsonErrorMsg(message string) *JsonResult {
	return &JsonResult{
		ErrorCode: 0,
		Message:   message,
		Data:      nil,
		Success:   false,
	}
}
func JsonErrorCode(code int, message string) *JsonResult {
	return &JsonResult{
		ErrorCode: code,
		Message:   message,
		Data:      nil,
		Success:   false,
	}
}

func JsonErrorData(code int, message string, data any) *JsonResult {
	return &JsonResult{
		ErrorCode: code,
		Message:   message,
		Data:      data,
		Success:   false,
	}
}

type RspBuilder struct {
	Data map[string]any
}

func NewEmptyRspBuilder() *RspBuilder {
	return &RspBuilder{Data: make(map[string]any)}
}

func NewRspBuilder(obj any) *RspBuilder {
	return NewRspBuilderExcludes(obj)
}

func NewRspBuilderExcludes(obj any, excludes ...string) *RspBuilder {
	return &RspBuilder{Data: structs.StructToMap(obj, excludes...)}
}

func (builder *RspBuilder) Put(key string, value any) *RspBuilder {
	builder.Data[key] = value
	return builder
}

func (builder *RspBuilder) Build() map[string]any {
	return builder.Data
}

func (builder *RspBuilder) JsonResult() *JsonResult {
	return JsonData(builder.Data)
}

func ConvertList[T any](results []T, conv func(item T) map[string]any) []map[string]any {
	list := make([]map[string]any, 0)
	for _, item := range results {
		if ret := conv(item); ret != nil {
			list = append(list, ret)
		}
	}
	return list
}
