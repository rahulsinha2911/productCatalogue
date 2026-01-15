package handler

import (
	"context"
	"errors"

	"highlevel/proto/product/v1/productv1connect"
	"highlevel/service"
	"highlevel/structs"

	"connectrpc.com/connect"
)

// UserServiceHandler implements the Connect RPC UserService
type ProductServiceHandler struct {
	productv1connect.UnimplementedProductServiceHandler
}

// NewUserServiceHandler creates a new UserServiceHandler
func NewProductServiceHandler() *ProductServiceHandler {
	return &ProductServiceHandler{}
}

// // GetUser retrieves user information by user ID
// // func (h *UserServiceHandler) GetUser(
// // 	ctx context.Context,
// // 	req *connect.Request[userv1.GetUserRequest],
// // ) (*connect.Response[userv1.GetUserResponse], error) {
// // 	userID := req.Msg.UserId
// // 	if userID == "" {
// // 		return nil, connect.NewError(
// // 			connect.CodeInvalidArgument,
// // 			errors.New("user_id is required"),
// // 		)
// // 	}

// // 	// Use service to get user info
// // 	userInfo, err := service.GetUserService(userID)
// // 	if err != nil {
// // 		connectCode := connect.CodeNotFound
// // 		if err.Error() != "user not found in database" {
// // 			connectCode = connect.CodeInternal
// // 		}
// // 		return nil, connect.NewError(connectCode, err)
// // 	}

// // 	// Convert to Connect response
// // 	response := &userv1.GetUserResponse{
// // 		UserId:  userInfo.UserID,
// // 		EmailId: userInfo.EmailID,
// // 		Name:    userInfo.Name,
// // 		Role:    userInfo.Role,
// // 	}

// // 	return connect.NewResponse(response), nil
// }

func (h *ProductServiceHandler) CreateProduct(ctx context.Context, req *connect.Request[productv1connect.CreateProductRequest],
) (*connect.Response[productv1connect.CreateProductResponse], error) {

	// string name = 1;
	// float64 price = 2;
	// int quantity = 3;

	name := req.Msg.Name
	price := req.Msg.Price
	quantity := req.Msg.Quantity

	if name == "" || price < 0 || quantity < 0 {
		return nil, connect.NewError(connect.CodeInvalidArgument, errors.New("invalid product data"))
	}

	product := &structs.Product{
		Name:     name,
		Price:    price,
		Quantity: quantity,
	}

	productResponse, err := service.CreateProduct(product)
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	response := &productv1connect.CreateProductResponse{
		ProductId: productResponse.ProductID,
		Name:      productResponse.Name,
		Price:     productResponse.Price,
		Quantity:  productResponse.Quantity,
	}
	return connect.NewResponse(response), nil
}

func (h *ProductServiceHandler) GetProductList(ctx context.Context, req *connect.Request[productv1connect.GetProductListRequest],
) (*connect.Response[productv1connect.GetProductListResponse], error) {

	productList, err := service.GetProductListService()
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	response := &productv1connect.GetProductListResponse{
		ProductList: productList,
	}
	return connect.NewResponse(response), nil
}

func (h *ProductServiceHandler) UpdateProduct(ctx context.Context, req *connect.Request[productv1connect.UpdateProductRequest],
) (*connect.Response[productv1connect.UpdateProductResponse], error) {

	productID := req.Msg.ProductId
	if productID == 0 {
		return nil, connect.NewError(connect.CodeInvalidArgument, errors.New("product_id is required"))
	}
	name := req.Msg.Name
	price := req.Msg.Price
	quantity := req.Msg.Quantity

	if name == "" || price < 0 || quantity < 0 {
		return nil, connect.NewError(connect.CodeInvalidArgument, errors.New("invalid product data"))
	}

	product := &structs.Product{
		ProductID: productID,
		Name:      name,
		Price:     price,
		Quantity:  quantity,
	}

	productResponse, err := service.UpdateProduct(product)
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	response := &productv1connect.UpdateProductResponse{
		ProductId: productResponse.ProductID,
		Name:      productResponse.Name,
		Price:     productResponse.Price,
		Quantity:  productResponse.Quantity,
	}
	return connect.NewResponse(response), nil
}
