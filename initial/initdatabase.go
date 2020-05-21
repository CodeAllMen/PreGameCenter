package initial

import (
	"bufio"
	"fmt"
	"github.com/MobileCPX/PreGameCenter/models"
	"github.com/astaxie/beego/orm"
	"io"
	"os"
	"strings"
)

func FileToSql() {
	f, err := os.Open("game_text")
	defer f.Close()
	if err != nil {

	} else {
		rd := bufio.NewReader(f)
		for {
			line, err := rd.ReadString('\n') // 以'\n'为结束符读入一行
			if err != nil || io.EOF == err {
				break
			}
			line = strings.Replace(line, "\n", "", -1)
			data := strings.Split(line, "\t")
			o := orm.NewOrm()
			fmt.Println(data)
			var game models.Game
			game.GameName = data[0]
			game.GameID = data[1]
			o.Insert(&game)
		}
	}
}
