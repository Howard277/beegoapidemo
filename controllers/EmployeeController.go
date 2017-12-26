package controllers

import (
	"beegoapidemo/daos"
	"beegoapidemo/models"
	"github.com/astaxie/beego"
)

/*
	Employee控制器
*/
type EmployeeController struct {
	beego.Controller
}

func (ec *EmployeeController) Insert() {
	employee := new(models.Employee)
	employee.Name = "小孩"
	employee.Code = "001"
	id := employeedao.Insert(employee)
	ec.Data["json"] = id
	ec.ServeJSON()
}

func (ec *EmployeeController) SelectAll() {
	employees:=employeedao.SelectAll()
	ec.Data["json"] = employees
	ec.ServeJSON()
}
