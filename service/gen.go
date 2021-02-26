package service

import (
	"codegen/util"
	"os"
	"text/template"
)

type Columns struct {
	TableCatalog           string `gorm:"column:TABLE_CATALOG" json:"table_catalog"`
	TableSchema            string `gorm:"column:TABLE_SCHEMA" json:"table_schema"`
	TableName              string `gorm:"column:TABLE_NAME" json:"table_name"`
	ColumnName             string `gorm:"column:COLUMN_NAME" json:"column_name"`
	OrdinalPosition        uint32 `gorm:"column:ORDINAL_POSITION" json:"ordinal_position"`
	ColumnDefault          string `gorm:"column:COLUMN_DEFAULT" json:"column_default"`
	IsNullable             string `gorm:"column:IS_NULLABLE" json:"is_nullable"`
	DataType               string `gorm:"column:DATA_TYPE" json:"data_type"`
	CharacterMaximumLength uint64 `gorm:"column:CHARACTER_MAXIMUM_LENGTH" json:"character_maximum_length"`
	CharacterOctetLength   uint64 `gorm:"column:CHARACTER_OCTET_LENGTH" json:"character_octet_length"`
	NumericPrecision       uint64 `gorm:"column:NUMERIC_PRECISION" json:"numeric_precision"`
	NumericScale           uint64 `gorm:"column:NUMERIC_SCALE" json:"numeric_scale"`
	DatetimePrecision      uint32 `gorm:"column:DATETIME_PRECISION" json:"datetime_precision"`
	CharacterSetName       string `gorm:"column:CHARACTER_SET_NAME" json:"character_set_name"`
	CollationName          string `gorm:"column:COLLATION_NAME" json:"collation_name"`
	ColumnType             string `gorm:"column:COLUMN_TYPE" json:"column_type"`
	ColumnKey              string `gorm:"column:COLUMN_KEY" json:"column_key"`
	Extra                  string `gorm:"column:EXTRA" json:"extra"`
	Privileges             string `gorm:"column:PRIVILEGES" json:"privileges"`
	ColumnComment          string `gorm:"column:COLUMN_COMMENT" json:"column_comment"`
	GenerationExpression   string `gorm:"column:GENERATION_EXPRESSION" json:"generation_expression"`
}

type Field struct {
	Name       string
	Type       string
	Tag        string
	Comment    string
	ColumnName string
}

type Data struct {
	HasTime    bool
	StructName string
	TableName  string
	Fields     []Field
}

var gc *GenCode

type GenCode struct {
	ColumnList []Columns
	Data       Data
}

func NewGenCode() *GenCode {
	return &GenCode{Data: Data{HasTime: false, StructName: util.SnakeToCamel(Cmd.TableName), TableName: Cmd.TableName}}
}

func (g *GenCode) Gen() (result string, err error) {
	g.Data.TableName = Cmd.TableName
	err = g.selectColumns(Cmd.TableName)
	if err != nil {
		panic(err)
	}

	err = g.assemblyFields()
	if err != nil {
		panic(err)
	}

	g.generateModel()
	g.generateDao()
	return
}

func (g *GenCode) selectColumns(tableName string) (err error) {
	return db.Where(&Columns{TableName: tableName}).Find(&g.ColumnList).Error
}

func (g *GenCode) assemblyFields() (err error) {
	for i := 0; i < len(g.ColumnList); i++ {
		g.Data.Fields = append(g.Data.Fields, g.assemblyField(g.ColumnList[i]))
	}
	return
}

func (g *GenCode) assemblyField(d Columns) (f Field) {
	f.Name = util.SnakeToCamel(d.ColumnName)
	f.Type = typeMap[d.DataType]
	f.Tag = "`gorm:\"column:" + d.ColumnName + "\" json:\"" + d.ColumnName + "\"`"
	f.Comment = d.ColumnComment
	f.ColumnName = d.ColumnName
	return
}

func (g *GenCode) generateModel() {
	t, err := template.New("model").Parse(modelTemplate)
	if err != nil {
		panic(err)
	}
	err = t.Execute(os.Stdout, g.Data)
	if err != nil {
		panic(err)
	}
}

func (g *GenCode) generateDao() {
	t, err := template.New("dao").Parse(DaoTemplate)
	if err != nil {
		panic(err)
	}
	err = t.Execute(os.Stdout, g.Data)
	if err != nil {
		panic(err)
	}
}
