/**
  create by yy on 2020/4/13
*/

package models

type CategoryModel struct {
	Id   int    `orm:"pk;column(id)" json:"id"`
	Name string `orm:"name" json:"name"`
}

