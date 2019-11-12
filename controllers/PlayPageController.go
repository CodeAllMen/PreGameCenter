package controllers

import "github.com/MobileCPX/PreGameCenter/models"

type PlayPageController struct {
	BaseController
}

func (c *PlayPageController) PlayPage() {
	gameID := c.Ctx.Input.Param(":gameID")

	game := new(models.Game)

	err := game.GetGameByID(gameID)
	if err != nil {
		c.redirect("/")
	}
	c.Data["game"] = game

	c.setTpl("game/page/play.html")
}

func (c *PlayPageController) PlayGame() {
	gameID := c.Ctx.Input.Param(":gameID")

	game := new(models.Game)

	err := game.GetGameByID(gameID)
	if err != nil {
		c.redirect("/")
	}
	c.redirect("http://static.gogamehub.com/game/" + game.GameID + "/index.html")
}
