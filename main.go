package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"UmiMeServer/model"
	"fmt"
)

func main() {
	//fmt.Println("Hello UmiMe!");
	//
	//user := "332114994@qq.com"
	//password := "uyblohfooprocbed"
	//host := "smtp.qq.com:587"
	//to := "j.xuyanjun@icloud.com;ophelia.ma@icloud.com"
	//
	//subject := "A Mail From Go Server."
	//
	//body := `
	//	您好：
	//		您收到一次报名提交，姓名为：[马玥]，联系方式为[15800705345]，请及时联系。
	//
	//							      From Go Server.
	//	`;
	//
	//fmt.Println("send email")
	//err := SendToMail(user, password, host, to, subject, body, "text")
	//if err != nil {
	//	fmt.Println("Send mail error!")
	//	fmt.Println(err)
	//} else {
	//	fmt.Println("Send mail success!")
	//}


	db, err := gorm.Open("mysql", "root:root@/UmiMe?charset=utf8&parseTime=True&loc=Local");

	defer db.Close()

	if (err == nil) {
		if (!db.HasTable(&model.EntryForm{})) {
			db.Set("gorm:table_options", "charset=utf8").CreateTable(model.EntryForm{})
		} else {
			fmt.Println("table exists")
		}
		form := model.EntryForm{Name: "张三", Sex: "男", School: "他妈的傻逼大学", Grade: "他妈的小学一年级", ParentPhone: "110", Remark: ""}
		if (db.NewRecord(form)) {
			db.Create(&form);
		}

		forms := []model.EntryForm{}
		db.Find(&forms)

		for _, v := range forms {
			fmt.Println("Id : ", v.ID)
			fmt.Println("Username : ", v.Name)
			fmt.Println("Visibility : ", v.Sex)
			fmt.Println("------------------")
		}

		fmt.Println("------------------")

		form = model.EntryForm{}
		db.First(&form);
		fmt.Println(form.School)


	} else {
		fmt.Println(err)
	}

}