package codegen

import "html/template"

var repositoryTmpl = template.Must(template.New("repository").Parse(repositoryTmplCode))

const repositoryTmplCode = `
package repositories

import (
	"{{.PkgName}}/model"
	"github.com/mlogclub/simple"
	"github.com/jinzhu/gorm"
)

type {{.Name}}Repository struct {
}

func New{{.Name}}Repository() *{{.Name}}Repository {
	return &{{.Name}}Repository{}
}

func (this *{{.Name}}Repository) Get(db *gorm.DB, id int64) *model.{{.Name}} {
	ret := &model.{{.Name}}{}
	if err := db.First(ret, "id = ?", id).Error; err != nil {
		return nil
	}
	return ret
}

func (this *{{.Name}}Repository) Take(db *gorm.DB, where ...interface{}) *model.{{.Name}} {
	ret := &model.{{.Name}}{}
	if err := db.Take(ret, where...).Error; err != nil {
		return nil
	}
	return ret
}

func (this *{{.Name}}Repository) QueryCnd(db *gorm.DB, cnd *simple.QueryCnd) (list []model.{{.Name}}, err error) {
	err = cnd.DoQuery(db).Find(&list).Error
	return
}

func (this *{{.Name}}Repository) Query(db *gorm.DB, queries *simple.ParamQueries) (list []model.{{.Name}}, paging *simple.Paging) {
	queries.StartQuery(db).Find(&list)
    queries.StartCount(db).Model(&model.{{.Name}}{}).Count(&queries.Paging.Total)
	paging = queries.Paging
	return
}

func (this *{{.Name}}Repository) Create(db *gorm.DB, t *model.{{.Name}}) (err error) {
	err = db.Create(t).Error
	return
}

func (this *{{.Name}}Repository) Update(db *gorm.DB, t *model.{{.Name}}) (err error) {
	err = db.Save(t).Error
	return
}

func (this *{{.Name}}Repository) Updates(db *gorm.DB, id int64, columns map[string]interface{}) (err error) {
	err = db.Model(&model.{{.Name}}{}).Where("id = ?", id).Updates(columns).Error
	return
}

func (this *{{.Name}}Repository) UpdateColumn(db *gorm.DB, id int64, name string, value interface{}) (err error) {
	err = db.Model(&model.{{.Name}}{}).Where("id = ?", id).UpdateColumn(name, value).Error
	return
}

func (this *{{.Name}}Repository) Delete(db *gorm.DB, id int64) {
	db.Model(&model.{{.Name}}{}).Delete("id", id)
}

`
