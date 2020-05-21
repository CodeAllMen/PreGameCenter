package models

import "github.com/astaxie/beego/orm"

type Game struct {
	GameID     string `orm:"pk;column(game_id)"`
	GameName   string `orm:"game_name"`
	CategoryId int    `orm:"category_id" json:"category_id"`
}

//  获取游戏列表，传页数及每页限制数量
func (game *Game) GetGamePage(pageNum, limitNum int) (*[]Game, int) {
	o := orm.NewOrm()
	gameList := new([]Game)
	_, _ = o.QueryTable("game").Limit(limitNum, pageNum*limitNum).All(gameList)
	return gameList, game.getTotalPage(limitNum)
}

// 获取游戏总页数
func (game *Game) getTotalPage(limitNum int) int {
	o := orm.NewOrm()
	totalNum, _ := o.QueryTable("game").Count()
	return int(totalNum) / limitNum
}

// 根据游戏ID查询游戏
func (game *Game) GetGameByID(gameID string) error {
	o := orm.NewOrm()
	return o.QueryTable("game").Filter("game_id", gameID).One(game)

}
