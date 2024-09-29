package handler

import (
	"api/genproto/question"
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreateTestCase godoc
// @Summary CreateTestCase
// @Description Create a new test case
// @Tags testCase
// @Security ApiKeyAuth
// @Param request body question.CreateTestCaseRequest true "Create test case request"
// @Success 200 {object} question.TestCaseId "Test case created successfully"
// @Failure 400 {object} string "Invalid request body"
// @Failure 500 {object} string "Server error"
// @Router /api/test-cases/create [post]
func (h *Handler) CreateTestCase(c *gin.Context) {
	h.Log.Info("CreateTestCase is starting")
	var request question.CreateTestCaseRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		h.Log.Error("Failed to bind JSON", "error", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}
	resp, err := h.TestCase.CreateTestCase(c, &request)
	if err != nil {
		h.Log.Error("Failed to create test case", "error", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Server error"})
		return
	}
	h.Log.Info("CreateTestCase ended successfully")
	c.JSON(http.StatusOK, resp)
}

// GetTestCaseById godoc
// @Summary GetTestCaseById
// @Description Get test case by id
// @Tags testCase
// @Security ApiKeyAuth
// @Param id path string true "Test case id"
// @Success 200 {object} question.GetTestCaseResponse "Test case retrieved successfully"
// @Failure 400 {object} string "Invalid request body"
// @Failure 500 {object} string "Server error"
// @Router /api/test-cases/{id} [get]
func (h *Handler) GetTestCaseById(c *gin.Context) {
	h.Log.Info("GetTestCaseById is starting")
	id := c.Param("id")
	if id == "" {
		h.Log.Error("Invalid request body")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}
	resp, err := h.TestCase.GetTestCase(c, &question.TestCaseId{Id: id})
	if err != nil {
		h.Log.Error("Failed to get test case", "error", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Server error"})
		return
	}
	h.Log.Info("GetTestCaseById ended successfully")
	c.JSON(http.StatusOK, resp)
}

// GetTestCasesByQuestionId godoc
// @Summary GetTestCasesByQuestionId
// @Description Get test cases by question id
// @Tags testCase
// @Security ApiKeyAuth
// @Param question_id path string true "Question id"
// @Success 200 {object} question.GetAllTestCasesByQuestionIdResponse "Test cases retrieved successfully"
// @Failure 400 {object} string "Invalid request body"
// @Failure 500 {object} string "Server error"
// @Router /api/test-cases/{question_id} [get]
func (h *Handler) GetTestCasesByQuestionId(c *gin.Context) {
	h.Log.Info("GetTestCasesByQuestionId is starting")
	questionId := c.Param("question_id")
	if questionId == "" {
		h.Log.Error("Invalid request body")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}
	resp, err := h.TestCase.GetAllTestCasesByQuestionId(c, &question.GetAllTestCasesByQuestionIdRequest{QuestionId: questionId})
	if err != nil {
		h.Log.Error("Failed to get test cases", "error", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Server error"})
		return
	}
	h.Log.Info("GetTestCasesByQuestionId ended successfully")
	c.JSON(http.StatusOK, resp)
}

// DeleteTestCase godoc
// @Summary DeleteTestCase
// @Description Delete test case
// @Tags testCase
// @Security ApiKeyAuth
// @Param id path string true "Test case id"
// @Success 200 {object} string "Test case deleted successfully"
// @Failure 400 {object} string "Invalid request body"
// @Failure 500 {object} string "Server error"
// @Router /api/test-cases/delete/{id} [delete]
func (h *Handler) DeleteTestCase(c *gin.Context) {
	h.Log.Info("DeleteTestCase is starting")
	id := c.Param("id")
	if id == "" {
		h.Log.Error("Invalid request body")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}
	_, err := h.TestCase.DeleteTestCase(c, &question.DeleteTestCaseRequest{Id: id})
	if err != nil {
		h.Log.Error("Failed to delete test case", "error", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Server error"})
		return
	}
	h.Log.Info("DeleteTestCase ended successfully")
	c.JSON(http.StatusOK, gin.H{"message": "Test case deleted successfully"})
}
