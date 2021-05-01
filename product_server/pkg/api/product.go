package api

import (
	"context"
	"golang-training-shop-grpc/product_server/pkg/data"
	pb "golang-training-shop-grpc/proto/go_proto"
	"gorm.io/gorm"
	"log"
)

const api = "v1"

type ProductServer struct {
	productData *data.ProductData
}

func NewProductServer(db *gorm.DB) *ProductServer {
	return &ProductServer{productData: data.NewProductData(db)}
}

func (p ProductServer) ReadAll(ctx context.Context, request *pb.ReadAllRequest) (*pb.ReadAllResponse, error) {
	var products []data.Product
	products, err := p.productData.ReadAll()
	if err != nil {
		log.Fatal(err)
	}
	var respProducts []*pb.Product
	for i := 0; i < len(products); i++ {
		product := &pb.Product{
			Id:                products[i].ID,
			Name:              products[i].Name,
			ProductCategoryId: products[i].ProductCategoryID,
			Quantity:          products[i].Quantity,
			UnitPrice:         products[i].UnitPrice,
		}
		respProducts = append(respProducts, product)
	}
	return &pb.ReadAllResponse{Product: respProducts}, nil
}

func (p ProductServer) Read(ctx context.Context, request *pb.ReadRequest) (*pb.ReadResponse, error) {
	var product data.Product
	product, err := p.productData.Read(request.Id)
	if err != nil {
		log.Fatal(err)
	}
	respProduct := pb.Product{
		Id:                product.ID,
		Name:              product.Name,
		ProductCategoryId: product.ProductCategoryID,
		Quantity:          product.Quantity,
		UnitPrice:         product.UnitPrice,
	}
	return &pb.ReadResponse{Product: &respProduct}, nil
}

func (p ProductServer) Create(ctx context.Context, request *pb.CreateRequest) (*pb.CreateResponse, error) {
	product := data.Product{
		Name:              request.Product.Name,
		ProductCategoryID: request.Product.ProductCategoryId,
		Quantity:          request.Product.Quantity,
		UnitPrice:         request.Product.UnitPrice,
	}
	id, err := p.productData.Create(product)
	if err != nil {
		log.Fatal(err)
	}
	return &pb.CreateResponse{Id: id}, nil
}

func (p ProductServer) Update(ctx context.Context, request *pb.UpdateRequest) (*pb.UpdateResponse, error) {
	panic("implement me")
}

func (p ProductServer) Delete(ctx context.Context, request *pb.DeleteRequest) (*pb.DeleteResponse, error) {
	panic("implement me")
}
