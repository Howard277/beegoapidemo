package employeedao

import (
	"beegoapidemo/models"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
)

/*
	插入Employee
*/
func Insert(employee *models.Employee) int64 {
	o := orm.NewOrm()
	id, err := o.Insert(employee)
	if nil != err {
		logs.Error(err)
	}
	return id
}

/**

 */
func SelectAll() []models.Employee {
	var employees []models.Employee
	o := orm.NewOrm()
	o.Raw("select * from employee ").QueryRows(&employees)
	return employees
}
