package controllers

import "github.com/MobileCPX/PreGameCenter/models"

type HomeController struct {
	BaseController
}

func (c *HomeController) Home() {

	pageNum, _ := c.GetInt("p", 0)

	game := new(models.Game)
	gameList, totalPage := game.GetGamePage(pageNum, 10)

	c.Data["GameList"] = gameList
	c.Data["currentPage"] = pageNum
	c.Data["pageNumList"] = c.GetPageNumList(totalPage)
	c.setTpl("game/page/home.html")
	//c.TplName = "game/page/home.html"
}
