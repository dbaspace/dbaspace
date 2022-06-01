package service

import (
	"awesomeProject/db-monitorProject/model"
)

func CheckArgs(pt *model.Tbl_dbarchiver_task) (flage bool) {
	if pt.Dhost == "" || pt.Dport == 0 || pt.Dschename == "" || pt.Dtabname == "" {
		flage = false
	}

	return
}
