package controllers

import (
	"BCP/models"
	"fmt"
	"github.com/astaxie/beego"
)

type RegisterController struct {
	beego.Controller
}

/**
*该方法处理用户注册逻辑
 */

func (r *RegisterController) Post() {
	var user models.User
	err := r.ParseForm(&user)
	if err != nil {
		r.Ctx.WriteString("抱歉，数据解析失败，请重试！")
		return
	}
	_, err = user.AddUser()
	if err != nil {
		fmt.Println(err.Error())
		r.Ctx.WriteString("抱歉，用户注册失败，请重试")
		return
	}

	//r.Ctx.WriteString("恭喜，注册用户信息成功！")

	r.TplName =  "login.html"

}
