package controllers

import (
	"errors"
)

var (
	ErrCantFindProducts   = errors.New("can not find the product")
	ErrCantDecodeProducts = errors.New("can not decode the product")
	ErrCantUpdateUser     = errors.New("can not update user")
	ErrCantRemoveItemCart = errors.New("can not remove item from the cart")
	ErrUserIdIsNotValid   = errors.New("user id is not valid")
	ErrCantGetItem        = errors.New("can not get item")
	ErrCantBuyCantItem    = errors.New("can not buy the cart item")
)

func AddProductToCart() {

}

func RemoveCartItem() {

}

func BuyItemFromCart() {

}

func InstantBuyer() {

}
