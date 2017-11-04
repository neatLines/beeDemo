package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql" // import your used driver
)

type User struct {
	Id       int
	Username string `orm:"size(255)"`
	Password string `orm:"size(255)"`
}

func init() {
	// set default database
	orm.RegisterDataBase("default", "mysql", "panis:oo7omysql@tcp(115.159.47.170:3306)/demo?charset=utf8", 30)

	// register model
	orm.RegisterModel(new(User))

	// create table
	orm.RunSyncdb("default", false, true)
}

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}

type UserController struct {
	beego.Controller
}

func (c *UserController) Get() {
	o := orm.NewOrm()

	user := User{Username: "slene", Password: "123"}

	// insert
	id, err := o.Insert(&user)
	fmt.Printf("ID: %d, ERR: %v\n", id, err)

}
