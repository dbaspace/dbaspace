package main

import (
	rddl "awesomeProject/db-monitorProject/controller/ddlschema"
	hpt "awesomeProject/db-monitorProject/controller/ptarchiver"
	"awesomeProject/db-monitorProject/dao"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	_, err := dao.InitDB()
	if err != nil {
		panic(err)
	}
	//加载静态文件
	//r.Static("/static", "./static")
	//加载模板
	r.LoadHTMLGlob("view/*.html")
	r.GET("/index.html", hpt.IndexPtarchiverList)
	r.POST("/addtask.html", hpt.PutPtarchiver)
	r.GET("/addtask.html", hpt.PutPtarchiver)
	r.GET("/taskddl.html", rddl.Taskddl)
	r.POST("/taskddl.html", rddl.Taskddl)
	fmt.Println("start")
	r.Run(":8000")
}
