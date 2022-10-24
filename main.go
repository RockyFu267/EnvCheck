package main

import (
	ebf "EnvCheck/basefunc"
	ec "EnvCheck/controller"
	"encoding/json"
	"flag"
	"log"
	"mime"
	"net/http"
	"os"
	"time"

	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

var starttrole *string = flag.String("role", "", "Use -role <master or client>")
var starttshow *string = flag.String("show", "", "Use -show Browse web pages")
var disktest *string = flag.String("disktest", "", "Use -disktest Show Disk IOPS")

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
	//web浏览模式
	if len(*starttshow) > 0 {
		jsonTmp, err := ec.ReadResJson(*starttshow)
		if err != nil {
			log.Println(err)
			return
		}
		// 1.创建路由
		r := gin.Default()
		// 2.绑定路由规则，执行的函数
		// gin.Context，封装了request和response
		r.GET("/health_check/", func(c *gin.Context) {
			c.String(http.StatusOK, "HelloWorld")
		})
		r.GET("/getinfo", func(c *gin.Context) {
			c.JSON(http.StatusOK, jsonTmp)
		})
		_ = mime.AddExtensionType(".js", "application/javascript")
		r.Use(static.Serve("/", static.LocalFile("./dist", true)))

		r.Run(":" + configTmp.MasterPort)
		return
	}
	chScanTurnBool := make(chan bool)
	//判断启动参数 是主服务还是客户端 是否是单机模式
	//如果是单机模式
	if *starttrole == "" && configTmp.Mode == "http" {
		if *disktest != "" {
			_, err := ebf.CmdFioTest(*disktest)
			if err != nil {
				log.Println("Get DiskTestInfo ERROR: ", err)
			}
		}
		ec.WriteRes()
		ec.WriteResImage()
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
				ec.WriteResImage()
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
