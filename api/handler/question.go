package handler

import (
	"api/config"
	"api/genproto/question"
	"api/genproto/topic"
	"api/model"
	"context"
	"fmt"
	"net/http"
	"path/filepath"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

// CreateQuestion godoc
// @Summary CreateQuestion
// @Description CreateQuestion
// @Tags question
// @Security ApiKeyAuth
// @Param info body question.CreateQuestionRequest true "question info"
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
		c.JSON(http.StatusBadRequest, gin.H{"error": "questions ID is required"})
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

// CreateQuestion godoc
// @Summary CreateQuestion
// @Description CreateQuestion
// @Tags question
// @Security ApiKeyAuth
// @Param limit query string false "limit"
// @Param offset query string false "offset"
// @Param topic_name query string false "topic_name"
// @Param type query string false "type"
// @Param name query string false "name"
// @Param number query string false "number"
// @Param difficulty query string false "difficulty"
// @Param input_info query string false "input_info"
// @Param output_info query string false "output_info"
// @Success 200 {object} question.QuestionId "id"
// @Failure 400 {object} string "Invalid request body"
// @Failure 500 {object} string "Server error"
// @Router /api/questions/getAll [get]
func (h *Handler) GetAllQuestions(c *gin.Context) {
	h.Log.Info("GetQuestions is starting")
	req2 := model.GetAllQuestionsRequest{}
	var limitstr, offsetstr int64
	if err := c.ShouldBindQuery(&req2); err != nil {
		h.Log.Error("Invalid query parameters", "error", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid query parameters"})
		return
	}
	limitStr := c.Query("limit")
	offsetStr := c.Query("offset")
	if limitStr != "" {
		limit, err := strconv.Atoi(limitStr)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest,
				gin.H{"error": err.Error()})
			h.Log.Error(err.Error())
			return
		}
		limitstr = int64(limit)
	} else {
		limitstr = 10
	}

	if offsetStr != "" {
		offset, err := strconv.Atoi(offsetStr)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest,
				gin.H{"error": err.Error()})
			h.Log.Error(err.Error())
			return
		}
		offsetstr = int64(offset)
	} else {
		offsetstr = 1
	}
	topicid := ""
	if req2.TopicName != "" {
		topicId, err := h.Topic.GetTopicIdByName(c, &topic.TopicNameReq{Name: req2.Name})
		if err != nil {
			h.Log.Error("Failed to get topic", "error", err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Server error"})
			return
		}
		topicid = topicId.Id
	}

	res, err := h.Question.GetAllQuestions(c, &question.GetAllQuestionsRequest{
		Limit:      limitstr,
		Offset:     offsetstr,
		TopicId:    topicid,
		Type:       req2.Type,
		Name:       req2.Name,
		Number:     req2.Number,
		Difficulty: req2.Difficulty,
		Language:   req2.Language,
	})
	if err != nil {
		h.Log.Error("Failed to get questions", "error", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Server error"})
		return
	}
	h.Log.Info("GetQuestions ended successfully")
	c.JSON(http.StatusOK, res)
}

// UpdateQuestion godoc
// @Summary UpdateQuestion
// @Description UpdateQuestion
// @Tags question
// @Security ApiKeyAuth
// @Param id path string true "id"
// @Param info body model.UpdateQuestionRequest true "question info"
// @Success 200 {object} string "Question updated successfully"
// @Failure 400 {object} string "Invalid request body"
// @Failure 500 {object} string "Server error"
// @Router /api/questions/update/{id} [put]
func (h *Handler) UpdateQuestion(c *gin.Context) {
	h.Log.Info("UpdateQuestion is starting")
	req := model.UpdateQuestionRequest{}
	if err := c.BindJSON(&req); err != nil {
		h.Log.Error("Invalid request body", "error", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}
	Id := c.Param("id")
	if len(Id) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "questions ID is required"})
		h.Log.Error("questions ID is required")
		return
	}

	req2 := question.UpdateQuestionRequest{
		Id:          Id,
		TopicId:     req.TopicId,
		Type:        req.Type,
		Name:        req.Name,
		Number:      req.Number,
		Difficulty:  req.Difficulty,
		InputInfo:   req.InputInfo,
		OutputInfo:  req.OutputInfo,
		Language:    req.Language,
		TimeLimit:   req.TimeLimit,
		MemoryLimit: req.MemoryLimit,
	}

	_, err := h.Question.UpdateQuestion(c, &req2)
	if err != nil {
		h.Log.Error("Failed to update question", "error", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Server error"})
		return
	}
	h.Log.Info("UpdateQuestion ended successfully")
	c.JSON(http.StatusOK, gin.H{"message": "Question updated successfully"})
}

