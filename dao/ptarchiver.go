package dao

import (
	"awesomeProject/db-monitorProject/model"
	"fmt"
)

func InsertPtarchiver(pt *model.Tbl_dbarchiver_task) (ptarchiveId int64, err error) {
	if pt == nil {
		return
	}
	fmt.Println("xxxxx:", pt.Taskid, pt.Shost, pt.Sport, pt.Sschname, pt.Stablename, pt.Dhost, pt.Dport, pt.Dschename, pt.Dtabname, pt.Exdtype, pt.Is_sharding, pt.Is_condition, pt.Is_store, pt.Schedule_time, pt.Is_scheduled, pt.Subtask, pt.Is_idx, pt.Approverby)
	sqlstr := "insert into tbl_dbarchiver_task(taskid,shost,sport,sschname,stablename,dhost,dport,dschename,dtabname,exdType,is_sharding,is_condition,is_store,schedule_time,is_scheduled,subtask,is_idx,approverby)values(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)"
	ret, err := Db.Exec(sqlstr, pt.Taskid, pt.Shost, pt.Sport, pt.Sschname, pt.Stablename, pt.Dhost, pt.Dport, pt.Dschename, pt.Dtabname, pt.Exdtype, pt.Is_sharding, pt.Is_condition, pt.Is_store, pt.Schedule_time, pt.Is_scheduled, pt.Subtask, pt.Is_idx, pt.Approverby)

	if err != nil {
		fmt.Println("AAAA", err)
		return
	}
	ptarchiveId, _ = ret.LastInsertId()
	return

}

//展示提交为执行的工单
func GetPtarchiveList() (ptarchiverList []*model.Tbl_dbarchiver_task, err error) {
	sqlstr := "select taskid,shost,sport,sschname,stablename,dhost,dport,dschename,dtabname,exdtype,is_sharding,is_condition,is_store,schedule_time,is_scheduled,approverby,subtask,is_idx from tbl_dbarchiver_task where is_state=0 order by id desc"
	err = Db.Select(&ptarchiverList, sqlstr)
	if err != nil {
		fmt.Println("select failed")
		return
	}
	return
}

//为执行得工单提前终止
func UpatePtarchiver(taskId string) (ptarchiverEffect int64, err error) {
	if taskId == "" {
		return
	}
	sqlstr := `update tbl_dbarchiver_task set is_state=4 where taskid=?`
	ret, err := Db.Exec(sqlstr, taskId)
	if err != nil {
		return
	}
	ptarchiverEffect, err = ret.RowsAffected()
	if err != nil {
		return
	}
	return

}

//工单审核处理
func AutoPtarchiver(ptuser string) (ptarchiver int, err error) {
	ptarchiver = 0
	//查询当前用户属于审核者具备得权限
	sqlstr := `select approverby  tbl_dbarchiver_task from where approverby='?'`
	err = Db.Get(&ptarchiver, sqlstr, ptuser)
	if err != nil {
		ptarchiver = 0
		return
	}
	if ptarchiver == 0 {
		ptarchiver = 0
		return
	}
	return
}

func GetPtarchiverTaskId(taskId string) (ptarchiverList *model.Tbl_dbarchiver_task, err error) {
	if taskId == "" {
		return
	}
	sqlstr := ` select taskid,shost,sport,sschname,stablename,dhost,dport,dschename,dtabname,exdtype,
	is_sharding,is_condition,is_store,schedule_time,is_scheduled,approverby,subtask,is_idx from 
	tbl_dbarchiver_task where taskid=? and is_state=0 order by id desc`
	err = Db.Get(&ptarchiverList, sqlstr, taskId)
	if err != nil {
		return

	}
	return
}

//获取需要执行的工单
func GetPtarchiverExecList() (ptarchiverList []*model.Tbl_dbarchiver_task, err error) {
	sqlstr := ` select taskid,shost,sport,sschname,stablename,dhost,dport,dschename,dtabname,exdtype,
	is_sharding,is_condition,is_store,schedule_time,is_scheduled,approverby,subtask,is_idx from 
	tbl_dbarchiver_task where  is_state=0 order by id desc`
	err = Db.Select(&ptarchiverList, sqlstr)
	if err != nil {
		return

	}
	return
}
