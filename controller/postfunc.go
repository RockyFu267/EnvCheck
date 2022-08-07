package controller

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

//PostAction 创建post请求
func PostLocalAction(InputData interface{}, url string) error {
	//转成[]byte
	bytesData, err := json.Marshal(InputData)
	if err != nil {
		log.Println("解析结构体失败: ", err)
		return err
	}

	reader := bytes.NewReader(bytesData)
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	//requeset
	req, err := http.NewRequest("POST", url, reader)
	if err != nil {
		log.Println("post url error1: ", err)
		return err
	}
	req.Header.Set("Content-Type", "application/json;charset=UTF-8")
	req.Header.Add("Connection", "close")
	client := &http.Client{Transport: tr, Timeout: time.Second * 5}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("post url error2: ", err)
		return err
	}
	resp.Header.Add("Connection", "close")
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("post request error1: ", err)
	}
	defer resp.Body.Close()
	var res ResData
	if resjson := json.Unmarshal([]byte(data), &res); resjson != nil {
		log.Println("post request error2: ", err)
		return resjson
	}

	return nil

}
