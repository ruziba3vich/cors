package handlers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ruziba3vich/cors/internal/items/models"
	"github.com/ruziba3vich/cors/internal/items/repo"
)

type (
	Handler struct {
		logger  *log.Logger
		service repo.UserServiceRepo
	}

	CORSHandler struct {
		logger  *log.Logger
		service repo.CORSRepo
	}
)

func New(logger *log.Logger, service repo.UserServiceRepo) *Handler {
	return &Handler{
		logger:  logger,
		service: service,
	}
}

func NewCORSHandler(logger *log.Logger, service repo.CORSRepo) *CORSHandler {
	return &CORSHandler{
		logger:  logger,
		service: service,
	}
}

func (h *Handler) Register(c *gin.Context) {
	var req models.User
	if err := c.ShouldBindJSON(&req); err != nil {
		h.logger.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	response, err := h.service.RegisterUser(c, &req)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"response": response})
}

func (h *Handler) Login(c *gin.Context) {
	var req models.User
	if err := c.ShouldBindJSON(&req); err != nil {
		h.logger.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	response, err := h.service.LoginUser(c, &req)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"response": response})
}

func (cors *CORSHandler) AddOriginToUser(c *gin.Context) {
	username := c.MustGet("username").(string)
	var req models.CreateOriginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		cors.logger.Println(err)
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	req.Username = username
	response, err := cors.service.CreateOrigin(c, &req)
	if err != nil {
		cors.logger.Println(err)
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusCreated, gin.H{"response": response})
}

func (cors *CORSHandler) GetOriginsByUsername(c *gin.Context) {
	username := c.MustGet("username").(string)

	response, err := cors.service.GetOriginsByUsername(c, username)
	if err != nil {
		cors.logger.Println(err)
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusCreated, gin.H{"response": response})
}

func (cors *CORSHandler) DeleteOriginByUsername(c *gin.Context) {
	username := c.MustGet("username").(string)
	var req models.CreateOriginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		cors.logger.Println(err)
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	response, err := cors.service.DeleteOrigin(c, username, req.Origin)
	if err != nil {
		cors.logger.Println(err)
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusCreated, gin.H{"response": response})
}
