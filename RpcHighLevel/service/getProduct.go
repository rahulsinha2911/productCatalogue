package service

import (
	"highlevel/connection"
	"highlevel/structs"
)

func GetProductService(productID int) (structs.ProductResponse, error) {
	var product structs.Product
	result := connection.ReadDB.Where("product_id = ?", productID).First(&product)
	if result.Error != nil {
		return structs.ProductResponse{}, result.Error
	}
	productResponse := &structs.ProductResponse{
		ProductID: product.ProductID,
		Name:      product.Name,
		Price:     product.Price,
		Quantity:  product.Quantity,
	}
	return *productResponse, nil
}
