package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
	//"github.com/bndr/gojenkins"
)


type Res struct {
	Code int `json:"code"`
	Content string `json:"content"`
}


//封装GET请求，返回http状态码
func Get(url string) (code int) {
	client := http.Client{Timeout: 5 * time.Second}
	resp, error := client.Get(url)
	if error != nil {
		fmt.Println("readyCheck get error: ", error)
		return 3001
	}
	defer resp.Body.Close()
	code = resp.StatusCode
	return
}

//封装POST请求 url:请求地址，data:POST请求提交的数据,contentType:请求体格式，如：application/json content:请求放回的内容
func Post(url string, data interface{}) (res Res) {
	jsonStr, _ := json.Marshal(data)
	// 查看post数据
	fmt.Println(string(jsonStr))

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Add("content-type", "application/json")
	if err != nil {
		res = Res {
			Code: 3002,
			Content: error.Error(err),
		}
	}
	defer req.Body.Close()

	client := &http.Client{Timeout: 5 * time.Second}
	resp, err2 := client.Do(req)
	if err2 != nil {
		res = Res {
			Code: 3003,
			Content: error.Error(err2),
		}
	}
	defer resp.Body.Close()
	code := resp.StatusCode
	result, _ := ioutil.ReadAll(resp.Body)
	content := string(result)
	res = Res {
		Code: code,
		Content: content,
	}
	return
}

//func Jenkins(){
//	jenkins := gojenkins.CreateJenkins(nil, "http://test-devops-jenkins.intramirror.com:8080/", "admin","Com123456")
//	_, err := jenkins.Init()
//	if err != nil {
//		fmt.Println(err)
//	}
//	//连接成功
//	fmt.Println("is ok")
//
//	m := make(map[string]string)
//	m["DEPLOY"] = "false"
//	m["CHECK_CODE_QUALITY"] = "true"
//	for i := 0; i <= 10; i++ {
//		res,err := jenkins.BuildJob("test-qa-opu-api",m)
//		fmt.Println(res,err)
//	}
//
//	job, err := jenkins.GetJob("test-qa-opu-api")
//	job.Poll()
//	if err != nil {
//		panic("Job Does Not Exist")
//	}
//}
