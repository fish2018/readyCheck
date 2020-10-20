package main

import (
	"fmt"
	"os"
)

type Job struct {
	Env string `json:"env"`
	App string `json:"app"`
}

//创建文件
func CreateFile() {
	f, err := os.Create("/dumps/up")
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()
	fmt.Println("创建文件")
}

// 判断所给路径文件是否存在
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

// code为200时，调用中控接口触发测试job
func callTest() {
	//创建文件
	CreateFile()

	job := Job{
		Env: os.Getenv("ENV"),
		App: os.Getenv("APP"),
	}
	fmt.Println("开始调用api")
	res := Post("http://demo.com/cmdb/v1/autotest", job)

	if res.Code != 200 {
		fmt.Println("************************* readyCheck 测试未通过 **********************************\n", res.Code, res.Content)
		//os.Exit(2)
	} else {
		fmt.Println("调用成功 ", res.Code, res.Content)
	}

}

func main() {
	code := Get("http://127.0.0.1:8888/actuator/info")
	if code == 200 {
		// 判断up文件是否存在
		exist, err := PathExists("/dumps/up")
		if err != nil {
			fmt.Printf("get dir error![%v]\n", err)
			return
		}
		if !exist {
			fmt.Println("不存在up文件")
			callTest()
		} else {
			fmt.Println("存在up文件")
		}

	} else {
		os.Exit(1)
	}
}
