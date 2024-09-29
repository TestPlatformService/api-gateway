package handler

import (
	"api/genproto/question"
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreateQuestionInput godoc
// @Summary CreateQuestionInput
// @Description CreateQuestionInput
// @Tags questionInput
// @Security ApiKeyAuth
// @Param questionInput body question.CreateQuestionInputRequest true "questionInput"
// @Success 200 {object} question.QuestionInputId
// @Failure 400 {object} string "Invalid request body"
// @Failure 500 {object} string "Server error"
// @Router /api/question-inputs/create [post]
func (h *Handler) CreateQuestionInput(c *gin.Context) {
	h.Log.Info("CreateQuestionInput is starting")
	var req question.CreateQuestionInputRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		h.Log.Error("Failed to bind JSON", "error", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}
	resp, err := h.QuestionInput.CreateQuestionInput(c, &req)
	if err != nil {
		h.Log.Error("Failed to create question input", "error", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Server error"})
		return
	}
	h.Log.Info("CreateQuestionInput ended successfully")
	c.JSON(http.StatusOK, resp)
}

// GetQuestionInputById godoc
// @Summary GetQuestionInputById
// @Description GetQuestionInputById
// @Tags questionInput
// @Security ApiKeyAuth
// @Param id path string true "id"
// @Success 200 {object} question.GetQuestionInputResponse
// @Failure 400 {object} string "Invalid request body"
// @Failure 500 {object} string "Server error"
// @Router /api/question-inputs/{id} [get]
func (h *Handler) GetQuestionInputById(c *gin.Context) {
	h.Log.Info("GetQuestionInputById is starting")
	id := c.Param("id")
	if id == "" {
		h.Log.Error("Invalid request body")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}
	resp, err := h.QuestionInput.GetQuestionInput(c, &question.QuestionInputId{Id: id})
	if err != nil {
		h.Log.Error("Failed to get question input by id", "error", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Server error"})
		return
	}
	h.Log.Info("GetQuestionInputById ended successfully")
	c.JSON(http.StatusOK, resp)
}

// GetQuestionInputsByQuestionId godoc
// @Summary GetQuestionInputsByQuestionId
// @Description GetQuestionInputsByQuestionId
// @Tags questionInput
// @Security ApiKeyAuth
// @Param question_id path string true "question_id"
// @Success 200 {object} question.GetAllQuestionInputsByQuestionIdResponse
// @Failure 400 {object} string "Invalid request body"
// @Failure 500 {object} string "Server error"
// @Router /api/question-inputs/{question_id} [get]
func (h *Handler) GetQuestionInputsByQuestionId(c *gin.Context) {
	h.Log.Info("GetQuestionInputsByQuestionId is starting")
	questionId := c.Param("question_id")
	if questionId == "" {
		h.Log.Error("Invalid request body")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}
	resp, err := h.QuestionInput.GetAllQuestionInputsByQuestionId(c, &question.GetAllQuestionInputsByQuestionIdRequest{QuestionId: questionId})
	if err != nil {
		h.Log.Error("Failed to get question inputs by question id", "error", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Server error"})
		return
	}
	h.Log.Info("GetQuestionInputsByQuestionId ended successfully")
	c.JSON(http.StatusOK, resp)
}

// DeleteQuestionInput godoc
// @Summary DeleteQuestionInput
// @Description DeleteQuestionInput
// @Tags questionInput
// @Security ApiKeyAuth
// @Param id path string true "id"
// @Success 200 {object} string "Question input deleted successfully"
// @Failure 400 {object} string "Invalid request body"
// @Failure 500 {object} string "Server error"
// @Router /api/question-inputs/delete/{id} [delete]
func (h *Handler) DeleteQuestionInput(c *gin.Context) {
	h.Log.Info("DeleteQuestionInput is starting")
	id := c.Param("id")
	if id == "" {
		h.Log.Error("Invalid request body")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}
	_, err := h.QuestionInput.DeleteQuestionInput(c, &question.DeleteQuestionInputRequest{Id: id})
	if err != nil {
		h.Log.Error("Failed to delete question input", "error", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Server error"})
		return
	}
	h.Log.Info("DeleteQuestionInput ended successfully")
	c.JSON(http.StatusOK, gin.H{"message": "Question input deleted successfully"})
}
