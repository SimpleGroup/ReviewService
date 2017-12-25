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
	json.Unmarshal(q.Ctx.Input.RequestBody, &question)
}
