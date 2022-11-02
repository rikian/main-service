package visitor

import (
	"log"
	"main/golang/app/api/auth"
	"main/golang/app/api/products"
	"main/golang/app/config"
	"main/golang/app/entities"
	pb "main/golang/app/grpc/protos"
	"main/golang/app/helper"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type Visitor interface {
	PostLogin() gin.HandlerFunc
	PostRegister() gin.HandlerFunc
	Home() gin.HandlerFunc
	GetProductById() gin.HandlerFunc
}

type visitorImpl struct {
	a auth.ApiAuth
	h helper.UtilsImpl
	p products.ApiProduct
}

func NewVisitor() Visitor {
	return &visitorImpl{
		h: helper.InitHelper(),
		a: auth.NewApiAuth(),
		p: products.NewApiProduct(),
	}
}

func (v *visitorImpl) PostLogin() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if !v.h.IsValidEmail(ctx.PostForm("user_email")) {
			ctx.JSON(200, gin.H{
				"message": "wrong email and password",
			})

			return
		}

		dataLogin := &pb.DataLogin{
			Email:      ctx.PostForm("user_email"),
			Password:   v.h.SHA256(ctx.PostForm("user_password")),
			RememberMe: v.h.StrToBool(ctx.PostForm("remember_me")),
		}

		loginUser, err := v.a.LoginUser(dataLogin)

		if err != nil {
			log.Print(err.Error())
			ctx.JSON(200, gin.H{
				"message": "sorry we are under maintenace",
			})
			return
		}

		if config.MainAddress != "127.0.0.1:8080" {
			var idCookie *http.Cookie = &http.Cookie{
				Name:     "id",
				Value:    loginUser.UserId,
				Path:     "/",
				HttpOnly: true,
			}

			var sessionCookie *http.Cookie = &http.Cookie{
				Name:     "session",
				Value:    strings.Split(loginUser.Session, ".")[2],
				Path:     "/",
				HttpOnly: true,
			}

			if dataLogin.RememberMe {
				idCookie.MaxAge = 31536000
				sessionCookie.MaxAge = 31536000
				http.SetCookie(ctx.Writer, idCookie)
				http.SetCookie(ctx.Writer, sessionCookie)
			} else {
				idCookie.MaxAge = 1800
				sessionCookie.MaxAge = 1800
				http.SetCookie(ctx.Writer, idCookie)
				http.SetCookie(ctx.Writer, sessionCookie)
			}
		} else {
			if dataLogin.RememberMe {
				ctx.SetCookie("id", loginUser.UserId, 31536000, "/", ".localhost", true, true)
				ctx.SetCookie("session", strings.Split(loginUser.Session, ".")[2], 31536000, "/", ".localhost", true, true)
			} else {
				ctx.SetCookie("id", loginUser.UserId, 1800, "/", ".localhost", true, true)
				ctx.SetCookie("session", strings.Split(loginUser.Session, ".")[2], 1800, "/", ".localhost", true, true)
			}
		}

		ctx.JSON(200, gin.H{
			"method":  "login",
			"message": "ok",
			"status":  200,
		})
	}
}

func (v *visitorImpl) PostRegister() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var user *entities.RequestRegister = &entities.RequestRegister{
			UserName:      ctx.PostForm("user_name"),
			UserEmail:     ctx.PostForm("user_email"),
			UserPassword1: ctx.PostForm("user_password_1"),
			UserPassword2: ctx.PostForm("user_password_2"),
			UserTerm:      v.h.StrToBool(ctx.PostForm("user_terms")),
		}

		if user.UserName == "" ||
			user.UserEmail == "" ||
			user.UserPassword1 == "" ||
			user.UserPassword2 == "" ||
			user.UserPassword1 != user.UserPassword2 ||
			!user.UserTerm ||
			!v.h.IsValidEmail(user.UserEmail) {
			ctx.JSON(200, gin.H{
				"message": "bad request",
			})

			return
		}

		registerUser, err := v.a.RegisterUser(&pb.DataRegister{
			UserName:     user.UserName,
			UserEmail:    user.UserEmail,
			UserPassword: v.h.SHA256(user.UserPassword1),
		})

		if err != nil {
			log.Print(err.Error())
			ctx.JSON(200, gin.H{
				"message": "sorry we are under maintenace",
			})
			return
		}

		log.Print(registerUser)

		ctx.JSON(200, gin.H{
			"method":  "register",
			"message": registerUser.Message,
			"status":  registerUser.Status,
		})
	}
}

func (v *visitorImpl) Home() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// get all new product by tanggal
		Products, err := v.p.GetAllProducts()

		if err != nil {
			ctx.Writer.Write([]byte("maintenance"))
			return
		}

		ctx.HTML(200, "public.html", gin.H{
			"host_main":   config.MainAddress,
			"host_image":  config.ImageAddressCDN,
			"host_static": config.StaticAddressCDN,
			"products":    Products,
		})
	}
}

type queryProductById struct {
	Pid   string `uri:"pid" binding:"required,uuid"`
	Pname string `uri:"pname" binding:"required"`
}

func (v *visitorImpl) GetProductById() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		prod := &queryProductById{}

		if err := ctx.ShouldBindUri(prod); err != nil {
			ctx.JSON(400, gin.H{"msg": "bad request"})
			return
		}

		P, err := v.p.GetProductById(prod.Pname, prod.Pid)

		if err != nil {
			ctx.JSON(404, gin.H{"msg": "product not found"})
			return
		}

		ctx.HTML(200, "product.html", gin.H{
			"host_image":  config.ImageAddressCDN,
			"host_static": config.StaticAddressCDN,
			"product":     P,
		})
	}
}
