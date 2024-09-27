package handler

import (
	pb "api/genproto/subject"
	"api/model"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary Create a new subject
// @Description This endpoint is used to create a new subject in the system.
// @Tags subjects
// @Accept  json
// @Produce  json
// @Param data body subject.CreateSubjectRequest true "Subject creation request"
// @Success 200 {object} string "Successful subject creation"
// @Failure 400 {object} model.Error "Bad request: invalid input data"
// @Failure 500 {object} model.Error "Internal server error"
// @Router /api/subjects/create [post]
func (h *Handler) CreateSubject(c *gin.Context) {
	req := pb.CreateSubjectRequest{}
	err := c.ShouldBindJSON(&req)
	if err != nil {
		h.Log.Error(fmt.Sprintf("error while creating the subject: %v", err))
		c.JSON(http.StatusBadRequest, model.Error{
			Message: "Incorrect input",
		})
		return
	}
	_, err = h.Subject.CreateSubject(c, &req)
	if err != nil {
		h.Log.Error(fmt.Sprintf("CreateSubject request error: %v", err))
		c.JSON(500, model.Error{
			Message: "Error",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Subject created successfully",
	})
}

// @Summary Get subject by ID
// @Description This endpoint retrieves the details of a subject by its ID.
// @Tags subjects
// @Accept  json
// @Produce  json
// @Param data body subject.GetSubjectRequest true "Subject ID request"
// @Success 200 {object} string "Successful subject retrieval"
// @Failure 400 {object} model.Error "Bad request: invalid input data"
// @Failure 500 {object} model.Error "Internal server error"
// @Router /api/subjects/getById [get]
func (h *Handler) GetSubject(c *gin.Context) {
	req := pb.GetSubjectRequest{}
	err := c.ShouldBindJSON(&req)
	if err != nil {
		h.Log.Error(fmt.Sprintf("error while getting the subject: %v", err))
		c.JSON(http.StatusBadRequest, model.Error{
			Message: "Incorrect data input",
		})
		return
	}
	_, err = h.Subject.GetSubject(c, &req)
	if err != nil {
		h.Log.Error(fmt.Sprintf("GetGroupById request error: %v", err))
		c.JSON(500, model.Error{
			Message: "Error",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Subject created successfully",
	})
}

// @Summary Update an existing subject
// @Description This endpoint is used to update the details of an existing subject.
// @Tags subjects
// @Accept  json
// @Produce  json
// @Param data body subject.UpdateGroupReq true "Group update request"
// @Success 200 {object} string
// @Failure 400 {object} model.Error "Bad request: invalid input data"
// @Failure 500 {object} model.Error "Internal server error"
// @Router /api/subjects/update [put]
func (h *Handler) UpdateSubject(c *gin.Context) {
	req := pb.UpdateSubjectRequest{}
	err := c.ShouldBindJSON(&req)
	if err != nil {
		h.Log.Error(fmt.Sprintf("error while updating the subject: %v", err))
		c.JSON(http.StatusBadRequest, model.Error{
			Message: "Incorrect data input",
		})
		return
	}
	resp, err := h.Subject.UpdateSubject(c, &req)
	if err != nil {
		h.Log.Error(fmt.Sprintf("UpdateSubject request error: %v", err))
		c.JSON(500, model.Error{
			Message: "Error",
		})
		return
	}
	c.JSON(http.StatusOK, resp)
}

// @Summary Delete a subject
// @Description This endpoint is used to delete a subject from the system.
// @Tags subjects
// @Accept  json
// @Produce  json
// @Param data body subject.DeleteSubjectRequest true "subject deletion request"
// @Success 200 {object} subject.DeleteResp "Successful subject deletion"
// @Failure 400 {object} model.Error "Bad request: invalid input data"
// @Failure 500 {object} model.Error "Internal server error"
// @Router /api/subjects/delete [delete]
func(h *Handler) DeleteSubject(c *gin.Context){
	req := pb.DeleteSubjectRequest{}
	err := c.ShouldBindJSON(&req)
	if err != nil{
		h.Log.Error(fmt.Sprintf("error while deleting the subject: %v", err))
		c.JSON(http.StatusBadRequest, model.Error{
			Message: "Incorrect data input",
		})
		return
	}
	resp, err := h.Subject.DeleteSubject(c, &req)
	if err != nil{
		h.Log.Error(fmt.Sprintf("DeleteSubject request error: %v", err))
		c.JSON(500, model.Error{
			Message: "Error",
		})
		return 
	}
	c.JSON(http.StatusOK, resp)
}