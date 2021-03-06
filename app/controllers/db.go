package controllers

import (
  "App42PaaS-Revel-MySQL-Sample/app/models"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	r "github.com/revel/revel"
)

var DB gorm.DB

func init() {
	var err error
	DB, err = gorm.Open("mysql", "root@/revel_mysql_sample")
	if err != nil {
		panic(fmt.Sprintf("Got error when connect database, the error is '%v'", err))
	} else {
		fmt.Println("Successfully connected to DB: %#v", DB)
	}
}

type ModelController struct {
	*r.Controller
	Orm gorm.DB
}

func (c *ModelController) Begin() r.Result {
  c.Orm = DB
	return nil
}

func (c *ModelController) AddTables() r.Result {
  c.Orm.CreateTable(models.User{})
  return nil
}
