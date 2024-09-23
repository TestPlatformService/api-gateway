package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	pb "api/genproto/user"
)

// Register godoc
// @Summary Register user
// @Description create new users
// @Tags user
// @Param info body user.RegisterReq true "User info"
// @Success 200 {object} string "Registered successfully"
// @Failure 400 {object} string "Invalid data"
// @Failure 500 {object} string "Server error"
// @Router /user/register [post]
func (h Handler) Register(c *gin.Context) {
	h.Log.Info("Register is starting")
	req := pb.RegisterRequest{}
	if err := c.BindJSON(&req); err != nil {
		h.Log.Error(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	_, err := h.User.Register(c, &req)
	if err != nil {
		h.Log.Error(err.Error())
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	h.Log.Info("Register ended")
	c.JSON(http.StatusOK, gin.H{"message":"Registered successfully"})
}

// @Summary      Login a user
// @Description  This endpoint logs in a user by checking the credentials and generating JWT tokens.
// @Tags         user
// @Accept       json
// @Produce      json
// @Param        credentials  body      user.LoginRequest  true  "User Login Data"
// @Success      200   {object}  string "Tokens"
// @Failure      400   {object}  user.LoginResponse
// @Failure      401   {object}  string
// @Failure      500   {object}  string
// @Router       /user/login [post]
func (h Handler) Login(c *gin.Context) {
	h.Log.Info("Login starting")
	req := pb.LoginRequest{}

	if err := c.BindJSON(&req); err != nil {
		h.Log.Error(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	
	
}