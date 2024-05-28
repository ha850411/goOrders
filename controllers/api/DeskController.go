package api

import (
	"goOrders/database"
	"goOrders/models"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

/**
 * @api {get} /api/desk 取得桌號列表
 */
func GetDesks(c *gin.Context) {
	// 接收參數並強制轉 int
	limitQuery := c.DefaultQuery("limit", "-1")
	pageQuery := c.DefaultQuery("page", "1")
	keyword := c.DefaultQuery("keyword", "")
	limit, _ := strconv.Atoi(limitQuery)
	page, _ := strconv.Atoi(pageQuery)
	var totalRows int64 // 總筆數

	db := database.GormConnect()
	desksList := []models.Desks{}
	queryBuilder := db.Debug().Model(models.Desks{}).Where("status = 1")
	// 關鍵字
	if keyword != "" {
		queryBuilder.Where("name LIKE ?", "%"+keyword+"%")
	}
	queryBuilder.Count(&totalRows)

	if limit > 0 {
		queryBuilder.Limit(limit)
	}

	if page > 1 {
		queryBuilder.Offset(page*limit - limit)
	}
	queryBuilder.Find(&desksList)

	c.JSON(http.StatusOK, map[string]interface{}{
		"page":      page,
		"perpage":   limit,
		"totlaRows": totalRows,
		"data":      desksList,
	})
}

/**
 * @api {post} /api/desk 新增桌號
 */
func CreateDesk(c *gin.Context) {
	type PostData struct {
		Name string `json:"name" required:"true"`
	}

	var postData PostData
	if err := c.ShouldBindJSON(&postData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db := database.GormConnect()

	insertData := models.Desks{
		Id:         int(uuid.New().ID()),
		Name:       postData.Name,
		Status:     1,
		CreateTime: time.Now().Format("2006-01-02 15:04:05"),
		UpdateTime: time.Now().Format("2006-01-02 15:04:05"),
	}

	db.Debug().Model(models.Desks{}).Create(&insertData)

	c.JSON(http.StatusOK, gin.H{"message": "新增成功"})
}

/**
 * @api {put} /api/desk 更新桌號
 */
func UpdateDesk(c *gin.Context) {
	type UpdateData struct {
		Id   int    `json:"id" required:"true"`
		Name string `json:"name" required:"true"`
	}

	var updateData UpdateData
	if err := c.ShouldBindJSON(&updateData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db := database.GormConnect()
	db.Debug().Model(models.Desks{}).
		Where("id = ?", updateData.Id).
		Updates(map[string]interface{}{
			"name":        updateData.Name,
			"update_time": time.Now().Format("2006-01-02 15:04:05"),
		})

	c.JSON(http.StatusOK, gin.H{"message": "更新成功"})
}

/**
 * @api {delete} /api/desk/:id 刪除桌號
 */
func DeleteDesk(c *gin.Context) {
	id := c.Param("id")
	db := database.GormConnect()

	// 刪除 product 產品資訊(假刪除)
	db.Debug().Model(models.Desks{}).Where("id = ?", id).Update("status", 0)

	c.JSON(http.StatusOK, gin.H{"message": "刪除成功"})
}
