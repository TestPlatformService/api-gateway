package handler

import (
	"api/genproto/question"
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreateQuestion godoc
// @Summary CreateQuestion
// @Description CreateQuestion
// @Tags question
// @Security ApiKeyAuth
// @Param info body question.QuestionId true "question info"
// @Success 200 {object} question.QuestionId "id"
// @Failure 400 {object} string "Invalid request body"
// @Failure 500 {object} string "Server error"
// @Router /api/questions/create [post]
func (h *Handler) CreateQuestion(c *gin.Context) {
	h.Log.Info("CreateQuestion is starting")
	req := question.CreateQuestionRequest{}
	if err := c.BindJSON(&req); err != nil {
		h.Log.Error("Invalid request body", "error", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	res, err := h.Question.CreateQuestion(c, &req)
	if err != nil {
		h.Log.Error("Failed to create question", "error", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Server error"})
		return
	}

	h.Log.Info("CreateQuestion ended successfully")
	c.JSON(http.StatusOK, res)

}

// GetQuestionById godoc
// @Summary GetQuestionById
// @Description GetQuestionById
// @Tags question
// @Security ApiKeyAuth
// @Param id path string true "id"
// @Success 200 {object} question.GetQuestionResponse "id"
// @Failure 400 {object} string "Invalid request body"
// @Failure 500 {object} string "Server error"
// @Router /api/questions/{id} [get]
func (h *Handler) GetQuestionById(c *gin.Context) {
	h.Log.Info("GetQuestionById is starting")
	req := question.QuestionId{}
	req.Id = c.Param("id")
	if len(req.Id) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Product ID is required"})
		h.Log.Error("Product ID is required")
		return
	}

	res, err := h.Question.GetQuestion(c, &req)
	if err != nil {
		h.Log.Error("Failed to create question", "error", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Server error"})
		return
	}

	h.Log.Info("GetQuestionById ended successfully")
	c.JSON(http.StatusOK, res)
}
