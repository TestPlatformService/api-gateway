package handler

import (
	pb "api/genproto/group"
	"api/model"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func(h *Handler) CreateGroup(c *gin.Context){
	req := pb.CreateGroupReq{}
	err := c.ShouldBindJSON(&req)
	if err != nil{
		h.Log.Error(fmt.Sprintf("Ma'lumotlarni olishda xatoli: %v", err))
		c.JSON(http.StatusBadRequest, model.Error{
			Message: "Noto'g'ri ma'lumot kiritildi",
		})
		return 
	}
	resp, err := h.Group.CreateGroup(c, &req)
	if err != nil{
		h.Log.Error(fmt.Sprintf("CreateGroup request error: %v", err))
		c.JSON(500, model.Error{
			Message: "Error",
		})
		return 
	}
	c.JSON(http.StatusOK, resp)
}

func(h *Handler) UpdateGroup(c *gin.Context){
	req := pb.UpdateGroupReq{}
	err := c.ShouldBindJSON(&req)
	if err != nil{
		h.Log.Error(fmt.Sprintf("Ma'lumotlarni olishda xatoli: %v", err))
		c.JSON(http.StatusBadRequest, model.Error{
			Message: "Noto'g'ri ma'lumot kiritildi",
		})
		return
	}
	resp, err := h.Group.UpdateGroup(c, &req)
	if err != nil{
		h.Log.Error(fmt.Sprintf("UpdateGroup request error: %v", err))
		c.JSON(500, model.Error{
			Message: "Error",
		})
		return 
	}
	c.JSON(http.StatusOK, resp)
}

func(h *Handler) DeleteGroup(c *gin.Context){
	req := pb.GroupId{}
	err := c.ShouldBindJSON(&req)
	if err != nil{
		h.Log.Error(fmt.Sprintf("Ma'lumotlarni olishda xatoli: %v", err))
		c.JSON(http.StatusBadRequest, model.Error{
			Message: "Noto'g'ri ma'lumot kiritildi",
		})
		return
	}
	resp, err := h.Group.DeleteGroup(c, &req)
	if err != nil{
		h.Log.Error(fmt.Sprintf("DeleteGroup request error: %v", err))
		c.JSON(500, model.Error{
			Message: "Error",
		})
		return 
	}
	c.JSON(http.StatusOK, resp)
}

func(h *Handler) GetGroupById(c *gin.Context){
	req := pb.GroupId{}
	err := c.ShouldBindJSON(&req)
	if err != nil{
		h.Log.Error(fmt.Sprintf("Ma'lumotlarni olishda xatoli: %v", err))
		c.JSON(http.StatusBadRequest, model.Error{
			Message: "Noto'g'ri ma'lumot kiritildi",
		})
		return
	}
	resp, err := h.Group.GetGroupById(c, &req)
	if err != nil{
		h.Log.Error(fmt.Sprintf("GetGroupById request error: %v", err))
		c.JSON(500, model.Error{
			Message: "Error",
		})
		return 
	}
	c.JSON(http.StatusOK, resp)
}

func(h *Handler) GetAllGroups(c *gin.Context){
	req := &pb.AllGroupsFilter{}
	err := c.ShouldBindJSON(&req)
	if err != nil{
		h.Log.Error(fmt.Sprintf("Ma'lumotlarni olishda xatoli: %v", err))
		c.JSON(http.StatusBadRequest, model.Error{
			Message: "Noto'g'ri ma'lumot kiritildi",
		})
		return
	}
	limit := c.Query("limit")
	offset := c.Query("offset")
	var lim, off int
	lim, err = strconv.Atoi(limit)
	if err != nil{
		lim = 1000
	}
	off, err = strconv.Atoi(offset)
	if err != nil{
		off = 0
	}
	
	resp, err := h.Group.GetAllGroups(c, &pb.GetAllGroupsReq{
		Room: req.Room,
		SubjectId: req.SubjectId,
		Limit: int32(lim),
		Offset: int32(off),
	})
	if err != nil{
		h.Log.Error(fmt.Sprintf("GetAllGroups request error: %v", err))
		c.JSON(500, model.Error{
			Message: "Error",
		})
		return 
	}
	c.JSON(http.StatusOK, resp)
}

func(h *Handler) AddStudentToGroup(c *gin.Context){
	req := pb.AddStudentReq{}
	err := c.ShouldBindJSON(&req)
	if err != nil{
		h.Log.Error(fmt.Sprintf("Ma'lumotlarni olishda xatoli: %v", err))
		c.JSON(http.StatusBadRequest, model.Error{
			Message: "Noto'g'ri ma'lumot kiritildi",
		})
		return
	}
	resp, err := h.Group.AddStudentToGroup(c, &req)
	if err != nil{
		h.Log.Error(fmt.Sprintf("AddStudentToGroup request error: %v", err))
		c.JSON(500, model.Error{
			Message: "Error",
		})
		return 
	
	}
	c.JSON(http.StatusOK, resp)
}

func(h *Handler) DeleteStudentFromGroup(c *gin.Context){
	req := pb.DeleteStudentReq{}
	err := c.ShouldBindJSON(&req)
	if err != nil{
		h.Log.Error(fmt.Sprintf("Ma'lumotlarni olishda xatoli: %v", err))
		c.JSON(http.StatusBadRequest, model.Error{
			Message: "Noto'g'ri ma'lumot kiritildi",
		})
		return
	}
	resp, err := h.Group.DeleteStudentFromGroup(c, &req)
	if err != nil{
		h.Log.Error(fmt.Sprintf("DeleteStudentFromGroup request error: %v", err))
		c.JSON(500, model.Error{
			Message: "Error",
		})
		return 
	}
	c.JSON(http.StatusOK, resp)
}

func(h *Handler) AddTeacherToGroup(c *gin.Context){
	req := pb.AddTeacherReq{}
	err := c.ShouldBindJSON(&req)
	if err != nil{
		h.Log.Error(fmt.Sprintf("Ma'lumotlarni olishda xatoli: %v", err))
		c.JSON(http.StatusBadRequest, model.Error{
			Message: "Noto'g'ri ma'lumot kiritildi",
		})
		return
	}
	resp, err := h.Group.AddTeacherToGroup(c, &req)
	if err != nil{
		h.Log.Error(fmt.Sprintf("AddTeacherToGroup request error: %v", err))
		c.JSON(500, model.Error{
			Message: "Error",
		})
		return 
	}
	c.JSON(http.StatusOK, resp)
}

func(h *Handler) DeleteTeacherFromGroup(c *gin.Context){
	req := pb.DeleteTeacherReq{}
	err := c.ShouldBindJSON(&req)
	if err != nil{
		h.Log.Error(fmt.Sprintf("Ma'lumotlarni olishda xatoli: %v", err))
		c.JSON(http.StatusBadRequest, model.Error{
			Message: "Noto'g'ri ma'lumot kiritildi",
		})
		return
	}
	resp, err := h.Group.DeleteTeacherFromGroup(c, &req)
	if err != nil{
		h.Log.Error(fmt.Sprintf("DeleteTeacherFromGroup request error: %v", err))
		c.JSON(500, model.Error{
			Message: "Error",
		})
		return 
	}
	c.JSON(http.StatusOK, resp)
}

func(h *Handler) GetStudentGroups(c *gin.Context){
	req := pb.StudentId{}
	err := c.ShouldBindJSON(&req)
	if err != nil{
		h.Log.Error(fmt.Sprintf("Ma'lumotlarni olishda xatoli: %v", err))
		c.JSON(http.StatusBadRequest, model.Error{
			Message: "Noto'g'ri ma'lumot kiritildi",
		})
		return
	}
	resp, err := h.Group.GetStudentGroups(c, &req)
	if err != nil{
		h.Log.Error(fmt.Sprintf("GetStudentGroups request error: %v", err))
		c.JSON(500, model.Error{
			Message: "Error",
		})
		return 
	}
	c.JSON(http.StatusOK, resp)
}

func(h *Handler) GetTeacherGroups(c *gin.Context){
	req := pb.TeacherId{}
	err := c.ShouldBindJSON(&req)
	if err != nil{
		h.Log.Error(fmt.Sprintf("Ma'lumotlarni olishda xatoli: %v", err))
		c.JSON(http.StatusBadRequest, model.Error{
			Message: "Noto'g'ri ma'lumot kiritildi",
		})
		return
	}
	resp, err := h.Group.GetTeacherGroups(c, &req)
	if err != nil{
		h.Log.Error(fmt.Sprintf("GetTeacherGroups request error: %v", err))
		c.JSON(500, model.Error{
			Message: "Error",
		})
		return 
	}
	c.JSON(http.StatusOK, resp)
}

func(h *Handler) GetGroupStudents(c *gin.Context){
	req := pb.GroupId{}
	err := c.ShouldBindJSON(&req)
	if err != nil{
		h.Log.Error(fmt.Sprintf("Ma'lumotlarni olishda xatoli: %v", err))
		c.JSON(http.StatusBadRequest, model.Error{
			Message: "Noto'g'ri ma'lumot kiritildi",
		})
		return
	}
	resp, err := h.Group.GetGroupStudents(c, &req)
	if err != nil{
		h.Log.Error(fmt.Sprintf("GetGroupStudents request error: %v", err))
		c.JSON(500, model.Error{
			Message: "Error",
		})
		return 
	}
	c.JSON(http.StatusOK, resp)
}
