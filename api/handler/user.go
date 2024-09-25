package handler

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"

	"api/api/token"
	pb "api/genproto/user"

	"github.com/gin-gonic/gin"
)

// Register godoc
// @Summary Register user
// @Description Create a new user
// @Tags user
// @Security ApiKeyAuth
// @Param info body user.RegisterRequest true "User info"
// @Success 200 {object} string "Registered successfully"
// @Failure 400 {object} string "Invalid request body"
// @Failure 500 {object} string "Server error"
// @Router /api/user/register [post]
func (h Handler) Register(c *gin.Context) {
	h.Log.Info("Register is starting")
	req := pb.RegisterRequest{}

	if err := c.BindJSON(&req); err != nil {
		h.Log.Error("Invalid request body", "error", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	_, err := h.User.Register(c, &req)
	if err != nil {
		h.Log.Error("Failed to register user", "error", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Server error"})
		return
	}

	h.Log.Info("Register ended successfully")
	c.JSON(http.StatusOK, gin.H{"message": "Registered successfully"})
}

// @Summary      Login a user
// @Description  This endpoint logs in a user by checking the credentials and generating JWT tokens.
// @Tags         user
// @Accept       json
// @Produce      json
// @Param        credentials  body user.LoginRequest  true  "User Login Data"
// @Success      200   {object}  user.LoginResponse "Tokens"
// @Failure      400   {object}  string "Invalid request body"
// @Failure      401   {object}  string "Unauthorized"
// @Failure      500   {object}  string "Server error"
// @Router       /all/user/login [post]
func (h Handler) Login(c *gin.Context) {
	h.Log.Info("Login starting")
	req := pb.LoginRequest{}

	if err := c.BindJSON(&req); err != nil {
		h.Log.Error("Invalid request body", "error", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	res, err := h.User.Login(c, &req)
	if err != nil {
		h.Log.Error("Login failed", "error", err.Error())
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	err = token.GeneratedAccessJWTToken(res)

	if err != nil {
		h.Log.Error(err.Error())
		c.JSON(500, gin.H{"error": err.Error()})
	}
	err = token.GeneratedRefreshJWTToken(res)
	if err != nil {
		h.Log.Error(err.Error())
		c.JSON(500, gin.H{"error": err.Error()})
	}

	h.Log.Info("Login ended successfully")
	c.JSON(http.StatusOK, res)
}

// @Summary      Get user profile
// @Description  This endpoint retrieves user profile.
// @Tags         user
// @Accept       json
// @Produce      json
// @Security     ApiKeyAuth
// @Success      200    {object}  user.GetProfileResponse
// @Failure      400    {object}  string "Invalid token"
// @Failure      500    {object}  string "Server error"
// @Router       /api/user/getprofile [get]
func (h Handler) GetProfile(c *gin.Context) {
	h.Log.Info("GetProfile starting")
	tokenn := c.GetHeader("Authorization")

	id, _, err := token.GetUserInfoFromAccessToken(tokenn)
	if err != nil {
		h.Log.Error("Invalid token", "error", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid token"})
		return
	}

	req := pb.GetProfileRequest{
		Id: id,
	}

	res, err := h.User.GetProfile(c, &req)
	if err != nil {
		h.Log.Error("Error while retrieving profile.", "error", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	h.Log.Info("GetProfile ended")
	c.JSON(http.StatusOK, res)
}

// @Summary      Get all users
// @Description  Retrieve all users with optional filters such as role, group, subject, teacher, and pagination.
// @Tags         user
// @Accept       json
// @Produce      json
// @Security     ApiKeyAuth
// @Param        Role         query    string  false  "Role to filter by"
// @Param        Group        query    string  false  "Group to filter by"
// @Param        Subject      query    string  false  "Subject to filter by"
// @Param        Teacher      query    string  false  "Teacher ID to filter by"
// @Param        HhId        query    string  false  "Unique household ID to filter by"
// @Param        PhoneNumber query    string  false  "Phone number to filter by"
// @Param        Gender       query    string  false  "Gender to filter by"
// @Param        Limit        query    int     false  "Number of users to return per page" default(10)
// @Param        Offset       query    int     false  "Pagination offset"
// @Success      200   {object}  user.GetAllUsersResponse  "Successfully retrieved users"
// @Failure      400   {object}  string "Invalid request parameters"
// @Failure      500   {object}  string "Internal server error"
// @Router       /api/user/all [get]
func (h Handler) GetAllUsers(c *gin.Context) {
	h.Log.Info("GetAllUsers starting")

	req := pb.GetAllUsersRequest{}

	if err := c.ShouldBindQuery(&req); err != nil {
		h.Log.Error("Invalid query parameters", "error", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid query parameters"})
		return
	}
	fmt.Print("1\n", req.HhId, "\n2")

	res, err := h.User.GetAllUsers(c, &req)
	if err != nil {
		h.Log.Error("Failed to get users", "error", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	h.Log.Info("GetAllUsers ended successfully")
	c.JSON(http.StatusOK, res)
}

// @Security ApiKeyAuth
// @Summary Update User
// @Description Update User profile
// @Tags user
// @Param info body user.UpdateProfileRequest true "info"
// @Success 200 {object} string "message"
// @Failure 400 {object} string "Invalid data"
// @Failure 500 {object} string "Server error"
// @Router /api/user/updateprofile [put]
func (h *Handler) UpdateProfile(c *gin.Context) {
	h.Log.Info("UpdateUser started")
	tokenn := c.GetHeader("Authorization")
	id, _, err := token.GetUserInfoFromAccessToken(tokenn)
	if err != nil {
		h.Log.Error(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var req pb.UpdateProfileRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		h.Log.Error("Failed to bind JSON", "error", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}
	req.Id = id
	_, err = h.User.UpdateProfile(c, &req)
	if err != nil {
		h.Log.Error("Failed to update user", "error", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	h.Log.Info("UpdateUser ended")
	c.JSON(http.StatusOK, gin.H{"message": "User profile updated"})
}

// @Summary      Update User by Admin
// @Description  Update User Profile by Admin
// @Tags         user
// @Security     ApiKeyAuth
// @Param        info body user.UpdateProfileAdminRequest true "info"
// @Success      200 {object} string "User profile updated"
// @Failure      400 {object} string "Invalid request body"
// @Failure      500 {object} string "Server error"
// @Router       /api/user/update [put]
func (h *Handler) UpdateProfileAdmin(c *gin.Context) {
	h.Log.Info("UpdateProfileAdmin started")
	tokenn := c.GetHeader("Authorization")

	id, _, err := token.GetUserInfoFromAccessToken(tokenn)
	if err != nil {
		h.Log.Error("Invalid token", "error", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid token"})
		return
	}

	var req pb.UpdateProfileAdminRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		h.Log.Error("Failed to bind JSON", "error", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	req.Id = id

	_, err = h.User.UpdateProfileAdmin(c, &req)
	if err != nil {
		h.Log.Error("Failed to update user", "error", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Server error"})
		return
	}

	h.Log.Info("UpdateProfileAdmin ended successfully")
	c.JSON(http.StatusOK, gin.H{"message": "User profile updated"})
}

// @Summary      Delete User Profile
// @Description  Marks a user profile as deleted by setting the deleted_at timestamp.
// @Tags         user
// @Security     ApiKeyAuth
// @Param        id   path      string  true  "User ID to delete"
// @Success      200  {object}  string "success"
// @Failure      400  {object}  string "Invalid request"
// @Failure      404  {object}  string "User not found"
// @Failure      500  {object}  string "Internal server error"
// @Router       /api/user/delete/{id} [delete]
func (h Handler) DeleteProfile(c *gin.Context) {
	h.Log.Info("DeleteProfile starting")
	id := c.Param("id")

	req := pb.DeleteProfileRequest{
		Id: id,
	}

	_, err := h.User.DeleteProfile(c, &req)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
			return
		}
		h.Log.Error("Failed to delete user profile", "error", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	h.Log.Info("DeleteProfile ended successfully")
	c.JSON(http.StatusOK, gin.H{"message": "User profile deleted"})
}
