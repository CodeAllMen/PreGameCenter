/**
  create by yy on 2020/3/19
*/

package ae

import "github.com/astaxie/beego"

type SubPage struct {
	beego.Controller
}

func (s *SubPage) Lp() {
	s.TplName = "ae/lp.html"
}

func (s *SubPage) Tan() {
	s.TplName = "ae/tan.html"
}

func (s *SubPage) Confirm() {
	s.TplName = "ae/confirm.html"
}

func (s *SubPage) Condition() {
	s.TplName = "ae/tnc.html"
}

func (s *SubPage) Help() {
	s.TplName = "ae/help.html"
}

func (s *SubPage) Privacy() {
	s.TplName = "ae/privacy.html"
}
