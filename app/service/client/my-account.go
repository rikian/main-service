package client

import (
	"context"
	"log"
	"main/golang/app/api/products"
	"main/golang/app/config"
	"main/golang/app/helper"
	"net/http"

	pb "main/golang/app/grpc/protos"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/gin-gonic/gin"
)

type MyAccountService interface {
	MyAccount() gin.HandlerFunc
	GetAllData(h *handler.Server) gin.HandlerFunc
	UpdateProduct() gin.HandlerFunc
	InsertProduct() gin.HandlerFunc
	DeleteProduct() gin.HandlerFunc
}

type myAccountImpl struct {
	api products.ApiProduct
	h   helper.UtilsImpl
}

func NewserviceMyAccount() MyAccountService {
	return &myAccountImpl{
		h:   helper.InitHelper(),
		api: products.NewApiProduct(),
	}
}

func (s *myAccountImpl) MyAccount() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.HTML(200, "my-account.html", gin.H{
			"host_static": config.StaticAddressCDN,
		})
	}
}

func (s *myAccountImpl) GetAllData(h *handler.Server) gin.HandlerFunc {
	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func (s *myAccountImpl) InsertProduct() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx context.Context = c.Request.Context()
		c.Request.Body = http.MaxBytesReader(c.Writer, c.Request.Body, 10000)
		identity, ok := s.h.GetContextId(ctx)

		if !ok {
			c.Writer.Write([]byte("id not found"))
			return
		}

		product, err := s.h.ParseMultiPartData(c, identity.Id)

		if err != nil {
			log.Print(err.Error())
			c.Writer.Write([]byte("bad request"))
			return
		}

		saveProductToDB, err := s.api.InsertProduct(product)

		if err != nil {
			log.Print(err.Error())
			c.Writer.Write([]byte("something wrong when strore data product"))
			return
		}

		c.JSON(200, gin.H{
			"method": "add new product",
			"info":   saveProductToDB.Product,
			"status": "ok",
			"error":  false,
		})
	}
}

func (s *myAccountImpl) UpdateProduct() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx context.Context = c.Request.Context()
		c.Request.Body = http.MaxBytesReader(c.Writer, c.Request.Body, 1000000)
		identity, ok := s.h.GetContextId(ctx)

		if !ok {
			c.Writer.Write([]byte("id not found"))
			return
		}

		// check data update
		dataUpdateProduct, err := s.h.ParseMultiPartDataUpdate(c, identity.Id)

		if err != nil {
			log.Print(err.Error())
			c.Writer.Write([]byte("failed update product"))
			return
		}

		updateProductToDB, err := s.api.UpdateProduct(dataUpdateProduct)

		if err != nil {
			log.Print(err.Error())
			c.Writer.Write([]byte("failed update product"))
			return
		}

		c.JSON(200, updateProductToDB)
	}
}

func (s *myAccountImpl) DeleteProduct() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx context.Context = c.Request.Context()
		c.Request.Body = http.MaxBytesReader(c.Writer, c.Request.Body, 1000)
		identity, ok := s.h.GetContextId(ctx)
		id := c.Param("product_id")

		if !ok || id == "" {
			c.Writer.Write([]byte("id not found"))
			return
		}

		log.Print(id)

		dataDeleteProduct := &pb.DataDeleteProduct{
			UserId:    identity.Id,
			ProductId: id,
		}

		responseDeleteProduct, err := s.api.DeleteProduct(dataDeleteProduct)

		if err != nil || responseDeleteProduct.Message != "ok" {
			log.Print(err.Error())
			c.Writer.Write([]byte("forbidden"))
			return
		}

		c.JSON(200, responseDeleteProduct)
	}
}
