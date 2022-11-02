package products

import (
	"context"
	"log"
	"main/golang/app/config"
	pb "main/golang/app/grpc/protos"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ApiProduct interface {
	GetAllProducts() (*pb.Products, error)
	GetProductById(productName, productId string) (*pb.Product, error)
	InsertProduct(p *pb.DataInsertProduct) (*pb.ResponseInsertProduct, error)
	DeleteProduct(p *pb.DataDeleteProduct) (*pb.ResponseDeleteProduct, error)
	UpdateProduct(p *pb.DataUpdateProduct) (*pb.ResponseUpdateProduct, error)
}

type apiProductImpl struct{}

func NewApiProduct() ApiProduct {
	return &apiProductImpl{}
}

func productConnection(ctx context.Context) (pb.ProductRPCClient, error) {
	conn, err := grpc.DialContext(ctx, config.GrpcAlphaAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Printf("connection to grpc server at %v failed", config.GrpcAlphaAddress)
		conn.Close()
		return nil, err
	}

	return pb.NewProductRPCClient(conn), nil
}

func (a *apiProductImpl) GetAllProducts() (*pb.Products, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	grpcClient, err := productConnection(ctx)

	defer cancel()

	if err != nil {
		log.Print(err.Error())
		return nil, err
	}

	products, err := grpcClient.GetAllProduct(ctx, &pb.User{
		Id: "1234",
	})

	if err != nil {
		log.Print(err.Error())
		return nil, err
	}

	return products, nil
}

func (a *apiProductImpl) GetProductById(productName, productId string) (*pb.Product, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	grpcClient, err := productConnection(ctx)

	defer cancel()

	if err != nil {
		log.Print(err.Error())
		return nil, err
	}

	product, err := grpcClient.GetProductById(ctx, &pb.RequestProduct{
		ProductName: productName,
		ProductId:   productId,
	})

	if err != nil {
		log.Print(err.Error())
		return nil, err
	}

	return product, nil
}

func (a *apiProductImpl) InsertProduct(p *pb.DataInsertProduct) (*pb.ResponseInsertProduct, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	grpcUser, err := productConnection(ctx)

	defer cancel()

	if err != nil {
		log.Print(err.Error())
		return nil, err
	}

	insertDataProduct, err := grpcUser.InsertProduct(ctx, p)

	if err != nil {
		log.Print(err.Error())
		return nil, err
	}

	return insertDataProduct, nil
}

func (a *apiProductImpl) DeleteProduct(p *pb.DataDeleteProduct) (*pb.ResponseDeleteProduct, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	grpcUser, err := productConnection(ctx)

	defer cancel()

	if err != nil {
		log.Print(err.Error())
		return nil, err
	}

	deleteDataProduct, err := grpcUser.DeleteProduct(ctx, p)

	if err != nil {
		return nil, err
	}

	return deleteDataProduct, nil
}

func (a *apiProductImpl) UpdateProduct(p *pb.DataUpdateProduct) (*pb.ResponseUpdateProduct, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	grpcClient, err := productConnection(ctx)

	defer cancel()

	if err != nil {
		log.Print(err.Error())
		return nil, err
	}

	updateDataProduct, err := grpcClient.UpdateProduct(ctx, p)

	if err != nil {
		return nil, err
	}

	return updateDataProduct, nil
}
