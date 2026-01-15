package service

import (
	"highlevel/connection"
	"highlevel/structs"
)

func GetProductListService() ([]structs.ProductResponse, error) {
	var products []structs.Product
	result := connection.ReadDB.Find(&products)
	if result.Error != nil {
		return nil, result.Error
	}
	productResponses := make([]structs.ProductResponse, len(products))
	for i, product := range products {
		productResponses[i] = structs.ProductResponse{
			ProductID: product.ProductID,
			Name:      product.Name,
			Price:     product.Price,
			Quantity:  product.Quantity,
		}
	}
	return productResponses, nil
}
