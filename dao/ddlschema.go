package dao

import (
	"awesomeProject/db-monitorProject/model"
	//"fmt"
	"strings"
)

func Insertddlschema(ddl *model.Tbl_add_ddl_task) (res int64, err error) {
	if ddl == nil {
		return
	}
	sqltext := strings.Join(strings.Fields(ddl.Sqltext), " ")
	sqlstr := "insert into tbl_add_ddl_task(taskid,shost,sport,dbname,tablename, command_exe,cmd_exe,exe_type,cmd_idc,db_type,search_type,sqltext)values(?,?,?,?,?,?,?,?,?,?,?,?)"
	r, err := Db.Exec(sqlstr, ddl.Taskid, ddl.Shost, ddl.Sport, ddl.Dbname, ddl.Tablename, ddl.Command_exe, ddl.Cmd_exe, ddl.Exe_type, ddl.Cmd_idc, ddl.Db_type, ddl.Search_type, sqltext)
	if err != nil {
		return
	}
	res, err = r.RowsAffected()
	if err != nil {
		return
	}

	return

}
