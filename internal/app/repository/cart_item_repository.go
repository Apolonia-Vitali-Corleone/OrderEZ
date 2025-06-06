package repository

import (
	"OrderEZ/internal/model"
	"gorm.io/gorm"
)

// CartItemRepository the struct
type CartItemRepository struct {
	db *gorm.DB
}

// NewCartItemRepository new a CartItemRepository
func NewCartItemRepository(db *gorm.DB) *CartItemRepository {
	return &CartItemRepository{db: db}
}

// CreateCartItem adds a new cart item
func (r *CartItemRepository) CreateCartItem(cartItem *model.CartItem) error {
	// Verify if the item exists
	var item model.Item
	if err := r.db.Table("oe_item").Where("item_id = ?", cartItem.ItemID).First(&item).Error; err != nil {
		return err
	}
	return r.db.Create(cartItem).Error
}

// GetCartItemByID retrieves a cart item by its ID
func (r *CartItemRepository) GetCartItemByID(cartItemID int64) (*model.CartItem, error) {
	var cartItem model.CartItem
	err := r.db.Where("cart_item_id = ?", cartItemID).First(&cartItem).Error
	if err != nil {
		return nil, err
	}
	return &cartItem, nil
}

// GetCartItemsByCartID retrieves all items in a cart
func (r *CartItemRepository) GetCartItemsByCartID(cartID int64) ([]model.CartItem, error) {
	var cartItems []model.CartItem
	err := r.db.Where("cart_id = ?", cartID).Find(&cartItems).Error
	if err != nil {
		return nil, err
	}
	return cartItems, nil
}

// UpdateCartItem updates a cart item
func (r *CartItemRepository) UpdateCartItem(cartItem *model.CartItem) error {
	// Verify if the item exists
	var item model.Item
	if err := r.db.Table("oe_item").Where("item_id = ?", cartItem.ItemID).First(&item).Error; err != nil {
		return err
	}
	return r.db.Save(cartItem).Error
}

// DeleteCartItem deletes a cart item by its ID
func (r *CartItemRepository) DeleteCartItem(cartItemID int64) error {
	return r.db.Where("cart_item_id = ?", cartItemID).Delete(&model.CartItem{}).Error
}

// DeleteCartItemsByCartID deletes all items in a cart
func (r *CartItemRepository) DeleteCartItemsByCartID(cartID int64) error {
	return r.db.Where("cart_id = ?", cartID).Delete(&model.CartItem{}).Error
}
