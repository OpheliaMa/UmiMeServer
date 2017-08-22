package utils

import (
	"strings"
	"net/smtp"
	"fmt"
	"encoding/base64"
)

func SendToMail(user, password, host, to, subject, body, mailtype string) error {

	hp := strings.Split(host, ":");
	auth := smtp.PlainAuth("", user, password, hp[0]);

	header := make(map[string]string)
	header["From"] = user
	header["To"] = to
	header["Subject"] = subject
	header["MIME-Version"] = "1.0"
	header["Content-Type"] = "text/plain; charset=\"utf-8\""
	header["Content-Transfer-Encoding"] = "base64"
	header["X-Priority"] = "3"
	header["X-MSMail-Priority"] = "Normal"
	header["X-Mailer"] = "Microsoft Outlook Express 6.00.2900.2869"; //本文以outlook名义发送邮件，不会被当作垃圾邮件
	header["X-MimeOLE"] = "Produced By Microsoft MimeOLE V6.00.2900.2869";
	header["ReturnReceipt"] = "1";

	message := ""
	for k, v := range header {
		message += fmt.Sprintf("%s: %s\r\n", k, v)
	}
	message += "\r\n" + base64.StdEncoding.EncodeToString([]byte(body))

	send_to := strings.Split(to, ";")

	err := smtp.SendMail(host, auth, user, send_to, []byte(message))
	return err
}
