package model

type Tbl_dbarchiver_task struct {
	Taskid        string `json:"taskid" gorm:"column:taskid" form:"taskid"`
	Shost         string `json:"shost" gorm:"column:shost" form:"shost" binding:"required"`
	Sport         int    `json:"sport" gorm:"column:sport"  form:"sport" binding:"required"`
	Sschname      string `json:"sschname" gorm:"column:sschname" form:"sschname"  binding:"required"`
	Stablename    string `json:"stablename" gorm:"column:stablename" form:"stablename" binding:"required"`
	Dhost         string `json:"dhost" gorm:"column:dhost" form:"dhost" binding:"required"`
	Dport         int    `json:"dport" gorm:"column:dport" form:"dport" binding:"required"`
	Dschename     string `json:"dschename" gorm:"column:dschename" form:"dschename" binding:"required"`
	Dtabname      string `json:"dtabname" gorm:"column:dtabname" form:"dtabname" binding:"required"`
	Exdtype       int    `json:"exdtype" gorm:"column:exdtype" form:"exdtype"`
	Is_sharding   int    `json:"is_sharding" gorm:"column:is_sharding" form:"is_sharding"`
	Is_condition  string `json:"is_condition" gorm:"column:is_condition" form:"is_condition" binding:"required"`
	Is_store      int    `json:"is_store" gorm:"column:is_store" form:"is_store" binding:"required"`
	Schedule_time string `json:"schedule_time" gorm:"column:schedule_time" form:"schedule_time"`
	Is_scheduled  int    `json:"is_scheduled" gorm:"column:is_scheduled" form:"is_scheduled"`
	Subtask       string `json:"subtask" gorm:"column:subtask" form:"subtask" binding:"required"`
	Is_idx        string `json:"is_idx" gorm:"column:is_idx" form:"is_idx" `
	Approverby    string `json:"approverby" gorm:"column:approverby" form:"approverby" binding:"required"`
}

type Tbl_dbarchiver_task_log struct {
	//Task         Tbl_dbarchiver_task
	Taskid       string `json:"taskid" gorm:"column:taskid" form:"taskid"`
	Shost        string `json:"shost" gorm:"column:shost" form:"shost" binding:"required"`
	Sport        int    `json:"sport" gorm:"column:sport"  form:"sport" binding:"required"`
	Sschname     string `json:"sschname" gorm:"column:sschname" form:"sschname"  binding:"required"`
	Stablename   string `json:"stablename" gorm:"column:stablename" form:"stablename" binding:"required"`
	Dhost        string `json:"dhost" gorm:"column:dhost" form:"dhost" binding:"required"`
	Dport        int    `json:"dport" gorm:"column:dport" form:"dport" binding:"required"`
	Dschename    string `json:"dschename" gorm:"column:dschename" form:"dschename" binding:"required"`
	Dtabname     string `json:"dtabname" gorm:"column:dtabname" form:"dtabname" binding:"required"`
	Is_condition string `json:"is_condition" gorm:"column:is_condition" form:"is_condition" binding:"required"`
	Is_state     int
	Is_store     int
	Delete_count int    `json:"delete_count" gorm:"column:delete_count" form:"delete_count"`
	Insert_count int    `json:"insert_count" gorm:"column:insert_count" form:"insert_count"`
	Start_time   string `json:"start_time" gorm:"column:start_time" form:"start_time"`
	End_time     string `json:"end_time" gorm:"column:end_time" form:"end_time"`
}

//
func NewTbl_dbarchiver_task(taskid, shost string, sport int, sschname, stablename, dhost string, dport int, dschename, dtabname string, exdtype, is_sharding int, is_condition string, is_store int, schedule_time string, is_scheduled int, subtask, is_idx, approverby string) *Tbl_dbarchiver_task {
	p := &Tbl_dbarchiver_task{
		Taskid:        taskid,
		Shost:         shost,
		Sport:         sport,
		Sschname:      sschname,
		Stablename:    stablename,
		Dhost:         dhost,
		Dport:         dport,
		Dschename:     dschename,
		Dtabname:      dtabname,
		Exdtype:       exdtype,
		Is_sharding:   is_sharding,
		Is_condition:  is_condition,
		Is_store:      is_store,
		Schedule_time: schedule_time,
		Is_scheduled:  is_scheduled,
		Subtask:       subtask,
		Is_idx:        is_idx,
		Approverby:    approverby,
	}
	return p
}
