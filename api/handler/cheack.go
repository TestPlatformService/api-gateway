package handler

import (
	"api/model"
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// InputOutput - I/O ma'lumotlari uchun tuzilma
type InputOutput struct {
	In  string `json:"in"`
	Out string `json:"out"`
}

// RunRequest - Kodni tekshirish uchun so'rov tuzilmasi
type RunRequest struct {
	Code        string        `json:"code"`        // Kod
	Lang        string        `json:"lang"`        // Dasturlash tili
	TimeLimit   int32         `json:"timeLimit"`   // Vaqt cheklovi
	MemoryLimit int64         `json:"memoryLimit"` // Xotira cheklovi
	IO          []InputOutput `json:"io"`          // I/O ma'lumotlari
}

// ProxyChecker - Kodni tekshirish uchun API
// @Summary Check code with the checker service
// @Description This API checks the submitted code using the checker service and returns the result via SSE.
// @Accept json
// @Produce text/event-stream
// @Param request body RunRequest true "Request body containing code, language, limits, and I/O"
// @Success 200 {string} string "Event stream with results"
// @Failure 400 {object} model.Error "Invalid request"
// @Failure 500 {object} model.Error "Internal server error"
// @Router /api/check/submit [post]
func (h *Handler) ProxyChecker(c *gin.Context) {
	var req RunRequest

	// Clientdan request ma'lumotlarni olish
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, model.Error{Message: "Invalid request"})
		return
	}

	// Checker service'ga request yuborish
	requestBody, err := json.Marshal(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.Error{Message: "Failed to encode request body"})
		return
	}

	// Checker service bilan bog'lanish
	resp, err := http.Post("http://3.121.214.21:50054/check", "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.Error{Message: "Failed to connect to checker service"})
		return
	}
	defer resp.Body.Close()

	// SSE uchun javob sarlavhalarini sozlash
	c.Writer.Header().Set("Content-Type", "text/event-stream")
	c.Writer.Header().Set("Cache-Control", "no-cache")
	c.Writer.Header().Set("Connection", "keep-alive")

	// Checker service'dan SSE orqali natijalarni qaytarish
	for {
		// Har bir chunk'ni o'qish
		buffer := make([]byte, 1024)
		n, err := resp.Body.Read(buffer)
		if n > 0 {
			// Qaytarilgan chunk'ni yozish va SSE orqali uzatish
			c.Writer.Write(buffer[:n])
			c.Writer.(http.Flusher).Flush()
		}
		if err != nil {
			// Agar end bo'lsa yoki xatolik yuz bersa loopdan chiqish
			if err == io.EOF {
				break
			}
			log.Printf("Error reading response: %v", err)
			break
		}
	}
}
