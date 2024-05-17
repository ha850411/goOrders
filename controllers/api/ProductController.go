package api

import (
	"fmt"
	"goOrders/conf"
	"goOrders/database"
	"goOrders/models"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

/**
 * @api {get} /api/product 取得商品列表
 * @apiParam {Number} [limit=10] 每頁筆數
 * @apiParam {Number} [page=0] 頁碼
 * @apiParam {String} [keyword] 關鍵字
 * @return {json} 商品列表
 */
func GetProducts(c *gin.Context) {
	db := database.GormConnect()
	productsList := []models.Products{}

	// 接收參數並強制轉 int
	limitQuery := c.Query("limit")
	pageQuery := c.DefaultQuery("page", "1")
	keyword := c.DefaultQuery("keyword", "")
	productTypeQuery := c.DefaultQuery("product_type_id", "")
	limit, _ := strconv.Atoi(limitQuery)
	page, _ := strconv.Atoi(pageQuery)
	productType, _ := strconv.Atoi(productTypeQuery)

	var totalRows int64 // 總筆數

	// 取得啟用的商品
	queryBuilder := db.Model(&models.Products{}).Where("status = 1")
	// 關鍵字
	if keyword != "" {
		queryBuilder.Where("name LIKE ?", "%"+keyword+"%")
	}
	// 產品類別
	if productType > 0 {
		// SELECT pid FROM product_type_detail WHERE tid = ?
		queryBuilder.Where(
			"id IN (?)",
			db.Table(models.ProductTypeDetail{}.GetTableName()).
				Select("pid").Where("tid = ?", productType))
	}
	queryBuilder.Count(&totalRows)

	if limit > 0 {
		queryBuilder.Limit(limit)
	}
	if page > 1 {
		queryBuilder.Offset(page*limit - limit)
	}

	queryBuilder.Debug().Order("sort asc").Find(&productsList)

	data := make([]interface{}, 0)
	for _, v := range productsList {
		db.Table("product_type_detail").
			Select("product_type.id, product_type.name").
			Joins("LEFT JOIN product_type ON product_type_detail.tid = product_type.id").
			Where("product_type_detail.pid = ?", v.Id).Scan(&v.ProductType)
		if v.ProductType == nil {
			v.ProductType = []models.ProductType{}
		}
		data = append(data, v)
	}

	result := map[string]interface{}{
		"page":      page,
		"totalRows": totalRows,
		"data":      data,
	}

	c.JSON(http.StatusOK, result)
}

/**
 * @api {put} /api/product 更新商品
 * @return {json}
 */
func UpdateProduct(c *gin.Context) {

	type UpdateData struct {
		Id            int    `form:"id" required:"true"`
		Name          string `form:"name" required:"true"`
		Amount        string `form:"amount" required:"true"`
		Price         string `form:"price" required:"true"`
		DiscountPrice string `form:"discount_price" required:"true"`
		Contetnt      string `form:"content" required:"true"`
		ProductType   string `form:"product_type" required:"true"`
	}
	var updateData UpdateData
	if err := c.ShouldBind(&updateData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// 上傳檔案
	file, err := c.FormFile("uploadFile")
	if err == nil {
		err2 := c.SaveUploadedFile(
			file,
			fmt.Sprintf(
				"%s/products/%d.png",
				conf.Settings.Common.UPLOADS_PATH,
				updateData.Id,
			),
		)
		if err2 != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err2.Error()})
			return
		}
	}

	db := database.GormConnect()

	// 更新 product 產品資訊
	db.Debug().
		Model(models.Products{}).
		Where("id = ?", updateData.Id).
		Updates(map[string]interface{}{
			"name":           updateData.Name,
			"amount":         updateData.Amount,
			"price":          updateData.Price,
			"discount_price": updateData.DiscountPrice,
			"content":        updateData.Contetnt,
		})

	// 更新 product_type_detail 產品類型
	db.Debug().
		Table(models.ProductTypeDetail{}.GetTableName()).
		Where("pid = ?", updateData.Id).
		Delete(&models.ProductTypeDetail{})

	var temp []models.ProductTypeDetail
	productTypeList := strings.Split(updateData.ProductType, ",")
	for _, v := range productTypeList {
		tid, _ := strconv.Atoi(v)
		temp = append(temp, models.ProductTypeDetail{
			Tid: tid,
			Pid: updateData.Id,
		})
	}
	db.Debug().Table(models.ProductTypeDetail{}.GetTableName()).Create(&temp)

	c.JSON(http.StatusOK, gin.H{"message": "更新成功"})
}

/**
 * @api {post} /api/product 新增商品
 */
