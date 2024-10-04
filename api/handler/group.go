package handler

import (
	pb "api/genproto/group"
	"api/model"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// @Summary Create a new group
// @Description This endpoint is used to create a new group in the system.
// @Tags groups
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @Param data body group.CreateGroupReq true "Group creation request"
// @Success 200 {object} group.CreateGroupResp "Successful group creation"
// @Failure 400 {object} model.Error "Bad request: invalid input data"
// @Failure 500 {object} model.Error "Internal server error"
// @Router /api/groups/create [post]
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

// @Summary Update an existing group
// @Description This endpoint is used to update the details of an existing group.
// @Tags groups
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @Param data body group.UpdateGroupReq true "Group update request"
// @Success 200 {object} group.UpdateGroupResp "Successful group update"
// @Failure 400 {object} model.Error "Bad request: invalid input data"
// @Failure 500 {object} model.Error "Internal server error"
// @Router /api/groups/update [put]
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

// @Summary Delete a group
// @Description This endpoint is used to delete a group from the system.
// @Tags groups
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @Param data body group.GroupId true "Group deletion request"
// @Success 200 {object} group.DeleteResp "Successful group deletion"
// @Failure 400 {object} model.Error "Bad request: invalid input data"
// @Failure 500 {object} model.Error "Internal server error"
// @Router /api/groups/delete [delete]
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

// @Summary Get group by ID
// @Description This endpoint retrieves the details of a group by its ID.
// @Tags groups
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @Param data path string true "Group ID request (group_id)"
// @Success 200 {object} group.Group "Successful group retrieval"
// @Failure 400 {object} model.Error "Bad request: invalid input data"
// @Failure 500 {object} model.Error "Internal server error"
// @Router /api/groups/getById/{group_id} [get]
func(h *Handler) GetGroupById(c *gin.Context){
	req := pb.GroupId{}
	req.Id = c.Param("group_id")
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

// @Summary Get all groups
// @Description This endpoint retrieves all groups with optional filters like room and subject.
// @Tags groups
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @Param room query string false "Room filter"
// @Param subject_id query string false "Subject ID filter"
// @Param limit query string false "Limit for pagination"
// @Param page query string false "page for pagination"
// @Success 200 {object} group.GetAllGroupsResp "Successful group retrieval"
// @Failure 400 {object} model.Error "Bad request: invalid input data"
// @Failure 500 {object} model.Error "Internal server error"
// @Router /api/groups/getAll [get]
func(h *Handler) GetAllGroups(c *gin.Context){
	req := &pb.AllGroupsFilter{}
	req.Room = c.Query("room")
	req.SubjectId = c.Query("subject_id")
	limit := c.Query("limit")
	page := c.Query("page")
	var lim, off int
	lim, err := strconv.Atoi(limit)
	if err != nil{
		lim = 1000
	}
	off, err = strconv.Atoi(page)
	if err != nil{
		off = 0
	}
	
	resp, err := h.Group.GetAllGroups(c, &pb.GetAllGroupsReq{
		Room: req.Room,
		SubjectId: req.SubjectId,
		Limit: int32(lim),
		Page: int32(off),
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

// @Summary Add student to group
// @Description This endpoint allows adding a student to a specific group.
// @Tags groups
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @Param data body group.AddStudentReq true "Student addition request"
// @Success 200 {object} group.AddStudentResp "Successful student addition"
// @Failure 400 {object} model.Error "Bad request: invalid input data"
// @Failure 500 {object} model.Error "Internal server error"
// @Router /api/groups/add-student [post]
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

// @Summary Delete student from group
// @Description This endpoint allows deleting a student from a specific group.
// @Tags groups
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @Param data body group.DeleteStudentReq true "Student deletion request"
// @Success 200 {object} group.DeleteResp "Successful student deletion"
// @Failure 400 {object} model.Error "Bad request: invalid input data"
// @Failure 500 {object} model.Error "Internal server error"
// @Router /api/groups/delete-student [delete]
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

// @Summary Add teacher to group
// @Description This endpoint allows adding a teacher to a specific group.
// @Tags groups
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @Param data body group.AddTeacherReq true "Teacher addition request"
// @Success 200 {object} group.AddTeacherResp "Successful teacher addition"
// @Failure 400 {object} model.Error "Bad request: invalid input data"
// @Failure 500 {object} model.Error "Internal server error"
// @Router /api/groups/add-teacher [post]
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

// @Summary Delete teacher from group
// @Description This endpoint allows deleting a teacher from a specific group.
// @Tags groups
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @Param data body group.DeleteTeacherReq true "Teacher deletion request"
// @Success 200 {object} group.DeleteResp "Successful teacher deletion"
// @Failure 400 {object} model.Error "Bad request: invalid input data"
// @Failure 500 {object} model.Error "Internal server error"
// @Router /api/groups/delete-teacher [delete]
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

// @Summary Get student groups
// @Description This endpoint retrieves the list of groups a specific student belongs to.
// @Tags groups
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @Param hh_id path string true "Student ID (hh_id)"
// @Success 200 {object} group.StudentGroups "Successful retrieval of student groups"
// @Failure 400 {object} model.Error "Bad request: invalid input data"
// @Failure 500 {object} model.Error "Internal server error"
// @Router /api/groups/student-groups/{hh_id} [get]
func(h *Handler) GetStudentGroups(c *gin.Context){
	req := pb.StudentId{}
	req.HhId = c.Param("hh_id")
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

// @Summary Get teacher groups
// @Description This endpoint retrieves the list of groups a specific teacher belongs to.
// @Tags groups
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @Param data path string true "Teacher ID (id)"
// @Success 200 {object} group.TeacherGroups "Successful retrieval of teacher groups"
// @Failure 400 {object} model.Error "Bad request: invalid input data"
// @Failure 500 {object} model.Error "Internal server error"
// @Router /api/groups/teacher-groups/{id} [get]
func(h *Handler) GetTeacherGroups(c *gin.Context){
	req := pb.TeacherId{}
	req.Id = c.Param("id")
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

// @Summary Get students of a group
// @Description This endpoint retrieves the list of students in a specific group.
// @Tags groups
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @Param data path string true "Group ID (group_id)"
// @Success 200 {object} group.GroupStudents "Successful retrieval of group students"
// @Failure 400 {object} model.Error "Bad request: invalid input data"
// @Failure 500 {object} model.Error "Internal server error"
// @Router /api/groups/students/{group_id} [get]
func(h *Handler) GetGroupStudents(c *gin.Context){
	req := pb.GroupId{}
	req.Id = c.Param("group_id")
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
