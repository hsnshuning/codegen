package service

var typeMap = map[string]string{
	"bigint": "int64",
	"int":"int32",
	"smallint":"int",
	"tinyint":"int8",
	"double":"float64",
	"float":"float64",
	"date":"time.Time",
	"datetime":"time.Time",
	"time":"time.Time",
	"timestamp":"time.Time",
	"bit":"int",
	"bool":"bool",
	"enum":"string",
	"set":"string",
	"char":"string",
	"varchar":"string",
}

var modelTemplate = `
package model

{{if .HasTime}}import "time"
{{end}}
type {{.StructName}} struct {`+"\n"+
	"{{range .Fields}}	{{.Name}}\t{{.Type}}\t{{.Tag}}\t//{{.Comment}}\n{{end}}"+
	"}\n"+
`
func ({{.StructName}}) TableName() string {
	return "{{.TableName}}"
}
`


var DaoTemplate = `
package dao

import (
	"github.com/go-gorm/gorm"
	"model"
	"sync"
)

var _{{.StructName}} *{{.StructName}}

var _once{{.StructName}} sync.Once

type {{.StructName}} struct {}

func New{{.StructName}} () (result *{{.StructName}}) {
	if _{{.StructName}} == nil {
		_once{{.StructName}}.Do(func () {
			_{{.StructName}} = new({{.StructName}})
		})
	}
	return _{{.StructName}}
}

func (*{{.StructName}}) Insert (db gorm.DB, data *model.{{.StructName}}) (err error) {
	err = db.Create(data).Error
	return 
}

func (*{{.StructName}}) GetOne (db gorm.DB, id int64) (result *model.{{.StructName}}, err error) {
	tmp := new(model.{{.StructName}})
	err = db.Where(where).First(tmp).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			err = nil
		}
		return
	}
	result = tmp
	err = db.First(result, id).Error
	return 
}

func (*{{.StructName}}) GetList (db gorm.DB, where *model.{{.StructName}}, limit int64, offset int64) (result []*model.{{.StructName}}, err error) {
	if limit <= 0 {
		limit = 100
	}
	err = db.Where(where).Limit(limit).Offset(offset).Find(&result).Error
	return
}

func (*{{.StructName}}) Update (db gorm.DB, where *model.{{.StructName}}, update *model.{{.StructName}}) (err error) {
	if where == nil || update == nil {
		return errors.New("where or update is nil")
	}
	return db.Table(where.TableName()).Where(where).Updates(update).Error
}
`
