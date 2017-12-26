package models

import "github.com/astaxie/beego/orm"

// 题目
// Category:题目类别
// Details:题目类别 0：安卓，1：java
type AllQuestion struct {
	Id       int              `orm:"pk"`
	Category int              `orm:"column(category)"`
	QuestionDetails  *QuestionDetails `orm:"rel(fk)"`
}

//自定义表名
func (q *AllQuestion) TableName() string {
	return "service_question"
}

//题目详情
// QuestionType:题目类别（0:选择题，1:文字描述题）
// Question：题目
// Option:选择题的选项	选项形如 "A:...;B:...;C...;D...;"选项之间用分号隔开
// Answer:选择题的答案
type QuestionDetails struct {
	Id           int          `orm:"pk"`
	QuestionType int          `orm:"column(question_type)"`
	Question     string       `orm:"column(question)"`
	Option       string       `orm:"column(option)"`
	Answer       string       `orm:"column(answer)"`
}

////自定义表名
//func (q *QuestionDetails) TableName() string {
//	return "service_question_details"
//}

func AddQuestionToDB(q *AllQuestion) bool {
	_, e := orm.NewOrm().Insert(q)
	if e != nil {
		return false
	}
	return true
}