// DeleteQuestion godoc
// @Summary DeleteQuestion
// @Description DeleteQuestion
// @Tags question
// @Security ApiKeyAuth
// @Param id path string true "id"
// @Success 200 {object} string "Question deleted successfully"
// @Failure 400 {object} string "Invalid request body"
// @Failure 500 {object} string "Server error"
// @Router /api/questions/delete/{id} [delete]
func (h *Handler) DeleteQuestion(c *gin.Context) {
	h.Log.Info("DeleteQuestion is starting")
	req := question.QuestionId{}
	req.Id = c.Param("id")
	if len(req.Id) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "questions ID is required"})
		h.Log.Error("questions ID is required")
		return
	}
	_, err := h.Question.DeleteQuestion(c, &question.DeleteQuestionRequest{Id: req.Id})
	if err != nil {
		h.Log.Error("Failed to delete question", "error", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Server error"})
		return
	}
	h.Log.Info("DeleteQuestion ended successfully")
	c.JSON(http.StatusOK, gin.H{"message": "Question deleted successfully"})
}

// UploadImageToQuestion godoc
// @Summary UploadImageToQuestion
// @Description UploadImageToQuestion
// @Tags question
// @Security ApiKeyAuth
// @Param file formData file true "file"
// @Param id path string true "id"
// @Success 200 {object} string "Image uploaded successfully"
// @Failure 400 {object} string "Invalid request body"
// @Failure 500 {object} string "Server error"
// @Router /api/questions/upload-image/{id} [post]
func (h *Handler) UploadImageToQuestion(c *gin.Context) {
	h.Log.Info("UploadImageToQuestion called")
	Id := c.Param("id")
	if len(Id) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "questions ID is required"})
		h.Log.Error("questions ID is required")
		return
	}

	file, header, err := c.Request.FormFile("file")
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Error retrieving the file"})
		return
	}
	defer file.Close()

	// minio start

	fileExt := filepath.Ext(header.Filename)
	println("\n File Ext:", fileExt)

	newFile := uuid.NewString() + fileExt
	minioClient, err := minio.New(config.Load().MINIO_URL, &minio.Options{
		Creds:  credentials.NewStaticV4("test", "minioadmin", ""),
		Secure: false,
	})
	if err != nil {
		c.AbortWithError(500, err)
		return
	}

	info, err := minioClient.PutObject(context.Background(), "questions", newFile, file, header.Size, minio.PutObjectOptions{
		ContentType: "image/jpeg",
	})
	if err != nil {
		c.AbortWithError(500, err)
		fmt.Println(err.Error())
		return
	}

	policy := fmt.Sprintf(`{
	 "Version": "2012-10-17",
	 "Statement": [
	  {
	   "Effect": "Allow",
	   "Principal": {
		"AWS": ["*"]
	   },
	   "Action": ["s3:GetObject"],
	   "Resource": ["arn:aws:s3:::%s/*"]
	  }
	 ]
	}`, "questions")

	err = minioClient.SetBucketPolicy(context.Background(), "questions", policy)
	if err != nil {
		c.AbortWithError(500, err)
		return
	}

	madeUrl := fmt.Sprintf("http://%s/questions/%s", config.Load().MINIO_URL, newFile)

	println("\n Info Bucket:", info.Bucket)

	// minio end

	_, err = h.Question.UploadImageQuestion(c, &question.UploadImageQuestionRequest{QuestionId: Id, Image: madeUrl})
	if err != nil {
		h.Log.Error("Failed to upload image to question", "error", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Server error"})
		return
	}
	h.Log.Info("UploadImageToQuestion ended successfully")
	c.JSON(http.StatusOK, gin.H{"Url": madeUrl})
}

// DeleteImageFromQuestion godoc
// @Summary DeleteImageFromQuestion
// @Description DeleteImageFromQuestion
// @Tags question
// @Security ApiKeyAuth
// @Param id path string true "id"
// @Success 200 {object} string "Image deleted successfully"
// @Failure 400 {object} string "Invalid request body"
// @Failure 500 {object} string "Server error"
// @Router /api/questions/delete-image/{id} [delete]
func (h *Handler) DeleteImageFromQuestion(c *gin.Context) {
	h.Log.Info("DeleteImageFromQuestion called")
	Id := c.Param("id")
	if len(Id) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "questions ID is required"})
		h.Log.Error("questions ID is required")
		return
	}

	_, err := h.Question.DeleteImageQuestion(c, &question.DeleteImageQuestionRequest{QuestionId: Id})
	if err != nil {
		h.Log.Error("Failed to delete image from question", "error", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Server error"})
		return
	}
	h.Log.Info("DeleteImageFromQuestion ended successfully")
	c.JSON(http.StatusOK, gin.H{"message": "Image deleted successfully"})
}
