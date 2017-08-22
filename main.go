package main

import (
	_ "github.com/go-sql-driver/mysql"
	//"github.com/jinzhu/gorm"
	//"UmiMeServer/model"
	"fmt"
	"gopkg.in/gin-gonic/gin.v1"
	"net"
	"UmiMeServer/model"
	"github.com/jinzhu/gorm"
	"UmiMeServer/utils"
)

type User struct {
	Name     string `form:"name" json:"name" binding:"required"`
	Password string `form:"pwd" json:"pwd" binding:"required"`
}

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
	// 		From Go Server.
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

	//db, err := gorm.Open("mysql", "root:root@/UmiMe?charset=utf8&parseTime=True&loc=Local");
	//
	//defer db.Close()
	//
	//if (err == nil) {
	//	if (!db.HasTable(&model.EntryForm{})) {
	//		db.Set("gorm:table_options", "charset=utf8").CreateTable(model.EntryForm{})
	//	} else {
	//		fmt.Println("table exists")
	//	}
	//	form := model.EntryForm{Name: "张三aaa", Sex: "男", School: "他妈的傻逼大学", Grade: "他妈的小学一年级", ParentPhone: "110", Remark: ""}
	//	if (db.NewRecord(form)) {
	//		db.Create(&form);
	//	}
	//
	//	forms := []model.EntryForm{}
	//	db.Find(&forms)
	//
	//	for _, v := range forms {
	//		fmt.Println("Id : ", v.ID)
	//		fmt.Println("Username : ", v.Name)
	//		fmt.Println("Visibility : ", v.Sex)
	//		fmt.Println("------------------")
	//	}
	//
	//	fmt.Println("------------------")
	//
	//	form = model.EntryForm{}
	//	db.First(&form);
	//	fmt.Println(form.School)
	//
	//
	//} else {
	//	fmt.Println(err)
	//}

	db, dbErr := gorm.Open("mysql", "root:root@/UmiMe?charset=utf8&parseTime=True&loc=Local");
	defer db.Close()

	router := gin.Default()
	v1 := router.Group("/v1")

	v1.Use(CORSMiddleware())

	v1.POST("/test", func(c *gin.Context) {
		var json User
		err := c.BindJSON(&json)
		if err == nil {
			fmt.Println(json.Name, json.Password);
			ip, _, _ := net.SplitHostPort(c.Request.RemoteAddr)

			userIP := net.ParseIP(ip)
			c.JSON(200, gin.H{"data": json, "status": "asdf", "address": userIP})
		} else {
			fmt.Println(err)
		}

	})

	v1.POST("/enroll", func(c *gin.Context) {
		var form model.EntryForm
		err := c.BindJSON(&form)
		if err == nil {
			if &form != nil {
				if checkForm(form, c) {
					if dbErr == nil {
						db.Create(&form)
						c.JSON(200, gin.H{"data": "success", "code": 0})
						go sendEmailAfterEnroll()
						return
					}
				} else {
					return
				}
			}
		} else {
			c.JSON(400, gin.H{"error": err, "errorMsg": "未知错误", "code": -100});
			return
		}
		c.JSON(400, gin.H{"data": nil, "errorMsg": "未知错误", "code": -100})
	})

	router.Run(":8080");
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func sendEmailAfterEnroll() {
	user := "332114994@qq.com"
	password := "uyblohfooprocbed"
	host := "smtp.qq.com:587"
	to := "332114994@qq.com;j.xuyanjun@gmail.com"
	subject := "A Mail From Go Server."
	body := `您好：
	From Go Server.
	`
	fmt.Println("send email")
	err := utils.SendToMail(user, password, host, to, subject, body, "text")
	if err != nil {
		fmt.Println("Send mail error!")
		fmt.Println(err)
	} else {
		fmt.Println("Send mail success!")
	}
}

func checkForm(form model.EntryForm, c *gin.Context) bool {
	if len(form.Name) == 0 {
		c.JSON(200, gin.H{"data": nil, "errorMsg": "Name is required", "code": -1})
		return false
	} else if len(form.Grade) == 0 {
		c.JSON(200, gin.H{"data": nil, "errorMsg": "Grade is required", "code": -1})
		return false
	} else if len(form.ParentPhone) == 0 {
		c.JSON(200, gin.H{"data": nil, "errorMsg": "ParentPhone is required", "code": -1})
		return false
	}

	return true;
}
