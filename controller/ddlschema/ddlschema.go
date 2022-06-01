package controller

import (
	"awesomeProject/db-monitorProject/model"
	dadd "awesomeProject/db-monitorProject/service/ddlschema"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/syohex/go-texttable"

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
		sqlp := c.PostForm("sqltext")
		sqltext := strings.Join(strings.Fields(sqlp), " ")
		ddl := model.NewTbl_add_ddl_task(taskid, shost, sport, dbname, tablename, command_exe, cmd_exe, exe_type, cmd_idc, db_type, search_type, sqltext)
		fmt.Println("add task info", taskid, shost, sport, dbname, tablename, command_exe, cmd_exe, exe_type, cmd_idc, db_type, search_type, sqltext)
		res, err := dadd.Putddlschema(ddl)
		if err != nil {
			c.HTML(http.StatusInternalServerError, "500.html", gin.H{"err_msg": err})
		}
		if res != 0 {
			c.HTML(http.StatusOK, "taskddl.html", gin.H{"code": "提交成功"})
		} else {
			c.HTML(http.StatusOK, "taskddl.html", gin.H{"code": "请重新提交"})
		}

	}
}

func GoinceptionChecksql(c *gin.Context) {
	if c.Request.Method == "GET" {
		//c.JSON(200, gin.H{"code": "get"})
		c.HTML(http.StatusOK, "taskddl.html", nil)
	} else {
		sqlp := c.PostForm("sqltext")
		sqltext := strings.Join(strings.Fields(sqlp), " ")
		rows, err := dadd.InceptionCheckSQL(sqltext)
		defer rows.Close()
		tbl := &texttable.TextTable{}
		if err != nil {
			log.Fatal(err)
			c.HTML(http.StatusInternalServerError, "500.html", gin.H{"err_msg": err})
		}
		cols, err := rows.Columns()
		if err != nil {
			log.Fatalln(err)
		}
		tbl.SetHeader("编号", cols[1], cols[2], cols[3], cols[4], cols[5], cols[6], cols[7], cols[8], cols[9], cols[10], cols[11])
		for rows.Next() {
			var order_id, affected_rows, stage, error_level, stage_status, error_message, sql, sequence, backup_dbname, execute_time, sqlsha1, backup_time []uint8
			err = rows.Scan(&order_id, &stage, &error_level, &stage_status, &error_message, &sql, &affected_rows, &sequence, &backup_dbname, &execute_time, &sqlsha1, &backup_time)
			if err != nil {
				fmt.Println("scan err:", err)
			}
			tbl.AddRow(string(order_id), string(affected_rows), string(stage), string(error_level), string(stage_status), string(error_message), string(sql), string(sequence), string(backup_dbname), string(execute_time))
			//fmt.Println(string(order_id), string(affected_rows), string(stage), string(error_level), string(stage_status), string(error_message), string(sql), string(sequence), string(backup_dbname), string(execute_time))
		}

		//c.JSON(http.StatusOK, gin.H{"authdata": tbl.Draw()})
		c.HTML(http.StatusOK, "taskddl.html", gin.H{"authdata": tbl.Draw()})
	}
}
