package handler

import (
	"OrderEZ/internal/app/service"
)

type CartHandler struct {
	cartService       *service.CartService
	cartDetailService *service.CartItemService
}

func NewCartHandler(cartService *service.CartService, cartDetailService *service.CartItemService) *CartHandler {
	return &CartHandler{cartService: cartService, cartDetailService: cartDetailService}
}

//// GetCart 方法用于获取本人的所有的购物车内容
//func (h *CartHandler) GetCart(c *gin.Context) {
//	// 获取token
//	tokenString := c.GetHeader("Authorization")
//
//	// 验证token
//	claims, err := util.ValidateToken(tokenString)
//	if err != nil {
//		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
//		return
//	}
//
//	// 验证成功，cartID
//	cartID, err := h.cartService.GetCartIDByUserID(int64(claims.UserID))
//	if err != nil {
//		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
//		return
//	}
//
//	// 获取cartDetailList
//	cartDetailList, err := h.cartDetailService.GetCartDetailListByCartID(cartID)
//	if err != nil {
//		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
//		return
//	}
//	c.JSON(http.StatusOK, gin.H{"cart_detail_list": cartDetailList})
//}
