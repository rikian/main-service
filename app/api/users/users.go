package users

import (
	"context"
	"errors"
	"log"
	"main/golang/app/config"
	"main/golang/app/graph/model"
	pb "main/golang/app/grpc/protos"
	h "main/golang/app/helper"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var helper = h.InitHelper()

type ApiUser interface {
	SelectUser(c context.Context) (*pb.ResponseSelectUser, error)
	SelectUsers(c context.Context) (*model.ResponseSelectUsers, error)
	Uint32ToInt(input *uint32) *int
}

type apiUserImpl struct {
	GrapUser  *model.ResponseSelectUser
	GrapUsers []*model.User
}

func NewApiUser() ApiUser {
	return &apiUserImpl{}
}

func userConnection(ctx context.Context) (pb.UserRPCClient, error) {
	conn, err := grpc.DialContext(ctx, config.GrpcAlphaAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Printf("connection to grpc server at %v failed", config.GrpcAlphaAddress)
		conn.Close()
		return nil, err
	}

	return pb.NewUserRPCClient(conn), nil
}

func (a *apiUserImpl) SelectUser(c context.Context) (*pb.ResponseSelectUser, error) {
	gc, ok := helper.GetContextId(c)

	if gc == nil || !ok {
		log.Print("could not retrieve gin.Context")
		return nil, errors.New("bad request")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	grpcClient, err := userConnection(ctx)
	defer cancel()

	if err != nil {
		log.Print(err.Error())
		return nil, err
	}

	user, err := grpcClient.SelectUser(ctx, &pb.DataSelectUser{
		Id: gc.Id,
	})

	if err != nil {
		log.Print(err.Error())
		return nil, err
	}

	return user, nil
}

func (a *apiUserImpl) SelectUsers(c context.Context) (*model.ResponseSelectUsers, error) {
	return nil, nil
}

func (a *apiUserImpl) Uint32ToInt(input *uint32) *int {
	i := int(*input)
	return &i
}
