package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	ebf "EnvCheck/basefunc"

	"github.com/gin-gonic/gin"
)

func EnvInfo(c *gin.Context) {
	// 表单发送 name=card, job=phper
	var postData ebf.HostInfo
	// err := json.NewDecoder(c.Request.Body).Decode(&postData)
	// if err != nil {
	// 	c.JSON(http.StatusOK, gin.H{
	// 		"message": "fuck",
	// 	})
	// 	fmt.Println("error---------------")
	// 	return
	// }
	if err := c.ShouldBindJSON(&postData); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"Code":   "400",
			"Reason": "formError",
		})
		return
	}
	//获取IP
	iptmp := GetRequestIP(c)
	if postData.Meta.Version != ebf.Version || postData.HostName.HostNameStr == "" {
		c.JSON(http.StatusOK, gin.H{
			"Code":   "400",
			"Reason": "infoError",
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"Code": "200",
	})

	resJson, _ := json.MarshalIndent(postData, "", " ")
	fmt.Println(string(resJson), iptmp)

}

func GetRequestIP(c *gin.Context) string {
	reqIP := c.ClientIP()
	if reqIP == "::1" {
		reqIP = "127.0.0.1"
	}
	return reqIP
}
