package codegen

import "html/template"

var serviceTmpl = template.Must(template.New("service").Parse(serviceTmplCode))

const serviceTmplCode = `
package services

import (
	"{{.PkgName}}/model"
	"{{.PkgName}}/repositories"
	"github.com/mlogclub/simple"
)

type {{.Name}}Service struct {
	{{.Name}}Repository *repositories.{{.Name}}Repository
}

func New{{.Name}}Service() *{{.Name}}Service {
	return &{{.Name}}Service {
        {{.Name}}Repository: repositories.New{{.Name}}Repository(),
    }
}

func (this *{{.Name}}Service) Get(id int64) *model.{{.Name}} {
	return this.{{.Name}}Repository.Get(simple.GetDB(), id)
}

func (this *{{.Name}}Service) Take(where ...interface{}) *model.{{.Name}} {
	return this.{{.Name}}Repository.Take(simple.GetDB(), where...)
}

func (this *{{.Name}}Service) QueryCnd(cnd *simple.QueryCnd) (list []model.{{.Name}}, err error) {
	return this.{{.Name}}Repository.QueryCnd(simple.GetDB(), cnd)
}

func (this *{{.Name}}Service) Query(queries *simple.ParamQueries) (list []model.{{.Name}}, paging *simple.Paging) {
	return this.{{.Name}}Repository.Query(simple.GetDB(), queries)
}

func (this *{{.Name}}Service) Create(t *model.{{.Name}}) error {
	return this.{{.Name}}Repository.Create(simple.GetDB(), t)
}

func (this *{{.Name}}Service) Update(t *model.{{.Name}}) error {
	return this.{{.Name}}Repository.Update(simple.GetDB(), t)
}

func (this *{{.Name}}Service) Updates(id int64, columns map[string]interface{}) error {
	return this.{{.Name}}Repository.Updates(simple.GetDB(), id, columns)
}

func (this *{{.Name}}Service) UpdateColumn(id int64, name string, value interface{}) error {
	return this.{{.Name}}Repository.UpdateColumn(simple.GetDB(), id, name, value)
}

func (this *{{.Name}}Service) Delete(id int64) {
	this.{{.Name}}Repository.Delete(simple.GetDB(), id)
}

`
