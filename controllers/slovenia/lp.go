/**
  create by yy on 2020/3/27
*/

package slovenia

import "github.com/astaxie/beego"

type SubPage struct {
	beego.Controller
}

func (s *SubPage) Lp() {
	s.TplName = "slovenia/lp.html"
}

func (s *SubPage) Tan() {
	s.TplName = "slovenia/tan.html"
}

func (s *SubPage) Confirm() {
	s.TplName = "slovenia/confirm.html"
}

func (s *SubPage) Condition() {
	s.TplName = "slovenia/tnc.html"
}

func (s *SubPage) Help() {
	s.TplName = "slovenia/help.html"
}

func (s *SubPage) Privacy() {
	s.TplName = "slovenia/privacy.html"
}
