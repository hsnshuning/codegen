package service


var modelTemplate = `
package model

{{if .HasTime}}import "time"
{{end}}
type {{.StructName}} struct {
{{range .Fields}}	{{.Name}}	{{.Type}}	{{.Tag}}	//{{.Comment}}
{{end}}
}

func ({{.StructName}}) TableName() string {
	return "{{.TableName}}"
}
`