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
// @Param	body		body 	models.question.AllQuestionResult		true		"QuestionBody"
// @Success 200 Add Success
// @Failure 403 Add fail
// @router /addQuestion [post]
// 将题目添加到数据库中
func (q *QuestionController) AddQuestion() {
	var question models.AllQuestionResult
	var resultData ResultData
	bytes := q.Ctx.Input.RequestBody
	println(string(bytes))
	json.Unmarshal(bytes, &question)
	isSuc := models.AddQuestionToDB(&question)
	if isSuc {
		resultData.Code = MSG_Suc
		resultData.Msg = "添加成功"
	} else {
		resultData.Code = MSG_ERR
		resultData.Msg = "添加失败"
	}
	q.Data["json"] = resultData
	q.ServeJSON()
}

// @Title GetAllQuestion
// @Description Get All Question From DB
// @Param	pageNum		query	int		true		"每一页的条数"
// @Param	pageSize	query	int		true		"请求的是第几页的数据"
// @Success 200 {object} models.question.AllQuestionResult
// @failure 403 找不到数据
// @router /getAllQuestion [get]
// 获取所有题目信息
func (q *QuestionController) GetAllQuestion() {
	pageNum, _ := q.GetInt("pageNum")
	pageSize, _ := q.GetInt("pageSize")
	question := models.FindAllQuestion(pageNum,pageSize)
	var resultData ResultData
	resultData.Code=MSG_Suc
	resultData.Msg="查询成功"
	resultData.Data=question
	q.Data["json"]=resultData
	q.ServeJSON()
}
