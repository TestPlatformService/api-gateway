package handler

import (
	pb "api/genproto/task"
	"api/model"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary Task yaratish
// @Description Yangi task yaratish uchun ma'lumotlarni qabul qiladi
// @Tags tasks
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param request body task.CreateTaskReq true "Task yaratish uchun zarur ma'lumotlar"
// @Success 200 {object} task.CreateTaskResp "Yaratilgan task haqida ma'lumot"
// @Failure 400 {object} model.Error "Noto'g'ri ma'lumot kiritilgan"
// @Failure 500 {object} model.Error "Serverda xato"
// @Router /api/task/create [post]
func(h *Handler) CreateTask(c *gin.Context){
	req := pb.CreateTaskReq{}
	err := c.ShouldBindJSON(&req)
	if err != nil{
		h.Log.Error(fmt.Sprintf("Ma'lumotlarni olishda xatolik: %v", err))
		c.JSON(http.StatusBadRequest, model.Error{
			Message: "Noto'g'ri ma'lumot kiritdingiz",
		})
		return
	}

	resp, err := h.Task.CreateTask(c, &req)
	if err != nil{
		h.Log.Error(fmt.Sprintf("CreateTask request error: %v", err))
		c.JSON(500, model.Error{
			Message: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, resp)
}

// @Summary Task o'chirish
// @Description Berilgan ID bo'yicha taskni o'chiradi
// @Tags tasks
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param request body task.DeleteTaskReq true "O'chirilishi kerak bo'lgan task ID"
// @Success 200 {object} task.DeleteTaskResp "O'chirilgan task haqida ma'lumot"
// @Failure 400 {object} model.Error "Noto'g'ri ma'lumot kiritilgan"
// @Failure 500 {object} model.Error "Serverda xato"
// @Router /api/task/delete [delete]
func(h *Handler) DeleteTask(c *gin.Context){
	req := pb.DeleteTaskReq{}
	err := c.ShouldBindJSON(&req)
	if err != nil{
		h.Log.Error(fmt.Sprintf("Ma'lumotlarni olishda xatolik: %v", err))
		c.JSON(http.StatusBadRequest, model.Error{
			Message: "Noto'g'ri ma'lumot kiritdingiz",
		})
		return
	}

	resp, err := h.Task.DeleteTask(c, &req)
	if err != nil{
		h.Log.Error(fmt.Sprintf("DeleteTask request error: %v", err))
		c.JSON(500, model.Error{
			Message: "Error",
		})
		return
	}
	c.JSON(http.StatusOK, resp)
}

// @Summary Taskni olish
// @Description Berilgan IDlar bo'yicha task ma'lumotlarini olish
// @Tags tasks
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param task_id query string true "Olish uchun task ID"
// @Param user_id query string true "Foydalanuvchi ID"
// @Param topic_id query string true "Mavzu ID"
// @Success 200 {object} task.GetTaskResp "Olingan task haqida ma'lumot"
// @Failure 500 {object} model.Error "Serverda xato"
// @Router /api/task/get [get]
func(h *Handler) GetTask(c *gin.Context){
	req := pb.GetTaskReq{}
	req.TaskId = c.Query("task_id")
	req.UserId = c.Query("user_id")
	req.TopicId = c.Query("topic_id")

	resp, err := h.Task.GetTask(c, &req)
	if err != nil{
		h.Log.Error(fmt.Sprintf("GetTask request error: %v", err))
		c.JSON(500, model.Error{
			Message: "Error",
		})
		return
	}
	c.JSON(http.StatusOK, resp)
}