package controllers

import (
	"BCP/models"
	"fmt"
	"github.com/astaxie/beego"
)

type LoginController struct {
	beego.Controller
}

/**
* 直接跳转展示用户登录页面
 */

func (l *LoginController) Get() {
	l.TplName = "login.html"
}


/**
* post方法处理用户的登录请求
 */
func (l *LoginController) Post() {
	//1、解析客户端用户提交的登录数据
	var user models.User
	err := l.ParseForm(&user)
	if err != nil {
		l.Ctx.WriteString("抱歉，用户登录信息解析失败，请重试")
		return
	}
	//2、执行数据库查询操作
	u, err := user.QueryUser()
	if err != nil {
		fmt.Println(err.Error())
		l.Ctx.WriteString("抱歉，用户登录失败，请重试")
		return
	}
	l.Data["Phone"] = u.Phone  //动态数据设置

	l.TplName = "home.html"

	//3、判断数据库查询结果
	//4、根据查询结果返回客户端相应的信息或者页面跳转
}