func CreateProduct(c *gin.Context) {
	type PostData struct {
		Name          string `form:"name" required:"true"`
		Amount        int    `form:"amount" required:"true"`
		Price         int    `form:"price" required:"true"`
		DiscountPrice int    `form:"discount_price" required:"true"`
		Contetnt      string `form:"content" required:"true"`
		ProductType   string `form:"product_type" required:"true"`
	}

	var postData PostData
	if err := c.ShouldBind(&postData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db := database.GormConnect()

	// 取得目前最大的 sort
	var maxSort int
	db.Debug().Model(&models.Products{}).Select("MAX(sort)").Row().Scan(&maxSort)

	// 新增 product 產品資訊
	insertData := models.Products{
		Name:          postData.Name,
		Amount:        postData.Amount,
		Price:         postData.Price,
		DiscountPrice: postData.DiscountPrice,
		Content:       postData.Contetnt,
		Status:        1,
		Sort:          maxSort,
		CreateTime:    time.Now().Format("2006-01-02 15:04:05"),
		UpdateTime:    time.Now().Format("2006-01-02 15:04:05"),
	}

	db.Debug().Model(&models.Products{}).Create(&insertData)

	// 新增 product_type_detail 產品類型
	var temp []models.ProductTypeDetail
	productTypeList := strings.Split(postData.ProductType, ",")
	for _, v := range productTypeList {
		tid, _ := strconv.Atoi(v)
		temp = append(temp, models.ProductTypeDetail{
			Tid: tid,
			Pid: insertData.Id,
		})
	}
	db.Debug().Table(models.ProductTypeDetail{}.GetTableName()).Create(&temp)

	// 上傳檔案
	file, err := c.FormFile("uploadFile")
	if err == nil {
		err2 := c.SaveUploadedFile(
			file,
			fmt.Sprintf(
				"%s/products/%d.png",
				conf.Settings.Common.UPLOADS_PATH,
				insertData.Id,
			),
		)
		if err2 != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err2.Error()})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"message": "新增成功"})
}

/**
 * @api {delete} /api/product/:id 刪除商品
 */
func DeleteProduct(c *gin.Context) {
	id := c.Param("id")
	db := database.GormConnect()

	// 刪除 product 產品資訊(假刪除)
	db.Debug().Model(&models.Products{}).Where("id = ?", id).Update("status", 0)

	c.JSON(http.StatusOK, gin.H{"message": "刪除成功"})
}

/**
 * @api {patch} /api/product/sort 商品排序
 */
func SortProduct(c *gin.Context) {
	type SortData struct {
		SortIds []string `json:"sortIds" required:"true"`
		Page    int      `json:"page" required:"true"`
		Perpage int      `json:"perpage" required:"true"`
	}

	var sortData SortData
	if err := c.ShouldBindJSON(&sortData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db := database.GormConnect()
	for idx, pid := range sortData.SortIds {
		db.Debug().Model(&models.Products{}).Where("id = ?", pid).Update("sort", (sortData.Page-1)*sortData.Perpage+idx+1)
	}
	c.JSON(http.StatusOK, gin.H{"message": "更新成功"})
}

/**
 * @api {get} /api/product-type 取得商品類型
 */
func GetProductType(c *gin.Context) {
	// 接收參數並強制轉 int
	limitQuery := c.DefaultQuery("limit", "-1")
	pageQuery := c.DefaultQuery("page", "1")
	keyword := c.DefaultQuery("keyword", "")
	limit, _ := strconv.Atoi(limitQuery)
	page, _ := strconv.Atoi(pageQuery)

	db := database.GormConnect()
	productTypeList := []models.ProductType{}
	var totalRows int64 // 總筆數
	queryBuilder := db.Table(models.ProductType{}.GetTableName())
	queryBuilder.Where("status = 1")

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

	queryBuilder.Find(&productTypeList)

	result := map[string]interface{}{
		"page":      page,
		"perpage":   limit,
		"totlaRows": totalRows,
		"data":      productTypeList,
	}
	c.JSON(http.StatusOK, result)
}

/**
 * @api {put} /api/product-type 更新商品類型
 */
func UpdateProductType(c *gin.Context) {
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
	db.Debug().Table(models.ProductType{}.GetTableName()).
		Where("id = ?", updateData.Id).
		Updates(map[string]interface{}{
			"name":        updateData.Name,
			"update_time": time.Now().Format("2006-01-02 15:04:05"),
		})

	c.JSON(http.StatusOK, gin.H{"message": "更新成功"})
}

/**
 * @api {post} /api/product-type 新增商品類型
 */
func CreateProductType(c *gin.Context) {
	type PostData struct {
		Name string `json:"name" required:"true"`
	}

	var postData PostData
	if err := c.ShouldBindJSON(&postData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db := database.GormConnect()
	insertData := models.ProductType{
		Name:       postData.Name,
		Status:     1,
		CreateTime: time.Now().Format("2006-01-02 15:04:05"),
		UpdateTime: time.Now().Format("2006-01-02 15:04:05"),
	}

	db.Debug().Table(models.ProductType{}.GetTableName()).Create(&insertData)

	c.JSON(http.StatusOK, gin.H{"message": "新增成功"})
}

/**
 * @api {delete} /api/product-type/:id 刪除商品類型
 */
func DeleteProductType(c *gin.Context) {
	id := c.Param("id")
	db := database.GormConnect()

	// 刪除 product 產品資訊(假刪除)
	db.Debug().Table(models.ProductType{}.GetTableName()).Where("id = ?", id).Update("status", 0)

	c.JSON(http.StatusOK, gin.H{"message": "刪除成功"})
}
