package handler

import (
	pb "api/genproto/subject"
	"api/model"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// @Summary Create a new Subject
// @Description This endpoint is used to create a new Subject in the system.
// @Tags subjects
// @Accept  json
// @Produce  json
// @Param data body subject.CreateSubjectRequest true "Subject creation request"
// @Success 200 {object} string "Successful Subject creation"
// @Failure 400 {object} model.Error "Bad request: invalid input data"
// @Failure 500 {object} model.Error "Internal server error"
// @Router /api/subjects/create [post]
func (h *Handler) CreateSubject(c *gin.Context) {
	req := pb.CreateSubjectRequest{}
	err := c.ShouldBindJSON(&req)
	if err != nil {
		h.Log.Error(fmt.Sprintf("error while getting information: %v", err))
		c.JSON(http.StatusBadRequest, model.Error{
			Message: "Incorrect data input",
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
		"Message": "Subject Created succesfully",
	})
}

// @Summary Get a Subject by ID
// @Description This endpoint retrieves a specific subject by its ID.
// @Tags subjects
// @Accept  json
// @Produce  json
// @Param id path string true "Subject ID"
// @Success 200 {object} subject.GetSubjectResponse "Successful retrieval of the subject"
// @Failure 400 {object} model.Error "Bad request: invalid input data"
// @Failure 404 {object} model.Error "Subject not found"
// @Failure 500 {object} model.Error "Internal server error"
// @Router /api/subjects/get/{id} [get]
func (h *Handler) GetSubject(c *gin.Context) {
	id := c.Param("id")
	req := pb.GetSubjectRequest{Id: id}

	resp, err := h.Subject.GetSubject(c, &req)
	if err != nil {
		h.Log.Error(fmt.Sprintf("GetSubject request error: %v", err))
		c.JSON(http.StatusInternalServerError, model.Error{
			Message: "Error retrieving subject",
		})
		return
	}

	c.JSON(http.StatusOK, resp)
}

// @Summary Get all Subjects
// @Description This endpoint retrieves all subjects with pagination.
// @Tags subjects
// @Accept  json
// @Produce  json
// @Param limit query int true "Limit of subjects"
// @Param offset query int true "Offset for pagination"
// @Success 200 {object} subject.GetAllSubjectsResponse "Successful retrieval of subjects"
// @Failure 400 {object} model.Error "Bad request: invalid input data"
// @Failure 500 {object} model.Error "Internal server error"
// @Router /api/subjects/getall [get]
func (h *Handler) GetAllSubjects(c *gin.Context) {
	limit, _ := strconv.Atoi(c.Query("limit"))
	offset, _ := strconv.Atoi(c.Query("offset"))

	req := pb.GetAllSubjectsRequest{
		Limit:  int64(limit),
		Offset: int64(offset),
	}

	resp, err := h.Subject.GetAllSubjects(c, &req)
	if err != nil {
		h.Log.Error(fmt.Sprintf("GetAllSubjects request error: %v", err))
		c.JSON(http.StatusInternalServerError, model.Error{
			Message: "Error retrieving subjects",
		})
		return
	}

	c.JSON(http.StatusOK, resp)
}

// @Summary Update a Subject
// @Description This endpoint updates an existing subject by its ID.
// @Tags subjects
// @Accept  json
// @Produce  json
// @Param id path string true "Subject ID"
// @Param data body subject.UpdateSubjectRequest true "Subject update request"
// @Success 200 {object} string "Successful Subject update"
// @Failure 400 {object} model.Error "Bad request: invalid input data"
// @Failure 404 {object} model.Error "Subject not found"
// @Failure 500 {object} model.Error "Internal server error"
// @Router /api/subjects/update/{id} [put]
func (h *Handler) UpdateSubject(c *gin.Context) {
	req := pb.UpdateSubjectRequest{}
	err := c.ShouldBindJSON(&req)
	if err != nil {
		h.Log.Error(fmt.Sprintf("error while binding request data: %v", err))
		c.JSON(http.StatusBadRequest, model.Error{
			Message: "Incorrect data input",
		})
		return
	}

	_, err = h.Subject.UpdateSubject(c, &req)
	if err != nil {
		h.Log.Error(fmt.Sprintf("UpdateSubject request error: %v", err))
		c.JSON(http.StatusInternalServerError, model.Error{
			Message: "Error updating subject",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"Message": "Subject updated successfully",
	})
}

// @Summary Delete a Subject by ID
// @Description This endpoint deletes a subject by its ID.
// @Tags subjects
// @Accept  json
// @Produce  json
// @Param id path string true "Subject ID"
// @Success 200 {object} string "Successful Subject deletion"
// @Failure 400 {object} model.Error "Bad request: invalid input data"
// @Failure 404 {object} model.Error "Subject not found"
// @Failure 500 {object} model.Error "Internal server error"
// @Router /api/subjects/delete/{id} [delete]
func (h *Handler) DeleteSubject(c *gin.Context) {
	id := c.Param("id")
	req := pb.DeleteSubjectRequest{Id: id}

	_, err := h.Subject.DeleteSubject(c, &req)
	if err != nil {
		h.Log.Error(fmt.Sprintf("DeleteSubject request error: %v", err))
		c.JSON(http.StatusInternalServerError, model.Error{
			Message: "Error deleting subject",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"Message": "Subject deleted successfully",
	})
}
