package handler

import (
	"OrderEZ/internal/app/service"
)

type ItemHandler struct {
	itemService *service.ItemService
}

func NewGoodHandler(itemService *service.ItemService) *ItemHandler {
	return &ItemHandler{itemService: itemService}
}

//// GetAllGoods 方法用于获取指定页码和每页数量的商品
//func (h *ItemHandler) GetAllGoods(c *gin.Context) {
//	// 从查询参数中获取 page 和 pageSize
//	pageStr := c.DefaultQuery("page", "1")
//	pageSizeStr := c.DefaultQuery("pageSize", "10")
//
//	var page, pageSize int
//	var err error
//
//	// 将 page 和 pageSize 从字符串转换为整数
//	page, err = strconv.Atoi(pageStr)
//	if err != nil {
//		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid page parameter"})
//		return
//	}
//
//	pageSize, err = strconv.Atoi(pageSizeStr)
//	if err != nil {
//		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid pageSize parameter"})
//		return
//	}
//
//	// 调用服务层的 GetAllGoods 方法并传递 page 和 pageSize
//	goods, err := h.itemService.GetAllGoods(page, pageSize)
//	if err != nil {
//		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
//		return
//	}
//
//	c.JSON(http.StatusOK, gin.H{"goods": goods})
//}
//
//// AddGood 处理添加商品的请求
//func (h *itemService) AddGood(c *gin.Context) {
//	var goodData struct {
//		GoodName  string `json:"good_name"`
//		GoodPrice int    `json:"good_price"`
//		GoodStock int    `json:"good_stock"`
//	}
//
//	if err := c.ShouldBindJSON(&goodData); err != nil {
//		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
//		return
//	}
//
//	snowflake, err := util.NewSnowflake(1, 1)
//	if err != nil {
//		log.Fatal(err)
//	}
//	id, err := snowflake.NextID()
//	if err != nil {
//		log.Fatal(err)
//	}
//	addGood := model.Good{
//		GoodID:    id,
//		GoodName:  goodData.GoodName,
//		GoodPrice: goodData.GoodPrice,
//		GoodStock: goodData.GoodStock,
//	}
//
//	// 调用商品服务添加商品
//	if err := h.goodService.AddGood(addGood); err != nil {
//		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add good: " + err.Error()})
//		return
//	}
//
//	// 假设添加成功，返回成功信息
//	c.JSON(http.StatusOK, gin.H{"message": "Good added successfully"})
//}
