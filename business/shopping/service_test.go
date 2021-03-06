package shopping_test

import (
	"AltaStore/api/v1/shopping/request"
	"AltaStore/business"
	"AltaStore/business/shopping"
	shoppingMock "AltaStore/business/shopping/mocks"
	"AltaStore/modules/shoppingdetail"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

const (
	userid          = "f9c8c2bf-d525-420e-86e5-4caf03cd8027"
	cartid          = "f9c8c2bf-d525-420e-86e5-4caf03cd8027"
	id              = "f9c8c2bf-d525-420e-86e5-4caf03cd8027"
	code            = "code"
	name            = "name"
	description     = "description"
	productid       = "f9c8c2bf-d525-420e-86e5-4caf03cd8027"
	productname     = "productname"
	ischeckout      = false
	qty         int = 10
	price       int = 100000000
)

var (
	shoppingRepository       shoppingMock.Repository
	shoppingDetailRepository shoppingMock.RepositoryCartDetail
	shoppingService          shopping.Service

	shoppCart            shopping.ShoppCart
	shoppCartDetail      shoppingdetail.ShoppingCartDetail
	detailItemInShopCart request.DetailItemInShopCart
	detailWithProduct    shoppingdetail.ShopCartDetailItemWithProductName
	detailWithProducts   []shoppingdetail.ShopCartDetailItemWithProductName
)

func TestMain(m *testing.M) {
	setup()
	os.Exit(m.Run())
}

func setup() {
	shoppCart = shopping.ShoppCart{
		ID:          id,
		IsCheckOut:  ischeckout,
		Description: description,
	}

	shoppCartDetail = shoppingdetail.ShoppingCartDetail{
		ID:        id,
		ProductId: productid,
		Price:     price,
		Qty:       qty,
	}
	detailWithProduct = shoppingdetail.ShopCartDetailItemWithProductName{
		ShoppingCartDetail: shoppCartDetail,
		ProductName:        productname,
	}
	detailWithProducts = append(detailWithProducts, detailWithProduct)

	detailItemInShopCart = request.DetailItemInShopCart{
		UserId:    userid,
		ProductId: productid,
		Price:     price,
		Qty:       qty,
	}
	shoppingService = shopping.NewService(&shoppingRepository, &shoppingDetailRepository)
}

func TestNewShoppingCart(t *testing.T) {
	t.Run("Expect Insert Shopping Cart Success", func(t *testing.T) {
		shoppingRepository.On("NewShoppingCart",
			mock.AnythingOfType("string"),
			mock.AnythingOfType("string"),
			mock.AnythingOfType("string"),
			mock.AnythingOfType("time.Time"),
		).Return(&shoppCart, nil).Once()

		cart, err := shoppingService.NewShoppingCart(userid, description)

		assert.Nil(t, err)
		assert.Equal(t, cart.ID, shoppCart.ID)
		assert.Equal(t, cart.IsCheckOut, shoppCart.IsCheckOut)
		assert.Equal(t, cart.Description, shoppCart.Description)
	})
	t.Run("Expect Insert Shopping Cart Fail", func(t *testing.T) {
		shoppingRepository.On("NewShoppingCart",
			mock.AnythingOfType("string"),
			mock.AnythingOfType("string"),
			mock.AnythingOfType("string"),
			mock.AnythingOfType("time.Time"),
		).Return(nil, business.ErrInternalServer).Once()

		_, err := shoppingService.NewShoppingCart(userid, description)

		assert.NotNil(t, err)
		assert.Equal(t, err, business.ErrInternalServer)
	})
}

func TestNewItemInShopCart(t *testing.T) {
	t.Run("Expect Insert New Item In Shopping Cart Success", func(t *testing.T) {
		shoppingDetailRepository.On("NewItemInShopCart",
			mock.AnythingOfType("string"),
			mock.AnythingOfType("*shoppingdetail.InsertItemInCartSpec"),
		).Return(nil).Once()

		err := shoppingService.NewItemInShopCart(cartid, &detailItemInShopCart)

		assert.Nil(t, err)
	})
	t.Run("Expect Insert New Item In Shopping Cart Fail", func(t *testing.T) {
		shoppingDetailRepository.On("NewItemInShopCart",
			mock.AnythingOfType("string"),
			mock.AnythingOfType("*shoppingdetail.InsertItemInCartSpec"),
		).Return(business.ErrInternalServer).Once()

		err := shoppingService.NewItemInShopCart(cartid, &detailItemInShopCart)

		assert.NotNil(t, err)
		assert.Equal(t, err, business.ErrInternalServer)
	})
}

func TestModifyItemInShopCart(t *testing.T) {
	t.Run("Expect Modify Item In Shopping Cart Success", func(t *testing.T) {
		shoppingDetailRepository.On("ModifyItemInShopCart",
			mock.AnythingOfType("string"),
			mock.AnythingOfType("string"),
			mock.AnythingOfType("*shoppingdetail.UpdateItemInCartSpec"),
		).Return(nil).Once()

		err := shoppingService.ModifyItemInShopCart(cartid, productid, &detailItemInShopCart)

		assert.Nil(t, err)
	})
	t.Run("Expect Modify Item In Shopping Cart Success", func(t *testing.T) {
		shoppingDetailRepository.On("ModifyItemInShopCart",
			mock.AnythingOfType("string"),
			mock.AnythingOfType("string"),
			mock.AnythingOfType("*shoppingdetail.UpdateItemInCartSpec"),
		).Return(business.ErrInternalServer).Once()

		err := shoppingService.ModifyItemInShopCart(cartid, productid, &detailItemInShopCart)

		assert.NotNil(t, err)
		assert.Equal(t, err, business.ErrInternalServer)
	})
}

func TestDeleteItemInShopCart(t *testing.T) {
	t.Run("Expect Delete Item In Shopping Cart Success", func(t *testing.T) {
		shoppingDetailRepository.On("DeleteItemInShopCart",
			mock.AnythingOfType("string"),
			mock.AnythingOfType("string"),
		).Return(nil).Once()

		err := shoppingService.DeleteItemInShopCart(cartid, productid)

		assert.Nil(t, err)
	})
	t.Run("Expect Delete Item In Shopping Cart Success", func(t *testing.T) {
		shoppingDetailRepository.On("DeleteItemInShopCart",
			mock.AnythingOfType("string"),
			mock.AnythingOfType("string"),
		).Return(business.ErrInternalServer).Once()

		err := shoppingService.DeleteItemInShopCart(cartid, productid)

		assert.NotNil(t, err)
		assert.Equal(t, err, business.ErrInternalServer)
	})
}

func TestGetShoppingCartByUserId(t *testing.T) {
	t.Run("Expect found the Shopping Cart", func(t *testing.T) {
		shoppingRepository.On("GetShoppingCartByUserId",
			mock.AnythingOfType("string"),
		).Return(&shoppCart, nil).Once()

		shopping, err := shoppingService.GetShoppingCartByUserId(userid)

		assert.Nil(t, err)
		assert.NotNil(t, shopping)

		assert.Equal(t, id, shopping.ID)
		assert.Equal(t, ischeckout, shopping.IsCheckOut)
		assert.Equal(t, description, shopping.Description)

	})

	t.Run("Expect Shopping Cart not found", func(t *testing.T) {
		shoppingRepository.On("GetShoppingCartByUserId",
			mock.AnythingOfType("string"),
		).Return(nil, business.ErrInternalServer).Once()

		shopping, err := shoppingService.GetShoppingCartByUserId(userid)

		assert.NotNil(t, err)
		assert.Nil(t, shopping)

		assert.Equal(t, err, business.ErrInternalServer)
	})
}

func TestGetShopCartDetailById(t *testing.T) {
	t.Run("Expect Shopping Cart Not Found", func(t *testing.T) {
		shoppingRepository.On("GetShoppingCartById",
			mock.AnythingOfType("string"),
		).Return(nil, business.ErrNotFound).Once()

		_, err := shoppingService.GetShopCartDetailById(id)

		assert.NotNil(t, err)
		assert.Equal(t, err, business.ErrNotFound)

	})
	t.Run("Expect Detail Shopping Cart Not Found", func(t *testing.T) {
		shoppingRepository.On("GetShoppingCartById",
			mock.AnythingOfType("string"),
		).Return(&shoppCart, nil).Once()
		shoppingDetailRepository.On("GetShopCartDetailById",
			mock.AnythingOfType("string"),
		).Return(nil, business.ErrNotFound).Once()

		_, err := shoppingService.GetShopCartDetailById(id)

		assert.NotNil(t, err)
		assert.Equal(t, err, business.ErrNotFound)

	})
	t.Run("Expect found the Detail Shopping Cart", func(t *testing.T) {
		shoppingRepository.On("GetShoppingCartById",
			mock.AnythingOfType("string"),
		).Return(&shoppCart, nil).Once()
		shoppingDetailRepository.On("GetShopCartDetailById",
			mock.AnythingOfType("string"),
		).Return(&detailWithProducts, nil).Once()

		shopping, err := shoppingService.GetShopCartDetailById(id)

		assert.Nil(t, err)
		assert.NotNil(t, shopping)

		assert.Equal(t, id, shopping.ID)
		assert.Equal(t, description, shopping.Description)
	})
}
