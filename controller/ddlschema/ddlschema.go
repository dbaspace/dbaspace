package controller

import (
	"awesomeProject/db-monitorProject/model"
	dadd "awesomeProject/db-monitorProject/service/ddlschema"
	"fmt"
	"net/http"
	"strconv"

	uuid "github.com/go.uuid"

	"github.com/gin-gonic/gin"
)

//数据归档工单提交
func Taskddl(c *gin.Context) {
	if c.Request.Method == "GET" {
		//c.JSON(200, gin.H{"code": "get"})
		c.HTML(http.StatusOK, "taskddl.html", nil)
	} else {
		uid := uuid.Must(uuid.NewV4())
		taskid := uid.String()
		shost := c.PostForm("shost")
		sport, _ := strconv.Atoi(c.PostForm("sport"))
		dbname := c.PostForm("dbname")
		tablename := c.PostForm("tablename")
		command_exe, _ := strconv.Atoi(c.PostForm("command_exe"))
		cmd_exe, _ := strconv.Atoi(c.PostForm("cmd_exe"))
		cmd_idc, _ := strconv.Atoi(c.PostForm("cmd_idc"))
		exe_type, _ := strconv.Atoi(c.PostForm("exe_type"))
		db_type, _ := strconv.Atoi(c.PostForm("db_type"))
		search_type, _ := strconv.Atoi(c.PostForm("search_type"))
		sqltext := c.PostForm("sqltext")
		ddl := model.NewTbl_add_ddl_task(taskid, shost, sport, dbname, tablename, command_exe, cmd_exe, exe_type, cmd_idc, db_type, search_type, sqltext)
		fmt.Println("add task info", taskid, shost, sport, dbname, tablename, command_exe, cmd_exe, exe_type, cmd_idc, db_type, search_type, sqltext)
		res, err := dadd.Putddlschema(ddl)
		if err != nil {
			c.HTML(http.StatusInternalServerError, "500.html", gin.H{"err_msg": err})
		}
		if res != 0 {
			c.HTML(http.StatusOK, "index.html", gin.H{"code": "提交成功"})
		} else {
			c.HTML(http.StatusOK, "index.html", gin.H{"code": "请重新提交"})
		}

	}
}
