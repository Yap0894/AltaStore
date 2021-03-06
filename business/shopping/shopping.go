package shopping

import (
	"AltaStore/api/v1/shopping/request"
	"AltaStore/modules/shoppingdetail"
	"time"

	"github.com/google/uuid"
)

type ShoppCart struct {
	ID          string
	IsCheckOut  bool
	Description string
	CreatedBy   string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type ItemInCart struct {
	ID          string
	ProductId   string
	ProductName string
	Price       int
	Qty         int
	UpdatedAt   time.Time
}

type ShopCartDetail struct {
	ID          string
	Description string
	CreatedBy   string
	UpdatedAt   time.Time
	Details     []ItemInCart
}

func getShopCartDetailFormat(sum *ShoppCart, details *[]ItemInCart) *ShopCartDetail {
	var shopCartDetail ShopCartDetail
	var itemInCart *ItemInCart

	shopCartDetail.ID = sum.ID
	shopCartDetail.Description = sum.Description
	shopCartDetail.CreatedBy = sum.CreatedBy
	shopCartDetail.UpdatedAt = sum.UpdatedAt

	for _, val := range *details {
		itemInCart = &val
		shopCartDetail.Details = append(shopCartDetail.Details, *itemInCart)
	}

	if shopCartDetail.Details == nil {
		shopCartDetail.Details = []ItemInCart{}
	}

	return &shopCartDetail
}

func insertItemFormat(item *request.DetailItemInShopCart) *shoppingdetail.InsertItemInCartSpec {
	return &shoppingdetail.InsertItemInCartSpec{
		ID:        uuid.NewString(),
		ProductId: item.ProductId,
		Price:     item.Price,
		Qty:       item.Qty,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

func updateItemFormat(productid string, item *request.DetailItemInShopCart) *shoppingdetail.UpdateItemInCartSpec {
	return &shoppingdetail.UpdateItemInCartSpec{
		ProductId: productid,
		Price:     item.Price,
		Qty:       item.Qty,
		UpdatedAt: time.Now(),
	}
}

func toItem(item shoppingdetail.ShopCartDetailItemWithProductName) ItemInCart {
	return ItemInCart{
		ID:          item.ID,
		ProductId:   item.ProductId,
		ProductName: item.ProductName,
		Price:       item.Price,
		Qty:         item.Qty,
		UpdatedAt:   item.UpdatedAt,
	}
}

func toDetailItemInCart(items *[]shoppingdetail.ShopCartDetailItemWithProductName) *[]ItemInCart {
	var details []ItemInCart

	for _, item := range *items {
		details = append(details, toItem(item))
	}

	if details == nil {
		details = []ItemInCart{}
	}

	return &details
}
