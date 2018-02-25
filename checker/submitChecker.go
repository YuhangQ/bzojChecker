package checker

import (
	"bzojChecker/user"
	"bzojChecker/bzoj"
	"bzojChecker/email"
	"time"
)

var runid map[string]int
var list []user.User

func LoadRunID(l []user.User) (error) {
	runid = make(map[string]int)
	list = l
	for _, u := range list {
		sub, err := bzoj.GetSubmit(u.Username)
		if err != nil {
			return err
		}
		runid[u.Username] = sub.ID
	}
	return nil
}

var flag bool

func RegistEvent(onSubmit func(user user.User, submit bzoj.Submit)) {
	flag = true
	for {
		time.Sleep(15 * time.Second)
		if bzoj.Alive() == false {
			if flag == true {
				flag = false
				for _, u := range list {
					email.SendMail(u.Email, "BZOJ 爆炸了！", "好了的话我会提醒你。")
				}
			}
			return
		} else {
			if flag == false {
				flag = true
				for _, u := range list {
					email.SendMail(u.Email, "BZOJ 好了！", "如果又炸了的话我会提醒你。")
				}
			}
		}
		for _, u := range list {
			sub, _ := bzoj.GetSubmit(u.Username)
			if sub.Result == "Compiling" || sub.Result == "Pending" || sub.Result == "Running_&_Judging" {
				continue
			}
			if runid[u.Username] == sub.ID {
				continue
			}
			runid[u.Username] = sub.ID
			onSubmit(u, sub)
		}
	}
}
