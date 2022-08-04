package controller

import (
	"net/http"

	ebf "EnvCheck/basefunc"

	"github.com/gin-gonic/gin"
)

func EnvInfo(c *gin.Context) {
	// 表单发送 name=card, job=phper
	var postData ebf.HostInfo
	if err := c.ShouldBindJSON(&postData); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":  "400",
			"error": "formError",
		})
		return
	}
	//获取IP
	postData.Meta.PostIP = GetRequestIP(c)
	//debug
	if postData.Meta.Version != ebf.Version || postData.HostName.HostNameStr == "" {
		c.JSON(http.StatusOK, gin.H{
			"code":  "400",
			"error": "infoError",
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"code": "200",
	})
	if postData.Meta.IP != postData.Meta.PostIP {
		mesTmp := postData.HostName.HostNameStr + " IP is inconsistent with the requested IP"
		ebf.CheckWarning = append(ebf.CheckWarning, mesTmp)
		//目前警告还未打印
	}
	//把上报的数据添加至结果列表
	ebf.HostInfoList = append(ebf.HostInfoList, postData)
	writeRes()
}

func GetRequestIP(c *gin.Context) string {
	reqIP := c.ClientIP()
	if reqIP == "::1" {
		reqIP = "127.0.0.1"
	}
	return reqIP
}
