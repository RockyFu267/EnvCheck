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

var starttrole *string = flag.String("role", "", "Use -role <master or client>")

func main() {
	//获取参数
	flag.Parse()
	//检查参数合法性
	if *starttrole != "master" && *starttrole != "client" && *starttrole != "" {
		log.Println("role error")
		return
	}
	pwdPath, err := os.Getwd()
	if err != nil {
		log.Println("Get pwdPATH ERROR: ", err)
		return
	}
	log.Println("当前路径：", pwdPath)
	//获取本机信息
	res, err := ebf.GetHostInfo()
	if err != nil {
		log.Println("Get OsInfo ERROR: ", err)
		return
	}
	resJson, _ := json.MarshalIndent(res, "", " ")
	log.Println(string(resJson))

	ebf.HostInfoList = append(ebf.HostInfoList, res)

	configTmp := ec.ReadConfig(pwdPath)
	posturlTmp := "http://" + configTmp.MasterIP + ":" + configTmp.MasterPort + "/env_info"
	chScanTurnBool := make(chan bool)
	//判断启动参数 是主服务还是客户端 是否是单机模式
	//如果是单机模式
	if *starttrole == "" && configTmp.Mode == "http" {
		ec.WriteRes()
		return
	}
	//如果是主服务
	if *starttrole == "master" && configTmp.Mode == "http" {
		// 1.创建路由
		r := gin.Default()
		// 2.绑定路由规则，执行的函数
		// gin.Context，封装了request和response
		r.GET("/health_check/", func(c *gin.Context) {
			c.String(http.StatusOK, "HelloWorld")
		})
		r.POST("/env_info", ec.EnvInfo)
		go r.Run(":" + configTmp.MasterPort)
		time.Sleep(5 * time.Second)
		//并发远程操作命令
		for _, v := range configTmp.Host {
			var tmpHost ec.HostPara = v
			go tmpHost.SSHClient(pwdPath, configTmp.RemotePath, configTmp.MasterIP, configTmp.MasterPort)
		}
		go func() {
			for {
				time.Sleep(10 * time.Second)
				checkListRes := ec.CheckInfoList()
				//检查客户端执行结果,如果完成(包括失败的)
				if checkListRes {
					chScanTurnBool <- true
				}
			}
		}()
		//等待结束信号
		for {
			select {
			//收到完成信号
			case <-chScanTurnBool:
				resJsonList, _ := json.MarshalIndent(ebf.HostInfoList, "", " ")
				//打印总的结果
				log.Println(string(resJsonList))
				ec.WriteRes()
				return
			case <-time.After(time.Duration(600 * time.Second)):
				log.Println("TimeOut")
				return
			}
		}
	}
	//如果是客户端
	if *starttrole == "client" && configTmp.Mode == "http" {
		err = ec.PostLocalAction(res, posturlTmp)
		if err != nil {
			log.Println("Failed to Post data : ", err)
		}
		return
	}

}
