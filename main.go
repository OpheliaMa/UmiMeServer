package main

import "fmt"

func main() {
	fmt.Println("Hello UmiMe!");

	user := "332114994@qq.com"
	password := "uyblohfooprocbed"
	host := "smtp.qq.com:587"
	to := "j.xuyanjun@icloud.com;ophelia.ma@icloud.com"

	subject := "A Mail From Go Server."

	body := `
		您好：
			您收到一次报名提交，姓名为：[马玥]，联系方式为[15800705345]，请及时联系。

								      From Go Server.
		`;

	fmt.Println("send email")
	err := SendToMail(user, password, host, to, subject, body, "text")
	if err != nil {
		fmt.Println("Send mail error!")
		fmt.Println(err)
	} else {
		fmt.Println("Send mail success!")
	}
}