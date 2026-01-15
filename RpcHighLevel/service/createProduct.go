package service

import (
	"highlevel/connection"
	"highlevel/structs"
)

func CreateProduct(product *structs.Product) (structs.ProductResponse, error) {
	result := connection.WriteDB.Create(product)
	if result.Error != nil {
		return structs.ProductResponse{}, result.Error
	}
	product.ProductID = int(result.RowsAffected)
	productResponse := &structs.ProductResponse{
		ProductID: product.ProductID,
		Name:      product.Name,
		Price:     product.Price,
		Quantity:  product.Quantity,
	}
	return *productResponse, nil
}
