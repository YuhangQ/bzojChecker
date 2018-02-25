package email

import (
	"net/smtp"
	"strings"
)

func sendToMail(user, password, host, to, subject, body, mailtype string) error {
	hp := strings.Split(host, ":")
	auth := smtp.PlainAuth("", user, password, hp[0])
	var contentType string
	if mailtype == "html" {
		contentType = "Content-Type: text/" + mailtype + "; charset=UTF-8"
	} else {
		contentType = "Content-Type: text/plain" + "; charset=UTF-8"
	}

	msg := []byte("To: " + to + "\r\nFrom: " + user + ">\r\nSubject: " + subject + "\r\n" + contentType + "\r\n\r\n" + body)
	sendTos := strings.Split(to, ";")
	err := smtp.SendMail(host, auth, user, sendTos, msg)
	return err
}

var user string
var password string
var host string

func SetConfig(smtphost string, username string, passwd string) {
	host = smtphost
	user = username
	password = passwd
}

func SendMail(to string, subject string, contect string) error {
	err := sendToMail(user, password, host, to, subject, contect, "nohtml")
	return err
}
