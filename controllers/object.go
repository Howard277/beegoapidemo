package controllers

import (
	"beegoapidemo/models"
	"encoding/json"

	"github.com/astaxie/beego"
)

// object控制器
type ObjectController struct {
	beego.Controller
}

/*
	post提交一个对象
*/
func (o *ObjectController) Post() {
	var ob models.Object
	json.Unmarshal(o.Ctx.Input.RequestBody, &ob)
	objectid := models.AddOne(ob)
	o.Data["json"] = map[string]string{"ObjectId": objectid}
	o.ServeJSON()
}

/*
	get 获取一个对象
	objectId
*/
func (o *ObjectController) Get() {
	objectId := o.GetString("objectId")
	if objectId != "" {
		ob, err := models.GetOne(objectId)
		if err != nil {
			o.Data["json"] = err.Error()
		} else {
			o.Data["json"] = ob
		}
	}
	o.ServeJSON()
}

/*
	获取所有
*/
func (o *ObjectController) GetAll() {
	obs := models.GetAll()
	o.Data["json"] = obs
	o.ServeJSON()
}

/*
	删除
*/
func (o *ObjectController) Delete() {
	objectId := o.GetString("objectId")
	models.Delete(objectId)
	o.Data["json"] = "delete success!"
	o.ServeJSON()
}
