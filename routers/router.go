package routers

import (
	"github.com/MobileCPX/PreBaseLib/loginutil"
	"github.com/MobileCPX/PreGameCenter/controllers"
	"github.com/MobileCPX/PreGameCenter/controllers/ae"
	"github.com/MobileCPX/PreGameCenter/controllers/slovenia"
	"github.com/astaxie/beego"
)

func init() {
	// 游戏首页
	beego.Router("/", &controllers.HomeController{}, "Get:Home")

	// 游戏详情页
	beego.Router("/play/:gameID", &controllers.PlayPageController{}, "Get:PlayPage")

	beego.Router("/play/game/:gameID", &controllers.PlayPageController{}, "Get:PlayGame")

	// 条款页面
	beego.Router("/AGB.html", &controllers.DimocoATController{}, "Get:AGB")
	beego.Router("/FAQ.html", &controllers.DimocoATController{}, "Get:FAQ")
	beego.Router("/KONTAKT.html", &controllers.DimocoATController{}, "Get:KONTAKT")
	beego.Router("/IMPRESSUM.html", &controllers.DimocoATController{}, "Get:IMPRESSUM")
	beego.Router("/KUNDIGUNG.html", &controllers.DimocoATController{}, "Get:KUNDIGUNG")
	beego.Router("/DATENSCHUTZERKLARUNG.html", &controllers.DimocoATController{}, "Get:DATENSCHUTZERKLARUNG")
	beego.Router("/Rucktrittsrechts.html", &controllers.DimocoATController{}, "Get:Rucktrittsrechts")

	beego.Router("/Widerrufsformular.pdf", &controllers.DimocoATController{}, "Get:Widerrufsformular")

	beego.Router("/at/sub", &controllers.DimocoATController{}, "Get:LPPage")

	beego.Router("/wifi/msisdn", &controllers.DimocoATController{}, "Get:InputMsisdnPage")

	beego.Router("/add/user", &loginutil.UserLoginController{}, "Get:AddUser")

	beego.Router("/unsub/result", &controllers.UnsubResultControllers{})

	beego.Router("/dm/at/welcome", &controllers.DimocoATController{}, "Get:WelcomePage")

	beego.Router("/lpage", &ae.SubPage{}, "*:Lp")
	beego.Router("/tan", &ae.SubPage{}, "*:Tan")
	beego.Router("/confirm", &ae.SubPage{}, "*:Confirm")
	beego.Router("/tnc", &ae.SubPage{}, "*:Condition")
	beego.Router("/help", &ae.SubPage{}, "*:Help")
	beego.Router("/privacy", &ae.SubPage{}, "*:Privacy")

	beego.Router("/sv/lp", &slovenia.SubPage{}, "*:Lp")
	beego.Router("/sv/tan", &slovenia.SubPage{}, "*:Tan")
	beego.Router("/sv/confirm", &slovenia.SubPage{}, "*:Confirm")
	beego.Router("/sv/tnc", &slovenia.SubPage{}, "*:Condition")
	beego.Router("/sv/help", &slovenia.SubPage{}, "*:Help")
	beego.Router("/sv/privacy", &slovenia.SubPage{}, "*:Privacy")
}
