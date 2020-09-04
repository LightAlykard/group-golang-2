package models

import (
	"fmt"
	"time"

	"github.com/astaxie/beego/orm"
)

type Messeges struct {
	Id   uint64
	Name string
	Text string
}

func (p *Messeges) TableName() string {
	return "messeges"
}

func NewPost(name, text string) (*Messeges, error) {
	if text == "" {
		return nil, fmt.Errorf("Empty Messege name")
	}

	return &Messeges{Text: text, Name: name}, nil
}
func ExPost(name, text string, id uint64) (*Messeges, error) {
	if text == "" {
		return nil, fmt.Errorf("Empty Messege name")
	}
	t := time.Now()
	return &Messeges{Id: id, Text: text, Name: name}, nil
}
func DelPost(id uint64) (*Messeges, error) {
	return &Messeges{Id: id}, nil
}

func init() {
	orm.RegisterModel(new(Messeges))
}
