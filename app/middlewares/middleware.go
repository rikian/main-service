package middlewares

import (
	"log"
	"main/golang/app/api/users"
	"main/golang/app/config"
	"main/golang/app/helper"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
)

// var utils = u.InitHelper()

type Minddleware interface {
	MyAccountMiddleware(r *redis.Client) gin.HandlerFunc
	LoginRegisterMiddleware(r *redis.Client) gin.HandlerFunc
	AuthMiddleware(r *redis.Client) gin.HandlerFunc
}

type minddlewareImpl struct {
	helper.UtilsImpl
	users.ApiUser
}

func NewMinddleware() Minddleware {
	return &minddlewareImpl{}
}

func (m *minddlewareImpl) LoginRegisterMiddleware(r *redis.Client) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Next()
	}
}

func (m *minddlewareImpl) MyAccountMiddleware(r *redis.Client) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		m.UtilsImpl = helper.InitHelper()
		log.Print(ctx.Request.Method + " " + ctx.Request.URL.Path)
		// check cookies here
		validCookies, err := m.CookieParser(ctx)
		if err != nil {
			ctx.Abort()
			ctx.HTML(200, "auth.html", gin.H{
				"host_static": config.StaticAddressCDN,
			})

			return
		}

		// check session if valid cookies
		session := m.CheckSession(ctx, r, validCookies)
		if !session {
			ctx.Abort()
			// clear cookie session
			var idCookie *http.Cookie = &http.Cookie{
				Name:     "id",
				Value:    "expired",
				Path:     "/",
				MaxAge:   1,
				HttpOnly: true,
			}

			var sessionCookie *http.Cookie = &http.Cookie{
				Name:     "session",
				Value:    "expired",
				Path:     "/",
				MaxAge:   1,
				HttpOnly: true,
			}

			http.SetCookie(ctx.Writer, idCookie)
			http.SetCookie(ctx.Writer, sessionCookie)
			ctx.HTML(200, "auth.html", gin.H{
				"host_static": config.StaticAddressCDN,
			})

			return
		}

		ctx.Next()
	}
}

func (m *minddlewareImpl) AuthMiddleware(r *redis.Client) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		m.UtilsImpl = helper.InitHelper()
		m.ApiUser = users.NewApiUser()
		// check cookies here
		validCookies, err := m.CookieParser(ctx)
		if err == nil {
			// user have valid cookies, so we need check session
			session := m.CheckSession(ctx, r, validCookies)
			if session {
				// user have valid session and try to login, so we need throw error
				ctx.Abort()
				ctx.JSON(403, gin.H{
					"method":  "login",
					"message": "forbidden",
					"status":  403,
				})
				return
			}

			// user have cookied but invalid session, so we need set current session to expired
			var idCookie *http.Cookie = &http.Cookie{
				Name:     "id",
				Value:    "expired",
				Path:     "/",
				MaxAge:   1,
				HttpOnly: true,
			}

			var sessionCookie *http.Cookie = &http.Cookie{
				Name:     "session",
				Value:    "expired",
				Path:     "/",
				MaxAge:   1,
				HttpOnly: true,
			}

			http.SetCookie(ctx.Writer, idCookie)
			http.SetCookie(ctx.Writer, sessionCookie)
			ctx.Abort()
			ctx.JSON(403, gin.H{
				"method":  "login",
				"message": "forbidden",
				"status":  403,
			})
			return
		}

		ctx.Next()
	}
}

// 	ctx.Header("Access-Control-Allow-Origin", "*")
// 	ctx.Header("Access-Control-Allow-Headers", "*")

// 	if ctx.Request.Method == http.MethodOptions {
// 		log.Print(ctx.Request.Method + " " + ctx.Request.URL.Path)
// 		ctx.AbortWithStatus(404)
// 		return
// 	}

// 	// auth
// 	if ctx.Request.Method == http.MethodGet {
// 		log.Print(ctx.Request.Method + " " + ctx.Request.URL.Path)
// 		// check cookies
// 		auth, err := utils.CookieParser(ctx)
// 		if err != nil {
// 			ctx.Abort()
// 			ctx.HTML(200, "auth.html", gin.H{
// 				"host": config.StaticAddressCDN,
// 			})
// 			return
// 		}

// 		// check session
// 		session := utils.CheckSession(ctx, r, auth)
// 		if !session {
// 			ctx.Abort()
// 			ctx.HTML(200, "auth.html", gin.H{
// 				"host": config.StaticAddressCDN,
// 			})
// 			return
// 		}

// 		ctx.Next()
// 		return
// 	}

// 	if ctx.Request.Method == http.MethodPost {
// 		log.Print(ctx.Request.Method + " " + ctx.Request.URL.Path)
// 		// check auth
// 		auth, err := utils.CookieParser(ctx)
// 		// if no auth
// 		if err != nil {
// 			switch ctx.Request.URL.Path {
// 			case "/login":
// 				ctx.Next()
// 				return
// 			case "/register":
// 				ctx.Next()
// 				return
// 			default:
// 				ctx.AbortWithStatus(http.StatusBadRequest)
// 				return
// 			}
// 		}

// 		// check session
// 		session := utils.CheckSession(ctx, r, auth)
// 		if !session {
// 			var idCookie *http.Cookie = &http.Cookie{
// 				Name:     "id",
// 				Value:    "expired",
// 				Path:     "/",
// 				MaxAge:   1,
// 				HttpOnly: true,
// 			}

// 			var sessionCookie *http.Cookie = &http.Cookie{
// 				Name:     "session",
// 				Value:    "expired",
// 				Path:     "/",
// 				MaxAge:   1,
// 				HttpOnly: true,
// 			}

// 			http.SetCookie(ctx.Writer, idCookie)
// 			http.SetCookie(ctx.Writer, sessionCookie)
// 			ctx.AbortWithStatus(http.StatusNotFound)
// 			return
// 		}

// 		ctx.Next()
// 		return
// 	}

// 	if ctx.Request.Method == http.MethodDelete {
// 		log.Print(ctx.Request.Method + " " + ctx.Request.URL.Path)
// 		// check auth
// 		auth, err := utils.CookieParser(ctx)
// 		// if no auth
// 		if err != nil {
// 			ctx.AbortWithStatus(http.StatusForbidden)
// 			return
// 		}

// 		// check session
// 		session := utils.CheckSession(ctx, r, auth)
// 		if !session {
// 			ctx.AbortWithStatus(http.StatusForbidden)
// 			return
// 		}

// 		ctx.Next()
// 		return
// 	}

// 	if ctx.Request.Method == http.MethodPut {
// 		log.Print(ctx.Request.Method + " " + ctx.Request.URL.Path)
// 		// check auth
// 		auth, err := utils.CookieParser(ctx)
// 		// if no auth
// 		if err != nil {
// 			ctx.AbortWithStatus(http.StatusForbidden)
// 			return
// 		}

// 		// check session
// 		session := utils.CheckSession(ctx, r, auth)
// 		if !session {
// 			ctx.AbortWithStatus(http.StatusForbidden)
// 			return
// 		}

// 		ctx.Next()
// 		return
// 	}

// 	ctx.AbortWithStatus(http.StatusMethodNotAllowed)
