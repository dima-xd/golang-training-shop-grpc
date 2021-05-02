package api

import (
	"context"

	"golang-training-shop-grpc/product_server/pkg/data"
	pb "golang-training-shop-grpc/proto/go_proto"

	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

const api = "v1"

type ProductServer struct {
	productData *data.ProductData
}

func NewProductServer(db *gorm.DB) *ProductServer {
	log.SetFormatter(&log.JSONFormatter{})
	return &ProductServer{productData: data.NewProductData(db)}
}

func (p ProductServer) ReadAll(ctx context.Context, request *pb.ReadAllRequest) (*pb.ReadAllResponse, error) {
	var products []data.Product
	products, err := p.productData.ReadAll()
	if err != nil {
		log.Fatal("got an error when tried to read all products")
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
		log.WithFields(log.Fields{
			"id":                  products[i].ID,
			"name":                products[i].Name,
			"product_category_id": products[i].ProductCategoryID,
			"quantity":            products[i].Quantity,
			"unit_price":          products[i].UnitPrice,
		}).Info("read product")
		respProducts = append(respProducts, product)
	}
	return &pb.ReadAllResponse{Product: respProducts}, nil
}

func (p ProductServer) Read(ctx context.Context, request *pb.ReadRequest) (*pb.ReadResponse, error) {
	var product data.Product
	product, err := p.productData.Read(request.Id)
	if err != nil {
		log.Fatal("got an error when tried to read product")
	}
	respProduct := pb.Product{
		Id:                product.ID,
		Name:              product.Name,
		ProductCategoryId: product.ProductCategoryID,
		Quantity:          product.Quantity,
		UnitPrice:         product.UnitPrice,
	}
	log.WithFields(log.Fields{
		"id":                  product.ID,
		"name":                product.Name,
		"product_category_id": product.ProductCategoryID,
		"quantity":            product.Quantity,
		"unit_price":          product.UnitPrice,
	}).Info("read product")
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
		log.Fatal("got an error when tried to create product")
	}
	log.WithFields(log.Fields{
		"id":                  id,
		"name":                request.Product.Name,
		"product_category_id": request.Product.ProductCategoryId,
		"quantity":            request.Product.Quantity,
		"unit_price":          request.Product.UnitPrice,
	}).Info("create product")
	return &pb.CreateResponse{Id: id}, nil
}

func (p ProductServer) Update(ctx context.Context, request *pb.UpdateRequest) (*pb.UpdateResponse, error) {
	err := p.productData.Update(request.Id, request.UnitPrice)
	if err != nil {
		log.Fatal("got an error when tried to update product")
	}
	log.WithFields(log.Fields{
		"id":         request.Id,
		"unit_price": request.UnitPrice,
	}).Info("update product")
	return &pb.UpdateResponse{}, nil
}

func (p ProductServer) Delete(ctx context.Context, request *pb.DeleteRequest) (*pb.DeleteResponse, error) {
	err := p.productData.Delete(request.Id)
	if err != nil {
		log.Fatal("got an error when tried to delete product")
	}
	log.WithFields(log.Fields{
		"id": request.Id,
	}).Info("delete product")
	return &pb.DeleteResponse{}, nil
}
