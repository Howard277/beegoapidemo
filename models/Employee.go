package models

import "github.com/astaxie/beego/orm"

/*
	员工Model
 */
type Employee struct {
	Id int
	Name string
	Code string
}

func init(){
	// 注册当前Model
	orm.RegisterModel(new(Employee))
}
