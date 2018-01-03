package models

import "github.com/astaxie/beego/orm"

// 题目
// Category:题目类别
// Details:题目类别 0：安卓，1：java
// QuestionType:题目类别（0:选择题，1:文字描述题）
// Question：题目
// Options:选择题的选项(如果是文字描述题，则改项为空)
// Answer:题目的答案，如果是选择题，则该项为Check的string，如果为文字描述体，该项则为具体的文字表述
type AllQuestion struct {
	Id           int       `pk:"auto"`
	Category     int       `orm:"column(category)"`
	QuestionType int
	Question     string
	Options      []*Option `orm:"reverse(many)"`
	Answer       string
}

//自定义表名
func (a *AllQuestion) TableName() string {
	return "service_question"
}

//选项
//Check:选择题选项是哪个（如A，B，C，D）
//Description:每个选项对应的描述
type Option struct {
	Id          int          `pk:"auto"`
	Check       string
	Description string
	AllQuestion *AllQuestion `orm:"rel(fk)"`
}

//往数据库中增加题目
func AddQuestionToDB(q *AllQuestion) bool {
	_, e := orm.NewOrm().Insert(q)
	if e != nil {
		return false
	}
	if len(q.Options) > 0 {
		for _, v := range q.Options { //range返回的是数据结构的索引和元素
			v.AllQuestion = q
			_, e2 := orm.NewOrm().Insert(v)
			if e2 != nil {
				orm.NewOrm().Delete(q)
				return false
			}
		}
	}
	return true
}
