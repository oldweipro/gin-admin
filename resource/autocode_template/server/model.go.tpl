// 自动生成模板{{.StructName}}
package {{.Package}}

import (
	"github.com/oldweipro/gin-admin/global"
	{{ if .HasTimer }}"time"{{ end }}
	{{ if .HasFile }}"gorm.io/datatypes"{{ end }}
)

// {{.StructName}} 结构体
type {{.StructName}} struct {
      global.MODEL {{- range .Fields}}
            {{- if eq .FieldType "enum" }}
      {{.FieldName}}  string `json:"{{.FieldJson}}" form:"{{.FieldJson}}" gorm:"column:{{.ColumnName}};type:enum({{.DataTypeLong}});comment:{{.Comment}};"`
            {{- else if eq .FieldType "picture" }}
      {{.FieldName}}  string `json:"{{.FieldJson}}" form:"{{.FieldJson}}" gorm:"column:{{.ColumnName}};comment:{{.Comment}};{{- if .DataTypeLong -}}size:{{.DataTypeLong}};{{- end -}}"`
            {{- else if eq .FieldType "file" }}
      {{.FieldName}}  datatypes.JSON `json:"{{.FieldJson}}" form:"{{.FieldJson}}" gorm:"column:{{.ColumnName}};comment:{{.Comment}};{{- if .DataTypeLong -}}size:{{.DataTypeLong}};{{- end -}}"`
            {{- else if ne .FieldType "string" }}
      {{.FieldName}}  *{{.FieldType}} `json:"{{.FieldJson}}" form:"{{.FieldJson}}" gorm:"column:{{.ColumnName}};comment:{{.Comment}};{{- if .DataTypeLong -}}size:{{.DataTypeLong}};{{- end -}}"`
            {{- else }}
      {{.FieldName}}  {{.FieldType}} `json:"{{.FieldJson}}" form:"{{.FieldJson}}" gorm:"column:{{.ColumnName}};comment:{{.Comment}};{{- if .DataTypeLong -}}size:{{.DataTypeLong}};{{- end -}}"`
            {{- end }} {{- end }}
      {{- if .AutoCreateResource }}
      CreatedBy  uint   `gorm:"column:created_by;comment:创建者"`
      UpdatedBy  uint   `gorm:"column:updated_by;comment:更新者"`
      DeletedBy  uint   `gorm:"column:deleted_by;comment:删除者"`
      {{- end}}
}

{{ if .TableName }}
// TableName {{.StructName}} 表名
func ({{.StructName}}) TableName() string {
  return "{{.TableName}}"
}
{{ end }}
