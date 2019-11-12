package controllers

import "github.com/astaxie/beego"

type BaseController struct {
	beego.Controller
}

func (c *BaseController) redirect(url string) {
	c.Redirect(url, 302)
	c.StopRun()
}

func (c *BaseController) CheckLogStatus() {

}

// 获取页数列表
func (c *BaseController) GetPageNumList(maxNum int) []int {
	var totalPageList []int

	for i := 1; i <= maxNum; i++ {
		totalPageList = append(totalPageList, i)
	}
	return totalPageList

}

func (c *BaseController)setTpl(tpl string){
	c.Layout = "game/layout/layout.html"
	c.TplName = tpl
}
