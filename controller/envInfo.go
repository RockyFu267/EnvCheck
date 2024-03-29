package controller

import (
	"net/http"

	ebf "EnvCheck/basefunc"

	"github.com/gin-gonic/gin"
)

func EnvInfo(c *gin.Context) {
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
	if postData.Meta.Version != ebf.Version || postData.HostName.HostNameStr == "" {
		c.JSON(http.StatusOK, gin.H{
			"code":  "400",
			"error": "infoError",
		})
		return
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

	//把成功上传的记录的加入 postlist
	var tmpPostInfo ebf.PostInfo
	tmpPostInfo.PostIP = postData.Meta.IP
	tmpPostInfo.Res = false
	ebf.PostInfoList = append(ebf.PostInfoList, tmpPostInfo)
}

func GetRequestIP(c *gin.Context) string {
	reqIP := c.ClientIP()
	if reqIP == "::1" {
		reqIP = "127.0.0.1"
	}
	return reqIP
}
