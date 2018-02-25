package email

import (
	"gopkg.in/gomail.v2"
	"log"
	"strings"
	"strconv"
)

var useremail string
var password string
var host string
var port int

func SetConfig(smtphost string, username string, passwd string) {
	host = strings.Split(smtphost, ":")[0]
	port, _ = strconv.Atoi(strings.Split(smtphost, ":")[1])
	useremail = username
	password = passwd
}

func SendMail(to string, subject string, content string) {
	m := gomail.NewMessage()

	m.SetAddressHeader("From", useremail, "bzojChecker")
	m.SetHeader("To", m.FormatAddress(to, "用户"))
	m.SetHeader("Subject", subject)
	m.SetBody("text/plain", content)

	d := gomail.NewDialer(host, port, useremail, password)

	if err := d.DialAndSend(m); err != nil {
		log.Println("发送失败", err)
		return
	}

	log.Println("成功发送给 " + to)
}
