package v1

import (
	"MiMengCore/global"
	"MiMengCore/model"
	"errors"
	"gorm.io/gorm"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Content 表示公告内容的类型
type Content = model.Content

// GetNoticeHandler 处理获取公告的GET请求
func GetNoticeHandler(c *gin.Context) {
	var content Content
	// 查询类型为"notice"的公告
	if err := global.DB.Where("type = ?", "notice").First(&content).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// 如果记录未找到，则创建一个新的公告记录
			content = Content{
				Type:    "notice",
				Content: "默认公告内容", // 设置默认公告内容
			}
			if err := global.DB.Create(&content).Error; err != nil {
				// 如果创建出错，返回内部服务器错误状态码及错误信息
				c.JSON(http.StatusInternalServerError, gin.H{
					"success": false,
					"message": "创建公告失败",
					"error":   err.Error(),
				})
				return
			}
			// 创建成功，返回新创建的公告内容
			c.JSON(http.StatusOK, gin.H{
				"success": true,
				"message": "公告不存在，已创建默认公告",
				"data":    content,
			})
			return
		} else {
			// 如果查询出错，返回内部服务器错误状态码及错误信息
			c.JSON(http.StatusInternalServerError, gin.H{
				"success": false,
				"message": "获取公告失败",
				"error":   err.Error(),
			})
			return
		}
	}
	// 如果查询成功，返回公告内容
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "获取公告成功",
		"data":    content,
	})
}

// UpdateNoticeHandler 处理更新公告的PUT请求
func UpdateNoticeHandler(c *gin.Context) {
	var content Content
	// 将请求体中的JSON数据绑定到content变量，并检查是否有错误
	if err := c.ShouldBindJSON(&content); err != nil {
		// 如果绑定出错，返回错误的请求状态码及错误信息
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "请求数据绑定失败",
			"error":   err.Error(),
		})
		return
	}

	// 查询类型为"notice"的公告
	if err := global.DB.Where("type = ?", "notice").First(&content).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// 如果记录未找到，则创建一个新的公告记录
			content = Content{
				Type:    "notice",
				Content: content.Content, // 使用请求中提供的公告内容
			}
			if err := global.DB.Create(&content).Error; err != nil {
				// 如果创建出错，返回内部服务器错误状态码及错误信息
				c.JSON(http.StatusInternalServerError, gin.H{
					"success": false,
					"message": "创建公告失败",
					"error":   err.Error(),
				})
				return
			}
			// 创建成功，返回新创建的公告内容
			c.JSON(http.StatusOK, gin.H{
				"success": true,
				"message": "公告不存在，已创建新的公告",
				"data":    content,
			})
			return
		} else {
			// 如果查询出错，返回内部服务器错误状态码及错误信息
			c.JSON(http.StatusInternalServerError, gin.H{
				"success": false,
				"message": "获取公告失败",
				"error":   err.Error(),
			})
			return
		}
	}

	// 如果查询成功，更新公告内容
	result := global.DB.Model(&Content{}).Where("type = ?", "notice").Updates(map[string]interface{}{
		"Content": content.Content,
	})
	// 检查更新操作是否有错误发生
	if result.Error != nil {
		// 如果更新出错，返回内部服务器错误状态码及错误信息
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "更新公告失败",
			"error":   result.Error.Error(),
		})
		return
	}

	// 如果更新成功，返回成功状态及消息
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "更新公告成功",
	})
}
