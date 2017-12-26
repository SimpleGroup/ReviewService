package controllers

import (
	"encoding/json"
	"ReviewService/models"
)

type QuestionController struct {
	BaseController
}

// @Title AddQuestion
// @Description Add question to DB
// @Param	body		body 	models.question.AllQuestion		true		"QuestionBody"
// @Success 200 Add Success
// @Failure 403 Add fail
// @router /addQuestion [post]
// 将题目添加到数据库中
func (q *QuestionController) AddQuestion() {
	var question models.AllQuestion
	var resultData ResultData
	bytes := q.Ctx.Input.RequestBody
	println(string(bytes))
	json.Unmarshal(bytes, &question)
	isSuc := models.AddQuestionToDB(&question)
	if isSuc {
		resultData.Code=MSG_Suc
		resultData.Msg="添加成功"
	}else {
		resultData.Code=MSG_ERR
		resultData.Msg="添加失败"
	}
	q.Data["json"]=resultData
	q.ServeJSON()
}
