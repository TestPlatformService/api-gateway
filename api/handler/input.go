package handler

import (
	"api/genproto/question"
	"api/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetQuestionInputById godoc
// @Summary GetQuestionInputById
// @Description GetQuestionInputById
// @Tags questionInputAndOutput
// @Security ApiKeyAuth
// @Param id path string true "id"
// @Success 200 {object} model.GetQuestionInputWithOutputsResponse
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

	// Retrieve the question input by ID
	inputResp, err := h.QuestionInput.GetQuestionInput(c, &question.QuestionInputId{Id: id})
	if err != nil {
		h.Log.Error("Failed to get question input by id", "error", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Server error"})
		return
	}

	// Retrieve outputs associated with the input
	outputsResp, err := h.QuestionOutput.GetQUestionOutPutByInputId(c, &question.GetQUestionOutPutByInputIdRequest{InputId: id})
	if err != nil {
		h.Log.Error("Failed to get outputs for question input", "error", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Server error"})
		return
	}

	// Create a response structure that includes both input and outputs
	response := model.GetQuestionInputWithOutputsResponse{
		Input:  inputResp,
		Output: outputsResp.QuestionOutputs[0],
	}

	h.Log.Info("GetQuestionInputById ended successfully")
	c.JSON(http.StatusOK, response)
}

// GetQuestionInputsByQuestionId godoc
// @Summary GetQuestionInputsByQuestionId
// @Description GetQuestionInputsByQuestionId
// @Tags questionInputAndOutput
// @Security ApiKeyAuth
// @Param question_id path string true "question_id"
// @Success 200 {object} model.GetAllQuestionInputsWithOutputsByQuestionIdResponse
// @Failure 400 {object} string "Invalid request body"
// @Failure 500 {object} string "Server error"
// @Router /api/question-inputs/question/{question_id} [get]
func (h *Handler) GetQuestionInputsByQuestionId(c *gin.Context) {
	h.Log.Info("GetQuestionInputsByQuestionId is starting")
	questionId := c.Param("question_id")
	if questionId == "" {
		h.Log.Error("Invalid request body")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	// Get all inputs for the question
	inputsResp, err := h.QuestionInput.GetAllQuestionInputsByQuestionId(c, &question.GetAllQuestionInputsByQuestionIdRequest{QuestionId: questionId})
	if err != nil {
		h.Log.Error("Failed to get question inputs by question id", "error", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Server error"})
		return
	}

	// Prepare the response structure
	var inputsWithOutputs []model.GetQuestionInputWithOutput

	// Iterate through each input and get the corresponding output
	for _, input := range inputsResp.QuestionInputs {
		outputResp, err := h.QuestionOutput.GetQUestionOutPutByInputId(c, &question.GetQUestionOutPutByInputIdRequest{InputId: input.Id})
		if err != nil {
			h.Log.Error("Failed to get output for question input", "error", err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Server error"})
			return
		}

		// Create a new struct for input and output
		inputWithOutput := model.GetQuestionInputWithOutput{
			Input:  input,
			Output: nil, // Default to nil
		}

		// Assign the first output if available
		if len(outputResp.QuestionOutputs) > 0 {
			inputWithOutput.Output = outputResp.QuestionOutputs[0]
		}

		inputsWithOutputs = append(inputsWithOutputs, inputWithOutput)
	}

	// Create the final response
	finalResponse := model.GetAllQuestionInputsWithOutputsByQuestionIdResponse{
		InputsWithOutputs: inputsWithOutputs,
	}

	h.Log.Info("GetQuestionInputsByQuestionId ended successfully")
	c.JSON(http.StatusOK, finalResponse)
}

// DeleteQuestionInput godoc
// @Summary DeleteQuestionInput
// @Description DeleteQuestionInput
// @Tags questionInputAndOutput
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

	// Retrieve outputs associated with the input
	outputsResp, err := h.QuestionOutput.GetQUestionOutPutByInputId(c, &question.GetQUestionOutPutByInputIdRequest{InputId: id})
	if err != nil {
		h.Log.Error("Failed to get outputs for question input", "error", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Server error"})
		return
	}

	// Delete all associated outputs
	for _, output := range outputsResp.QuestionOutputs {
		_, err := h.QuestionOutput.DeleteQuestionOutput(c, &question.DeleteQuestionOutputRequest{Id: output.Id})
		if err != nil {
			h.Log.Error("Failed to delete question output", "error", err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Server error"})
			return
		}
	}

	// Now delete the question input
	_, err = h.QuestionInput.DeleteQuestionInput(c, &question.DeleteQuestionInputRequest{Id: id})
	if err != nil {
		h.Log.Error("Failed to delete question input", "error", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Server error"})
		return
	}

	h.Log.Info("DeleteQuestionInput ended successfully")
	c.JSON(http.StatusOK, gin.H{"message": "Question input deleted successfully"})
}
