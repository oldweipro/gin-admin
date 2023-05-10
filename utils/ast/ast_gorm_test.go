package ast

import (
	"github.com/oldweipro/gin-admin/global"
	"github.com/oldweipro/gin-admin/model/example"
	"testing"
)

const A = 123

func TestAddRegisterTablesAst(t *testing.T) {
	AddRegisterTablesAst("D:\\gin-vue-admin\\server\\utils\\ast_test.go", "Register", "test", "testDB", "testModel")
}

func Register() {
	test := global.GetGlobalDBByDBName("test")
	test.AutoMigrate(example.ExaFile{})
}
