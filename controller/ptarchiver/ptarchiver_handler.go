package controller

import (
	"awesomeProject/db-monitorProject/model"
	pt "awesomeProject/db-monitorProject/service/ptarchiver"
	"net/http"
	"strconv"

	uuid "github.com/go.uuid"

	"github.com/gin-gonic/gin"
)

//访问主页的控制器
func IndexPtarchiverList(c *gin.Context) {

	ptarchiverlist, err := pt.GetALLPtarchiverList()
	if err != nil {
		c.HTML(http.StatusInternalServerError, "500.html", gin.H{"err_msg": err})
		return
	}
	c.HTML(http.StatusOK, "index.html", gin.H{
		"ptarchiver_list": ptarchiverlist,
	})
}

//数据归档工单提交
func PutPtarchiver(c *gin.Context) {
	if c.Request.Method == "GET" {
		//c.JSON(200, gin.H{"code": "get"})
		c.HTML(http.StatusOK, "addtask.html", nil)
	} else {
		uid := uuid.Must(uuid.NewV4())
		taskid := uid.String()
		shost := c.PostForm("shost")
		sport, _ := strconv.Atoi(c.PostForm("sport"))
		sschname := c.PostForm("sschname")
		stablename := c.PostForm("stablename")
		dhost := c.PostForm("dhost")
		dport, _ := strconv.Atoi(c.PostForm("dport"))
		dschename := c.PostForm("dschename")
		dtabname := c.PostForm("dtabname")
		exdtype, _ := strconv.Atoi(c.PostForm("exdtype"))
		is_condition := c.PostForm("is_condition")
		is_store, _ := strconv.Atoi(c.PostForm("is_store"))
		schedule_time := c.PostForm("schedule_time")
		is_sharding, _ := strconv.Atoi(c.PostForm("is_sharding"))
		is_scheduled, _ := strconv.Atoi(c.PostForm("is_scheduled"))
		subtask := "dba"
		is_idx := c.PostForm("is_idx")
		approverby := c.PostForm("approverby")
		pd := model.NewTbl_dbarchiver_task(taskid, shost, sport, sschname, stablename, dhost, dport, dschename, dtabname, exdtype, is_sharding, is_condition, is_store, schedule_time, is_scheduled, subtask, is_idx, approverby)

		ptarchiverId, err := pt.PutPtarchiver(pd)
		if err != nil {
			c.HTML(http.StatusInternalServerError, "500.html", gin.H{"err_msg": "empty"})
			return
		}
		if ptarchiverId != 0 {
			//c.HTML(http.StatusOK,"",gin.H{"code":"ok"})
			//c.JSON(200, gin.H{"code": "提交成功"})
			c.HTML(http.StatusOK, "index.html", gin.H{"code": "提交成功"})
		} else {
			c.HTML(http.StatusOK, "index.html", gin.H{"code": "请重新提交"})
		}
	}
}
