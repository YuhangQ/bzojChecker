package main

import (
	"bzojChecker/checker"
	"bzojChecker/email"
	"bzojChecker/json"
	"bzojChecker/user"
	"bzojChecker/bzoj"
	"fmt"
	"io/ioutil"
	"strconv"
)

var config map[string]interface{}

func loadConfig() error {
	data, err := ioutil.ReadFile("config.json")
	if err != nil {
		return err
	}
	config, err = json.Prase(data)
	return err
}

var list []user.User

func onSubmit(user user.User, submit bzoj.Submit) {
	fmt.Println(submit)
	for _, u := range list {
		//if u == user {
		//	continue
		//}
		if u.Receive == false {
			continue
		}
		title := u.Username + " 提交了 " + strconv.Itoa(submit.Prob) + " 结果: " + submit.Result
		content := `提交人: %s
题目编号: %d
评测结果: %s
评测时间: %s`
		email.SendMail(u.Email, title, fmt.Sprintf(content, submit.User, submit.Prob, submit.Result, submit.Time))
	}
}

func main() {
	fmt.Println("BZOJ 爬虫程序 v1.1")
	fmt.Print("检查网络状态...")
	if bzoj.Alive() == false {
		fmt.Println("×\nERROR: 无法连接到 BZOJ.")
		return
	}
	fmt.Println("√")
	fmt.Print("检查配置文件...")
	err := loadConfig()
	if err != nil {
		fmt.Println("×\nERROR: 配置文件不存在或格式错误(config.json)")
		return
	}
	fmt.Println("√")
	fmt.Print("配置邮件服务...")
	email.SetConfig(config["smtphost"].(string), config["username"].(string), config["password"].(string))
	fmt.Println("√")
	fmt.Print("载入用户列表...")
	m := config["users"].([]interface{})
	for _, v := range m {
		t := v.(map[string]interface{})
		username := t["username"]
		email := t["email"]
		receive := t["receive"]
		list = append(list, user.NewUser(username.(string), email.(string), receive.(bool)))
	}
	fmt.Println("√")
	fmt.Print("处理用户信息...")
	err = checker.LoadRunID(list)
	if err != nil {
		fmt.Println("×\nERROR: 处理用户信息失败，请检查网络.")
		return
	}
	fmt.Println("√")
	fmt.Print("启动监听线程...")
	go checker.RegistEvent(onSubmit)
	fmt.Println("√")
	var cmd string
	for {
		fmt.Scanf("%s", &cmd)
		switch cmd {
		case "stop":
			fmt.Println("程序运行结束.")
			return
		case "help":
			fmt.Println("--指令帮助--")
			fmt.Println("stop -- 关闭此程序")
		default:
			fmt.Println("未知命令. 输入 help 获取帮助.")
		}
	}
}
