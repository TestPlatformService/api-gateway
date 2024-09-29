package handler

import (
	"api/genproto/question"
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreateQuestionOutput godoc
// @Summary CreateQuestionOutput
// @Description CreateQuestionOutput
// @Tags questionOutput
// @Security ApiKeyAuth
// @Param info body question.CreateQuestionOutputRequest true "questionOutput info"
// @Success 200 {object} question.QuestionOutputId "id"
// @Failure 400 {object} string "Invalid request body"
// @Failure 500 {object} string "Server error"
// @Router /api/question-outputs/create [post]
func (h *Handler) CreateQuestionOutput(c *gin.Context) {
	h.Log.Info("CreateQuestionOutput is starting")
	req := question.CreateQuestionOutputRequest{}
	if err := c.BindJSON(&req); err != nil {
		h.Log.Error("Invalid request body", "error", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}
	resp, err := h.QuestionOutput.CreateQuestionOutput(c, &req)
	if err != nil {
		h.Log.Error("Failed to create question output", "error", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Server error"})
		return
	}
	h.Log.Info("CreateQuestionOutput ended successfully")
	c.JSON(http.StatusOK, resp)
}

// GetQuestionOutputById godoc
// @Summary GetQuestionOutputById
// @Description GetQuestionOutputById
// @Tags questionOutput
// @Security ApiKeyAuth
// @Param id path string true "id"
// @Success 200 {object} question.GetQuestionOutputResponse "questionOutput info"
// @Failure 400 {object} string "Invalid request body"
// @Failure 500 {object} string "Server error"
// @Router /api/question-outputs/{id} [get]
func (h *Handler) GetQuestionOutputById(c *gin.Context) {
	h.Log.Info("GetQuestionOutputById is starting")
	id := c.Param("id")
	if len(id) == 0 {
		h.Log.Error("Invalid request id")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request id"})
		return
	}
	resp, err := h.QuestionOutput.GetQuestionOutput(c, &question.QuestionOutputId{Id: id})
	if err != nil {
		h.Log.Error("Failed to get question output by id", "error", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Server error"})
		return
	}
	h.Log.Info("GetQuestionOutputById ended successfully")
	c.JSON(http.StatusOK, resp)
}

// DeleteQuestionOutput godoc
// @Summary DeleteQuestionOutput
// @Description DeleteQuestionOutput
// @Tags questionOutput
// @Security ApiKeyAuth
// @Param id path string true "id"
// @Success 200 {object} question.Void "Void"
// @Failure 400 {object} string "Invalid request body"
// @Failure 500 {object} string "Server error"
// @Router /api/question-outputs/delete/{id} [delete]
func (h *Handler) DeleteQuestionOutput(c *gin.Context) {
	h.Log.Info("DeleteQuestionOutput is starting")
	id := c.Param("id")
	if len(id) == 0 {
		h.Log.Error("Invalid request id")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request id"})
		return
	}
	_, err := h.QuestionOutput.DeleteQuestionOutput(c, &question.DeleteQuestionOutputRequest{Id: id})
	if err != nil {
		h.Log.Error("Failed to delete question output", "error", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Server error"})
		return
	}
	h.Log.Info("DeleteQuestionOutput ended successfully")
	c.JSON(http.StatusOK, gin.H{"message": "Question output deleted successfully"})
}

// GetQuestionOutputsByQuestionId godoc
// @Summary GetQuestionOutputsByQuestionId
// @Description GetQuestionOutputsByQuestionId
// @Tags questionOutput
// @Security ApiKeyAuth
// @Param question_id path string true "question_id"
// @Success 200 {object} question.GetAllQuestionOutputsByQuestionIdResponse "questionOutput info"
// @Failure 400 {object} string "Invalid request body"
// @Failure 500 {object} string "Server error"
// @Router /api/question-outputs/{question_id} [get]
func (h *Handler) GetQuestionOutputsByQuestionId(c *gin.Context) {
	h.Log.Info("GetQuestionOutputsByQuestionId is starting")
	questionId := c.Param("question_id")
	if len(questionId) == 0 {
		h.Log.Error("Invalid request question id")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request question id"})
		return
	}
	resp, err := h.QuestionOutput.GetAllQuestionOutputsByQuestionId(c, &question.GetAllQuestionOutputsByQuestionIdRequest{QuestionId: questionId})
	if err != nil {
		h.Log.Error("Failed to get question outputs by question id", "error", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Server error"})
		return
	}
	h.Log.Info("GetQuestionOutputsByQuestionId ended successfully")
	c.JSON(http.StatusOK, resp)
}
