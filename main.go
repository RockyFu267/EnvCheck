package main

import (
	ebf "EnvCheck/basefunc"
	ec "EnvCheck/controller"
	"encoding/json"
	"flag"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

var starttrole *string = flag.String("role", "master", "Use -role <master or client>")

func main() {
	//获取参数
	flag.Parse()
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
	resJson, _ := json.MarshalIndent(res, "", " ")
	log.Println(string(resJson))

	ebf.HostInfoList = append(ebf.HostInfoList, res)

	configTmp := ec.ReadConfig(pwdPath)
	posturlTmp := "http://" + configTmp.MasterIP + ":" + configTmp.MasterPort + "/env_info"
	log.Println(posturlTmp)
	if *starttrole != "master" && *starttrole != "client" {
		log.Println("role error")
		return
	}
	chScanTurnBool := make(chan bool)
	if *starttrole == "master" && configTmp.Mode == "http" {
		// 1.创建路由
		r := gin.Default()
		// 2.绑定路由规则，执行的函数
		// gin.Context，封装了request和response
		r.GET("/health_check/", func(c *gin.Context) {
			c.String(http.StatusOK, "hello World!")
		})
		r.POST("/env_info", ec.EnvInfo)
		go r.Run(":8282")
		time.Sleep(5 * time.Second)
		for _, v := range configTmp.Host {
			var tmpHost ec.HostPara = v
			go tmpHost.SSHClient(pwdPath, configTmp.RemotePath)
		}
		go func() {
			for {
				time.Sleep(10 * time.Second)
				checkListRes := ec.CheckInfoList()
				if checkListRes {
					chScanTurnBool <- true
				}
			}
		}()
		//等待结束信号
		for {
			select {
			case <-chScanTurnBool:
				ec.WriteRes()
				return
			case <-time.After(time.Duration(600 * time.Second)):
				log.Println("TimeOut")
				return
			}
		}
	}
	if *starttrole == "client" && configTmp.Mode == "http" {
		err = ec.PostLocalAction(res, posturlTmp)
		if err != nil {
			log.Println("Failed to Post data : ", err)
			// mesTmp := res.Meta.IP + ": There is an error uploading data from the main server"
			// ebf.CheckWarning = append(ebf.CheckWarning, mesTmp)
		}
		return
	}

}
