package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"main/golang/app/api/users"
	"main/golang/app/graph/generated"
	"main/golang/app/graph/model"
)

// User is the resolver for the user field.
func (r *queryResolver) User(ctx context.Context) (*model.ResponseSelectUser, error) {
	r.users = users.NewApiUser()
	dataUser, err := r.Resolver.users.SelectUser(ctx)

	if err != nil {
		return nil, err
	}

	data := &model.ResponseSelectUser{
		UserID:      &dataUser.UserId,
		UserEmail:   &dataUser.UserEmail,
		UserName:    &dataUser.UserName,
		UserImage:   &dataUser.UserImage,
		UserStatus:  &dataUser.UserStatus,
		CreatedDate: &dataUser.CreatedDate,
		LastUpdate:  &dataUser.LastUpdate,
	}

	for i := 0; i < len(dataUser.Products); i++ {
		data.Products = append(data.Products, &model.Product{
			UserID:       &dataUser.Products[i].UserId,
			ProductID:    &dataUser.Products[i].ProductId,
			ProductName:  &dataUser.Products[i].ProductName,
			ProductImage: &dataUser.Products[i].ProductImage,
			ProductInfo:  &dataUser.Products[i].ProductInfo,
			ProductStock: r.Resolver.users.Uint32ToInt(&dataUser.Products[i].ProductStock),
			ProductPrice: r.Resolver.users.Uint32ToInt(&dataUser.Products[i].ProductPrice),
			ProductSell:  r.Resolver.users.Uint32ToInt(&dataUser.Products[i].ProductSell),
			CreatedDate:  &dataUser.Products[i].CreatedDate,
			LastUpdate:   &dataUser.Products[i].LastUpdate,
		})
	}

	return data, nil
}

// Users is the resolver for the users field.
func (r *queryResolver) Users(ctx context.Context) (*model.ResponseSelectUsers, error) {
	return nil, nil
}

// Product is the resolver for the product field.
func (r *queryResolver) Product(ctx context.Context) (*model.Product, error) {
	panic(fmt.Errorf("not implemented"))
}

// Products is the resolver for the products field.
func (r *queryResolver) Products(ctx context.Context) ([]*model.Product, error) {
	panic(fmt.Errorf("not implemented"))
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
