package handler

import (
	pb "api/genproto/topic"
	"api/model"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// @Summary Create a new topic
// @Description Ushbu API orqali yangi topic yaratishingiz mumkin.
// @Tags Topic
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param data body topic.CreateTopicReq true "Create Topic request body"
// @Success 200 {object} topic.CreateTopicResp "Muvaffaqiyatli yaratildi"
// @Failure 400 {object} model.Error "Noto'g'ri ma'lumot kiritdingiz"
// @Failure 500 {object} model.Error "Serverda xatolik yuz berdi"
// @Router /api/topics/create [post]
func (h *Handler) CreateTopic(c *gin.Context) {
	req := pb.CreateTopicReq{}
	err := c.ShouldBindJSON(&req)
	if err != nil {
		h.Log.Error(fmt.Sprintf("Ma'lumotlarni olishda xatolik: %v", err))
		c.JSON(http.StatusBadRequest, model.Error{
			Message: "Noto'g'ri ma'lumot kiritdingiz",
		})
		return
	}
	resp, err := h.Topic.CreateTopic(c, &req)
	if err != nil {
		h.Log.Error(fmt.Sprintf("CreateTopic request error: %v", err))
		c.JSON(500, model.Error{
			Message: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, resp)
}

// @Summary Update an existing topic
// @Description Ushbu API orqali mavjud topicni yangilashingiz mumkin.
// @Tags Topic
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param data body topic.UpdateTopicReq true "Update Topic request body"
// @Success 200 {object} topic.UpdateTopicResp "Muvaffaqiyatli yangilandi"
// @Failure 400 {object} model.Error "Noto'g'ri ma'lumot kiritdingiz"
// @Failure 500 {object} model.Error "Serverda xatolik yuz berdi"
// @Router /api/topics/update [put]
func (h *Handler) UpdateTopic(c *gin.Context) {
	req := pb.UpdateTopicReq{}
	err := c.ShouldBindJSON(&req)
	if err != nil {
		h.Log.Error(fmt.Sprintf("Ma'lumotlarni olishda xatolik: %v", err))
		c.JSON(http.StatusBadRequest, model.Error{
			Message: "Noto'g'ri ma'lumot kiritdingiz",
		})
		return
	}
	resp, err := h.Topic.UpdateTopic(c, &req)
	if err != nil {
		h.Log.Error(fmt.Sprintf("UpdateTopic request error: %v", err))
		c.JSON(500, model.Error{
			Message: "Error",
		})
		return
	}
	c.JSON(http.StatusOK, resp)
}

// @Summary Delete an existing topic
// @Description Ushbu API orqali mavjud topicni o'chirishingiz mumkin.
// @Tags Topic
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param data body topic.DeleteTopicReq true "Delete Topic request body"
// @Success 200 {object} topic.DeleteTopicResp "Mavzu muvaffaqiyatli o'chirildi"
// @Failure 400 {object} model.Error "Noto'g'ri ma'lumot kiritdingiz"
// @Failure 500 {object} model.Error "Serverda xatolik yuz berdi"
// @Router /api/topics/delete [delete]
func (h *Handler) DeleteTopic(c *gin.Context) {
	req := pb.DeleteTopicReq{}
	err := c.ShouldBindJSON(&req)
	if err != nil {
		h.Log.Error(fmt.Sprintf("Ma'lumotlarni olishda xatolik: %v", err))
		c.JSON(http.StatusBadRequest, model.Error{
			Message: "Noto'g'ri ma'lumot kiritdingiz",
		})
		return
	}
	resp, err := h.Topic.DeleteTopic(c, &req)
	if err != nil {
		h.Log.Error(fmt.Sprintf("DeleteTopic request error: %v", err))
		c.JSON(500, model.Error{
			Message: "Error",
		})
		return
	}
	c.JSON(http.StatusOK, resp)
}

// GetAllTopics godoc
// @Summary Get all topics
// @Description Bu API barcha mavzularni qaytaradi.
// @Tags Topic
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @Param limit query int false "Limit of topics (optional)" default(1000)
// @Param page query int false "Page for topics (optional)" default(1)
// @Param data query string false "Filter for subjects (subject_id)"
// @Success 200 {object} topic.GetAllTopicsResp "Mavzular ro'yxati"
// @Failure 400 {object} model.Error "Noto'g'ri ma'lumot kiritildi"
// @Failure 500 {object} model.Error "Ichki xatolik"
// @Router /api/topics/getAll [get]
func (h *Handler) GetAllTopics(c *gin.Context) {
	req := pb.GetAllFilter{}
	req.SubjectId = c.Query("subject_id")
	limit := c.Query("limit")
	page := c.Query("page")
	var lim, off int
	lim, err := strconv.Atoi(limit)
	if err != nil {
		lim = 1000
	}
	off, err = strconv.Atoi(page)
	if err != nil {
		off = 1
	}
	resp, err := h.Topic.GetAllTopics(c, &pb.GetAllTopicsReq{
		SubjectId: req.SubjectId,
		Limit:     int32(lim),
		Page:    int32(off),
	})
	if err != nil {
		h.Log.Error(fmt.Sprintf("GetAllTopics request error: %v", err))
		c.JSON(500, model.Error{
			Message: "Error",
		})
		return
	}
	c.JSON(http.StatusOK, resp)
}
