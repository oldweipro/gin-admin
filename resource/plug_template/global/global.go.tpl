package global

{{- if .HasGlobal }}

import "github.com/oldweipro/gin-admin/plugin/{{ .Snake}}/config"

var GlobalConfig = new(config.{{ .PlugName}})
{{ end -}}