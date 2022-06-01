package model

//表名字tbl_add_info_task
//taskid, shost, sport, dbname, tablename, command_exe, cmd_exe,exe_type, cmd_idc, db_type, search_type, sqltext
type Tbl_add_ddl_task struct {
	Taskid      string
	Shost       string
	Sport       int
	Dbname      string //输入的执行数据库名字
	Tablename   string //输入的执行表名字
	Command_exe int    //选择执行DDL方式 1-GH 2原生alter
	Cmd_exe     int    //1执行DDL 2、添加schema
	Exe_type    int    // 1已知批量执行2自定义执行
	Cmd_idc     int    //1北京机房IDC  2阿里云数据库
	Db_type     int    //执行类型1会员2订单3支付4结算5老板通6自定义数据源
	Search_type int    //执行查找数据匹配方式1，精确查找2模糊查找
	Sqltext     string //输入执行的DDLsql

}
//taskid, shost, sport, dbname, tablename, command_exe, cmd_exe, exe_type, cmd_idc, db_type, search_type, sqltext
func NewTbl_add_ddl_task(taskid, shost string, sport int, dbname, tablename string, command_exe, cmd_exe, exe_type, cmd_idc, db_type, search_type int, sqltext string) *Tbl_add_ddl_task {
	return &Tbl_add_ddl_task{
		Taskid:      taskid,
		Shost:       shost,
		Sport:       sport,
		Dbname:      dbname,
		Tablename:   tablename,
		Command_exe: command_exe,
		Cmd_exe:     cmd_exe,
		Exe_type:    exe_type,
		Cmd_idc:     cmd_idc,
		Db_type:     db_type,
		Search_type: search_type,
		Sqltext:     sqltext,
	}
}

type dblist struct {
	Dbname    string
	Tablename string
}

func NewDblist(dbname, tablename string) *dblist {
	return &dblist{
		Dbname:    dbname,
		Tablename: tablename,
	}
}
