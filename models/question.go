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
	Id           int `pk:"auto"`
	Category     int
	QuestionType int
	Question     string
	Answer       string
}

//对外暴露的结构
type AllQuestionResult struct {
	Id           int
	Category     int
	QuestionType int
	Question     string
	Options      []*Option
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
	Id            int   `pk:"auto"`
	Check         string
	Description   string
	AllQuestionId int64 `orm:"column(all_question_id)"`
}

//自定义表名
func (a *Option) TableName() string {
	return "service_option"
}

//往数据库中增加题目
func AddQuestionToDB(a *AllQuestionResult) bool {
	question := AllQuestion{Id:0,Category: a.Category, QuestionType: a.QuestionType, Question: a.Question, Answer: a.Answer}
	id, e := orm.NewOrm().Insert(&question)
	if e != nil {
		println(e.Error())
		return false
	}
	i := len(a.Options)
	if i > 0 {
		for k := range a.Options {
			a.Options[k].AllQuestionId = id
		}
		orm.NewOrm().InsertMulti(i, a.Options)
	}
	return true
}

//查看全部题库信息
func FindAllQuestion(pageNum int, pageSize int) ([]*AllQuestionResult) {
	var allQuestionResults []*AllQuestionResult
	var allQuestions []*AllQuestion
	_, e := orm.NewOrm().QueryTable("service_question").Limit(pageNum,pageNum*pageSize).All(&allQuestions)
	if e != nil {
		println(e.Error())
	}
	for k,v:= range allQuestions {
		var options []*Option
		_, e2 := orm.NewOrm().QueryTable("service_option").Filter("all_question_id", allQuestions[k].Id).All(&options)
		if e2 !=nil {
			options=nil
		}
		var allQuestionResult=AllQuestionResult{Category:v.Category,QuestionType:v.QuestionType,Question:v.Question,Options:options,Answer:v.Answer}
		allQuestionResults=append(allQuestionResults, &allQuestionResult)
	}
	return allQuestionResults
}
