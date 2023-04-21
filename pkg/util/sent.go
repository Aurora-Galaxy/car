package util

import (
	"github.com/jordan-wright/email"
	"log"
	"net/smtp"
)

//发送通知
func sendMail(text []byte, mail string) {
	e := email.NewEmail()
	//设置发送方的邮箱
	e.From = "1302997173@qq.com"
	// 设置接收方的邮箱
	e.To = []string{mail}
	//设置主题
	e.Subject = "充电桩预约系统"
	//设置文件发送的内容
	e.Text = text
	//设置服务器相关的配置
	err := e.Send("smtp.qq.com:25", smtp.PlainAuth("",
		"1302997173@qq.com", "kmibjpfvgamxfggh", "smtp.qq.com"))
	if err != nil {
		log.Fatal(err)
	}
	//kmibjpfvgamxfggh
}
