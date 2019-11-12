package controllers

import (
	"github.com/MobileCPX/PreBaseLib/splib/tracking"
	"github.com/MobileCPX/PreBaseLib/util"
	"github.com/astaxie/beego"
)

type DimocoATController struct {
	BaseController
}

func (c *DimocoATController) LPPage() {
	trackID := c.GetString("track")
	if trackID == "" {
		track := new(tracking.Track)
		track.ServiceID = "110543"
		track.ServiceName = "GameCenter"
		track.IP = util.GetIPAddress(c.Ctx.Request)
		track.UserAgent = c.Ctx.Request.UserAgent()
		track.Refer = c.Ctx.Request.Referer()
		track.ClickTime, _ = util.GetNowTime()
		track.AffName = "Test"
		trackID = track.RequestInsertTrack("http://dat.leadernet-hk.com/aff/click")

		c.Redirect("http://dat.leadernet-hk.com/identify?track="+trackID, 302)
		c.StopRun()
	}

	c.Data["URL"] = "http://dat.leadernet-hk.com/start-sub?track=" + trackID

	subID := c.Ctx.GetCookie("user")
	if subID != "" {
		c.Redirect("/", 302)
		c.StopRun()
	}

	c.TplName = "game/sp/dm_at/lp.html"
}

// 电话号码输入页面
func (c *DimocoATController) InputMsisdnPage() {
	trackID := c.GetString("track")
	c.Data["trackID"] = trackID
	c.TplName = "game/sp/dm_at/msisdn_input.html"
}

func (c *DimocoATController) WelcomePage() {
	//parm := c.Ctx.Request.URL.String()
	subStatus := c.GetString("sph-x")
	subID := c.GetString("sph-s")

	if subStatus == "s" {
		c.Ctx.SetCookie("user", subID)
		c.Data["text"] = "Die Transaktion war erfolgreich. Vielen Dank, dass Sie den Service nutzen. Sie werden in Kürze auf die" +
			" Seite weitergeleitet. Wenn dies nicht innerhalb weniger Minuten geschieht, klicken Sie auf Weiter."
		c.Data["URL"] = "/"
	} else {
		c.Data["text"] = "Abo fehlgeschlagen"
		c.Data["URL"] = "/at/sub"
	}
	c.TplName = "game/sp/dm_at/welcome.html"
}

func (c *DimocoATController) AGB() {
	c.setTpl("game/sp/dm_at/AGB.html")
}

func (c *DimocoATController) FAQ() {
	c.setTpl("game/sp/dm_at/FAQ.html")
}

func (c *DimocoATController) KONTAKT() {
	c.setTpl("game/sp/dm_at/KONTAKT.html")
}

func (c *DimocoATController) IMPRESSUM() {
	c.setTpl("game/sp/dm_at/IMPRESSUM.html")
}

func (c *DimocoATController) KUNDIGUNG() {
	c.setTpl("game/sp/dm_at/KUNDIGUNG.html")
}

func (c *DimocoATController) DATENSCHUTZERKLARUNG() {
	c.setTpl("game/sp/dm_at/DATENSCHUTZERKLARUNG.html")
}

func (c *DimocoATController) Rucktrittsrechts() {
	c.setTpl("game/sp/dm_at/Rucktrittsrechts.html")
}

func (c *DimocoATController) Widerrufsformular() {
	c.setTpl("game/sp/dm_at/Rucktrittsrechts.html")
}

const (
	//UnsubSuccessCode 退订成功
	UnsubSuccessCode = "0"

	//MsisdnIsEmptyCode 退订电话号码为空
	MsisdnIsEmptyCode = "-1"

	//MsisdnNotExistCode 退订电话号码不存在
	MsisdnNotExistCode = "-2"

	// XMLErrorCode xml解析错误
	XMLErrorCode = "-3"

	//UnsubFaieldCode  退订失败
	UnsubFaieldCode = "-4"
)

type UnsubResultControllers struct {
	beego.Controller
}

func (c *UnsubResultControllers) Get() {
	code := c.GetString("status")
	if code == UnsubSuccessCode {
		c.Data["text"] = "Ihr Abonnement wurde erfolgreich gekündigt"
		c.Data["URL"] = "/"
	} else {
		c.Data["text"] = "Die Kündigung des Abonnements ist fehlgeschlagen"
		c.Data["failed_desc"] = "Sie können Ihr Abonnement per E-Mail an info@mobileinfo.cc oder telefonisch unter 0800 100 336 (EUR 0,00/Min.) kündigen."
		c.Data["URL"] = "/"
	}
	c.TplName = "dm/at/unsub_result.html"
}
