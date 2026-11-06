package handler

import (
	"net/http"
	"req-forge-server/internal/model"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type DecompositionHandler struct {
	db *gorm.DB
}

func NewDecompositionHandler(db *gorm.DB) *DecompositionHandler {
	return &DecompositionHandler{db: db}
}

type createRequest struct {
	Requirement string `json:"requirement" binding:"required"`
	Result      string `json:"result" binding:"required"`
	Model       string `json:"model"`
}

// Create 保存一次需求拆解结果
func (h *DecompositionHandler) Create(c *gin.Context) {
	var req createRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	d := model.Decomposition{Requirement: req.Requirement, Result: req.Result, Model: req.Model}
	if err := h.db.Create(&d).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, d)
}

// List 最近 50 条(按创建时间倒序)
func (h *DecompositionHandler) List(c *gin.Context) {
	var items []model.Decomposition
	if err := h.db.Order("created_at desc").Limit(50).Find(&items).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, items)
}

// Get 单条详情
func (h *DecompositionHandler) Get(c *gin.Context) {
	var d model.Decomposition
	if err := h.db.First(&d, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "not found"})
		return
	}
	c.JSON(http.StatusOK, d)
}

// Delete 删除一条
func (h *DecompositionHandler) Delete(c *gin.Context) {
	if err := h.db.Delete(&model.Decomposition{}, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"ok": true})
}
