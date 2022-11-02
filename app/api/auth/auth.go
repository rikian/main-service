package auth

import (
	"context"
	"log"
	"main/golang/app/config"
	pb "main/golang/app/grpc/protos"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ApiAuth interface {
	SelectSessionUserById(id string) (*pb.ResponseSession, error)
	LoginUser(data *pb.DataLogin) (*pb.ResponseLogin, error)
	RegisterUser(data *pb.DataRegister) (*pb.ResponseRegister, error)
}

type apiAuthImpl struct{}

func NewApiAuth() ApiAuth {
	return &apiAuthImpl{}
}

func authConnection(ctx context.Context) (pb.AuthRPCClient, error) {
	conn, err := grpc.DialContext(ctx, config.GrpcAlphaAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Printf("connection to grpc server at %v failed", config.GrpcAlphaAddress)
		conn.Close()
		return nil, err
	}

	return pb.NewAuthRPCClient(conn), nil
}

func (a *apiAuthImpl) SelectSessionUserById(id string) (*pb.ResponseSession, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	grpcClient, err := authConnection(ctx)

	defer cancel()

	if err != nil {
		log.Print(err.Error())
		return nil, err
	}

	session, err := grpcClient.SelectSessionUserById(ctx, &pb.DataSession{
		Id: id,
	})

	if err != nil {
		log.Print(err.Error())
		return nil, err
	}

	log.Print(session.RememberMe)

	return session, nil
}

func (a *apiAuthImpl) LoginUser(data *pb.DataLogin) (*pb.ResponseLogin, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	grpcClient, err := authConnection(ctx)

	defer cancel()

	if err != nil {
		log.Print(err.Error())
		return nil, err
	}

	loginUser, err := grpcClient.LoginUser(ctx, data)

	if err != nil {
		log.Print(err.Error())
		return nil, err
	}

	return loginUser, nil
}

func (a *apiAuthImpl) RegisterUser(data *pb.DataRegister) (*pb.ResponseRegister, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	grpcClient, err := authConnection(ctx)

	defer cancel()

	if err != nil {
		log.Print(err.Error())
		return nil, err
	}

	registerUser, err := grpcClient.RegisterUser(ctx, data)

	if err != nil {
		log.Print(err.Error())
		return nil, err
	}

	return registerUser, nil
}
