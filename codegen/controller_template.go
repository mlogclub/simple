package codegen

import "html/template"

var controllerTmpl = template.Must(template.New("controller").Parse(controllerTmplCode))

const controllerTmplCode = `
package admin

import (
	"{{.PkgName}}/model"
	"{{.PkgName}}/services"
	"github.com/m-log/simple"
	"github.com/kataras/iris"
	"strconv"
)

type {{.Name}}Controller struct {
	Ctx             iris.Context
	{{.Name}}Service      *services.{{.Name}}Service
}

func (this *{{.Name}}Controller) GetBy(id int64) *simple.JsonResult {
	t := this.{{.Name}}Service.Get(id)
	if t == nil {
		return simple.ErrorMsg("Not found, id=" + strconv.FormatInt(id, 10))
	}
	return simple.JsonData(t)
}

func (this *{{.Name}}Controller) AnyList() *simple.JsonResult {
	list, paging := this.{{.Name}}Service.Query(simple.NewParamQueries(this.Ctx).PageAuto().Desc("id"))
	return simple.JsonData(&simple.PageResult{Results: list, Page: paging})
}

func (this *{{.Name}}Controller) PostCreate() *simple.JsonResult {
	t := &model.{{.Name}}{}
	this.Ctx.ReadForm(t)

	err := this.{{.Name}}Service.Create(t)
	if err != nil {
		return simple.ErrorMsg(err.Error())
	}
	return simple.JsonData(t)
}

func (this *{{.Name}}Controller) PostUpdate() *simple.JsonResult {
	id, err := simple.FormValueInt64(this.Ctx, "id")
	if err != nil {
		return simple.ErrorMsg(err.Error())
	}
	t := this.{{.Name}}Service.Get(id)
	if t == nil {
		return simple.ErrorMsg("entity not found")
	}

	this.Ctx.ReadForm(t)

	err = this.{{.Name}}Service.Update(t)
	if err != nil {
		return simple.ErrorMsg(err.Error())
	}
	return simple.JsonData(t)
}

`
