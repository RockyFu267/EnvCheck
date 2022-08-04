package controller

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

//PostAction 创建post请求
func PostLocalAction(InputData interface{}, url string) error {
	//转成[]byte
	bytesData, err := json.Marshal(InputData)
	if err != nil {
		log.Println("解析结构体失败", err)
		return err
	}

	reader := bytes.NewReader(bytesData)
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	//requeset
	req, err := http.NewRequest("POST", url, reader)
	if err != nil {
		log.Println(err)
		return err
	}
	req.Header.Set("Content-Type", "application/json;charset=UTF-8")
	req.Header.Add("Connection", "close")
	client := &http.Client{Transport: tr}
	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return err
	}
	resp.Header.Add("Connection", "close")
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
	}
	defer resp.Body.Close()
	var res ResData
	if resjson := json.Unmarshal([]byte(data), &res); resjson != nil {
		log.Println(err)
		return resjson
	}

	return nil

}
