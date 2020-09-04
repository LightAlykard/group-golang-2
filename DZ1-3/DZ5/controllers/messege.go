package controllers

import (
	"DZ5/models"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type MessegeController struct {
	beego.Controller
}

func (c *MessegeController) Get() {
	beeOrm := orm.NewOrm()

	messeges := []models.Messeges{}

	_, err := beeOrm.QueryTable("Messeges").All(&messeges)
	if err != nil {
		c.Ctx.Output.SetStatus(400)
		c.Ctx.Output.Body([]byte("Error getting Messeges from BD"))
	}
	c.Data["Title"] = "Test title"
	c.Data["Messeges"] = messeges
	c.TplName = "messeges.tpl"
}

func (c *MessegeController) GetOneMessege() {
	id := c.Ctx.Input.Param(":id")
	uid64, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		c.Ctx.Output.SetStatus(400)
		c.Ctx.Output.Body([]byte("Messege id is incorrect1"))
		return
	}
	beeOrm := orm.NewOrm()
	messege := models.Messeges{Id: uid64}
	err2 := beeOrm.QueryTable("Messeges").Filter("Id", uid64).One(&messege)
	if err2 != nil {
		c.Ctx.Output.SetStatus(400)
		c.Ctx.Output.Body([]byte("Messege id is incorrect2"))
		return
	}
	c.Data["Title"] = "Test title"
	c.Data["Messeges"] = messege
	c.TplName = "messege.tpl"
}

func (c *MessegeController) Post() {
	req := struct {
		Name string `json: "name"`
		Text string `json: "text"`
	}{}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &req); err != nil {
		c.Ctx.Output.SetStatus(400)
		c.Ctx.Output.Body([]byte("Body is empty"))
		return
	}

	messege, err := models.NewPost(req.Name, req.Text)
	if err != nil {
		c.Ctx.Output.SetStatus(400)
		c.Ctx.Output.Body([]byte(err.Error()))
		return
	}

	beeOrm := orm.NewOrm()

	id, err := beeOrm.Insert(messege)
	if err != nil {
		fmt.Println(err)
		c.Ctx.Output.SetStatus(400)
		c.Ctx.Output.Body([]byte("Error inserting messege in BD"))
		return
	}
	_ = id
	c.Data["json"] = messege
	c.ServeJSON()
}

func (c *MessegeController) UpdateMessege() {
	req := struct {
		Name string `json: "name"`
		Text string `json: "text"`
	}{}
	id := c.Ctx.Input.Param(":id")
	uid64, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		c.Ctx.Output.SetStatus(400)
		c.Ctx.Output.Body([]byte("Messege id is incorrect"))
		return
	}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &req); err != nil {
		c.Ctx.Output.SetStatus(400)
		c.Ctx.Output.Body([]byte("Body is empty"))
		return
	}

	messege, err := models.ExPost(req.Name, req.Text, uid64)
	if err != nil {
		c.Ctx.Output.SetStatus(400)
		c.Ctx.Output.Body([]byte(err.Error()))
		return
	}
	beeOrm := orm.NewOrm()

	pid, err := beeOrm.Update(messege)
	if err != nil {
		fmt.Println(err)
		c.Ctx.Output.SetStatus(400)
		c.Ctx.Output.Body([]byte("Error updating messege in BD"))
		return
	}
	_ = pid
	c.Data["json"] = messege
	c.ServeJSON()
}
func (c *MessegeController) DeleteMessege() {

	id := c.Ctx.Input.Param(":id")
	uid64, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		c.Ctx.Output.SetStatus(400)
		c.Ctx.Output.Body([]byte("Messege id is incorrect1"))
		return
	}

	messege, err := models.DelPost(uid64)
	if err != nil {
		c.Ctx.Output.SetStatus(400)
		c.Ctx.Output.Body([]byte(err.Error()))
		return
	}
	beeOrm := orm.NewOrm()
	fmt.Println(uid64)
	pid, err := beeOrm.Delete(messege)
	if err != nil {
		c.Ctx.Output.SetStatus(400)
		c.Ctx.Output.Body([]byte("Error deleting messege in BD"))
		return
	}
	_ = pid
	c.Data["json"] = messege
	c.ServeJSON()
}
