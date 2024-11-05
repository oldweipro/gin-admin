package system

import (
	"fmt"
	"github.com/oldweipro/gin-admin/model/system/response"
	"github.com/oldweipro/gin-admin/pkg/app"
)

var AutoCodeMssql = new(autoCodeMssql)

type autoCodeMssql struct{}

// GetDB 获取数据库的所有数据库名
// Author [oldweipro](https://github.com/oldweipro)
func (s *autoCodeMssql) GetDB(businessDB string) (data []response.Db, err error) {
	var entities []response.Db
	sql := "select name AS 'database' from sysdatabases;"
	if businessDB == "" {
		err = app.DBClient.Raw(sql).Scan(&entities).Error
	} else {
		err = app.DbList[businessDB].Raw(sql).Scan(&entities).Error
	}
	return entities, err
}

// GetTables 获取数据库的所有表名
// Author [oldweipro](https://github.com/oldweipro)
func (s *autoCodeMssql) GetTables(businessDB string, dbName string) (data []response.Table, err error) {
	var entities []response.Table

	sql := fmt.Sprintf(`select name as 'table_name' from %s.DBO.sysobjects where xtype='U'`, dbName)
	if businessDB == "" {
		err = app.DBClient.Raw(sql).Scan(&entities).Error
	} else {
		err = app.DbList[businessDB].Raw(sql).Scan(&entities).Error
	}

	return entities, err
}

// GetColumn 获取指定数据库和指定数据表的所有字段名,类型值等
// Author [oldweipro](https://github.com/oldweipro)
func (s *autoCodeMssql) GetColumn(businessDB string, tableName string, dbName string) (data []response.Column, err error) {
	var entities []response.Column
	sql := fmt.Sprintf(`
SELECT
    sc.name AS column_name,
    st.name AS data_type,
    sc.max_length AS data_type_long,
    CASE
        WHEN pk.object_id IS NOT NULL THEN 1
        ELSE 0
    END AS primary_key,
    sc.column_id
FROM
    %s.sys.columns sc
JOIN
    sys.types st ON sc.user_type_id=st.user_type_id
LEFT JOIN
    %s.sys.objects so ON so.name='%s' AND so.type='U'
LEFT JOIN
    %s.sys.indexes si ON si.object_id = so.object_id AND si.is_primary_key = 1
LEFT JOIN
    %s.sys.index_columns sic ON sic.object_id = si.object_id AND sic.index_id = si.index_id AND sic.column_id = sc.column_id
LEFT JOIN
    %s.sys.key_constraints pk ON pk.object_id = si.object_id
WHERE
    st.is_user_defined=0 AND sc.object_id = so.object_id
ORDER BY
    sc.column_id
`, dbName, dbName, tableName, dbName, dbName, dbName)

	if businessDB == "" {
		err = app.DBClient.Raw(sql).Scan(&entities).Error
	} else {
		err = app.DbList[businessDB].Raw(sql).Scan(&entities).Error
	}

	return entities, err
}
