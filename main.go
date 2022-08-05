package main

import (
	ebf "EnvCheck/basefunc"
	ec "EnvCheck/controller"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	pwdPath, err := os.Getwd()
	if err != nil {
		log.Println("Get pwdPATH ERROR: ", err)
		//***结束***
		return
	}
	log.Println("当前路径：", pwdPath)
	//获取本机信息
	res, err := ebf.GetHostInfo()
	if err != nil {
		log.Println("Get OsInfo ERROR: ", err)
		//***结束***
		return
	}
	ebf.HostInfoList = append(ebf.HostInfoList, res)

	configTmp := ec.ReadConfig(pwdPath)
	posturlTmp := "http://" + configTmp.MasterIP + ":" + configTmp.MasterPort + "/env_info"
	log.Println(posturlTmp)
	if configTmp.Role == "master" && configTmp.Mode == "http" {
		// 1.创建路由
		r := gin.Default()
		// 2.绑定路由规则，执行的函数
		// gin.Context，封装了request和response
		r.GET("/health_check/", func(c *gin.Context) {
			c.String(http.StatusOK, "hello World!")
		})
		r.POST("/env_info", ec.EnvInfo)
		r.Run(":8282")
		return
	}
	if configTmp.Role == "client" && configTmp.Mode == "http" {
		err = ec.PostLocalAction(res, posturlTmp)
		if err != nil {
			log.Println("Failed to Post data : ", err)
			// mesTmp := res.Meta.IP + ": There is an error uploading data from the main server"
			// ebf.CheckWarning = append(ebf.CheckWarning, mesTmp)
		}
		return
	}

}
